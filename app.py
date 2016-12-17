from flask import Flask, request, jsonify, abort, make_response, json
from flask_pymongo import PyMongo
from healthcheck import HealthCheck, EnvironmentDump
import requests
import random
from time import sleep
import datetime
from timeit import default_timer as timer
import os
import logging


__version__ = '0.1'

# Enable Debug in env
debug = os.getenv('DEBUG', False)

# Define default MONGO_URI - Can be modified through environment var
MONGO_URI = 'mongodb://mongo:27017/test'
MONGO_MAX_POOL_SIZE = '75'

app = Flask(__name__)
app.config['MONGO_URI'] = os.getenv('MONGO_URI', MONGO_URI)
app.config['MONGO_MAX_POOL_SIZE'] = os.getenv('MONGO_MAX_POOL_SIZE', MONGO_MAX_POOL_SIZE)

# wrap the flask app and give a heathcheck url
health = HealthCheck(app, "/healthz")
envdump = EnvironmentDump(app, "/envz")

mongo = PyMongo(app)

# Check if debug mode is required
debug = os.getenv('DEBUG', False)
logger = logging.getLogger()
handler = logging.StreamHandler()
formatter = logging.Formatter(
    '%(asctime)s %(name)-12s %(levelname)-8s %(message)s')
handler.setFormatter(formatter)
logger.addHandler(handler)

if debug:
    loglevel = logging.DEBUG
else:
    loglevel = os.environ.get("LOG_LEVEL", logging.INFO)
logger.setLevel(loglevel)

# add your own check function to the healthcheck
def mongo_available():
    result = mongo.db.test.insert_one({'status': 'OK'})
    if result:
        return True, "mongoDB status"
    else:
        return False, "Failed to Insert"

health.add_check(mongo_available)


# add your own data to the environment dump
def application_data():
    return {"maintainer": "Mitch Anderson",
            "git_repo": "https://github.com/undeadops/giles"}

envdump.add_section("application", application_data)

@app.route("/v1/posts", methods=['POST'])
def create_post():
    """
    Push Post to MongoDB
    """
    if request.headers['Content-Type'] == 'application/json':
        result = mongo.db.twitter.insert_one(request.json)
        if result.inserted_id:
            status = '{"status": "OK", "id": %s}' % result.inserted_id
            return jsonify(status), 201
        else:
            status = '{"status": "Failed"}'
            return jsonify(status), 406
    else:
        abort(415)

@app.route("/v1/topics/<string:name>", methods=['PUT'])
def updateTopics(name):
    """
    Update Topic monitors
    """
    logging.debug()
    if not request.json or not 'topics' in request.json:
        abort(400)
    if not 'action' in request.json:
        abort(400)
    r = request.json()
    topic = mongo.db.monitors.find_one_or_404({'name': name})

    if r['action'] == 'update':
        topic['topics'].append(r['topics'])
        topic['date_updated'] = datetime.datetime.utcnow()
    if r['action'] == 'set':
        topic['topics'] = r['topics']
        topic['date_updated'] = datetime.datetime.utcnow()

    result = mongo.db.monitors.update_one(topic)
    if result.modified_count == 1:
        return jsonify({"status": "Success", "id": result.upserted_id}), 201


@app.route("/v1/topics", methods=['POST'])
def addTopics():
    """
    Add Topic monitors
    """
    if not request.json or not 'name' in request.json or not 'topics' in request.json:
        abort(400)

    topic = {
        'name': request.json['name'],
        'topics': request.json['topics'],
        'date_updated': datetime.datetime.utcnow(),
        'date_created': datetime.datetime.utcnow()
    }

    result = mongo.db.monitors.insert_one(topic)
    if result.inserted_id:
        return jsonify('{"status": "Success", "_id": %s }' % result.inserted_id), 201
    else:
        return jsonify('{"status": "Fail", "message": "Failed to insert topic"}'), 400


@app.route("/v1/topics", methods=['GET'])
def topics():
    """
    Grab Monitor Topics
    """
    topics = []
    try:
        results = mongo.db.monitors.find({}, {'topics': 1})
        print results
        for result in results:
            for r in result['topics']:
                topics.append(r)
        return jsonify('{"status": "OK", "topics", %s }' % topics), 201
    except:
        return jsonify('{"status": "Fail"}'), 400


# Needs Changing
@app.route('/')
def index():
    return jsonify("hello world")


def main():
    print "%s (v%s) starting up..." % ('giles', __version__)
    if debug:
        print "Running in Debug Mode"
        app.run(host='0.0.0.0', port=5000, debug=True)
    else:
        app.run(host='0.0.0.0', port=5000)

if __name__ == '__main__':
    main()
