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


__version__ = '0.1.2'

# Enable Debug in env
debug = os.getenv('DEBUG', False)

app = Flask(__name__)

#MONGO_URI="mongodb://mongo:mongo@mongo:27017/test"

# This... doesn't work so well...
if os.getenv('MONGO_URI'):
    MONGO_URI = os.getenv('MONGO_URI')
else:
    app.config['MONGO_HOST'] = os.getenv('MONGO_HOST', 'mongo')
    app.config['MONGO_PORT'] = os.getenv('MONGO_PORT', 27017)
    app.config['MONGO_USERNAME'] = os.getenv('MONGO_USERNAME', 'mongo')
    app.config['MONGO_PASSWORD'] = os.getenv('MONGO_PASSWORD', 'mongo')
    app.config['MONGO_DBNAME'] = os.getenv('MONGO_DBNAME', 'test')

    MONGO_URI = "mongodb://%s:%s@%s:%d/%s" % (app.config['MONGO_USERNAME'].rstrip("\n"),
                                              app.config['MONGO_PASSWORD'].rstrip("\n"),
                                              app.config['MONGO_HOST'].rstrip("\n"),
                                              int(app.config['MONGO_PORT']),
                                              app.config['MONGO_DBNAME'].rstrip("\n")
                                            )

app.config['MONGO_CONNECT'] = os.getenv('MONGO_CONNECT', False)

# wrap the flask app and give a heathcheck url
health = HealthCheck(app, "/healthz")
envdump = EnvironmentDump(app, "/envz")

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

logger.debug("MONGO_URI: %s" % MONGO_URI)
app.config['MONGO_URI'] = MONGO_URI

mongo = PyMongo(app)

# add your own check function to the healthcheck
def mongo_available():
    if debug:
        logger.debug('MONGO_URI: %s' % app.config['MONGO_URI'])
    try:
        result = mongo.db.test.insert_one({'status': 'OK'})
        if result:
            return True, "mongoDB status"
        else:
            return False, "Failed to Insert"
    except:
        logger.info("Failed to insert 'status: ok' in test collection")
        logger.info("MONGO_URI: %s" % app.config['MONGO_URI'])

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
    data = request.get_json()
    if not data or not 'name' in data or not 'topics' in data:
        abort(400)
    topic = {
        'name': data['name'],
        'topics': data['topics'],
        'date_updated': datetime.datetime.utcnow(),
        'date_created': datetime.datetime.utcnow()
    }
    query = {'name': data['name']}
    existing = mongo.db.monitors.find_one(query)
    if existing:
        import pprint
        pprint.pprint(existing)
        logger.info("Add/Update Topics: Found id: %s" % existing['_id'])
        update = { '$set': {
            'topics': data['topics'],
            'date_updated': datetime.datetime.utcnow()
        }}
        result = mongo.db.monitors.update_one({'_id': existing['_id']}, update)
    else:
        logger.info("Monitor Name Does not Exist, creating new")
        result = mongo.db.monitors.insert_one(topic)

    if result.acknowledged:
        return jsonify('{"status": "Success", "message": "Updated Monitor: %s"}' % data['name']), 201
    else:
        return jsonify('{"status": "Fail", "message": "Failed to create topic"}'), 400


@app.route("/v1/topics", methods=['GET'])
def topics():
    """
    Grab Monitor Topics
    """
    topics = []
    try:
        results = mongo.db.monitors.find({}, {'topics': 1})
        for result in results:
            logger.info("result: %s" % result)
            for r in result['topics']:
                topics.append(r)
        if len(topics) > 0:
            return jsonify({"status": "OK", "topics": topics }), 200
        else:
            return jsonify('{"status": "None"}'), 404
    except:
        return jsonify('{"status": "Error"}'), 400


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
