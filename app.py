from flask import Flask, request, jsonify, abort, make_response
from flask_pymongo import PyMongo
from healthcheck import HealthCheck, EnvironmentDump
import requests
import random
from time import sleep
import os


__version__ = '0.1'

# Define default MONGO_URI - Can be modified through environment var
MONGO_URI = 'mongodb://mongo:27017/social_scraps'
MONGO_MAX_POOL_SIZE = '75'

app = Flask(__name__)
app.config['MONGO_URI'] = os.getenv('MONGO_URI', MONGO_URI)
app.config['MONGO_MAX_POOL_SIZE'] = os.getenv('MONGO_MAX_POOL_SIZE', MONGO_MAX_POOL_SIZE)

# wrap the flask app and give a heathcheck url
health = HealthCheck(app, "/healthz")
envdump = EnvironmentDump(app, "/envz")

mongo = PyMongo(app)


# add your own check function to the healthcheck
def mongo_available():
    result = mongo.db.test.insert_one({'status': 'OK'})

    return True, "mongoDB status"

health.add_check(mongo_available)


# add your own data to the environment dump
def application_data():
    return {"maintainer": "Mitch Anderson",
            "git_repo": "https://github.com/undeadops/giles"}

envdump.add_section("application", application_data)

@app.route("/v1/posts", methods=['PUT'])
def create_post():
    """
    Push Post to MongoDB
    """
    if not request.json or not 'process_time' in request.json:
        abort(400)
    else:
        result = mongo.db.twitter.insert_one(request.json)
        return jsonify({'status': result}), 201


# Needs Changing
@app.route('/')
def index():
    return jsonify("hello world")


def main():
    print "%s (v%s) starting up..." % ('giles', __version__)
    app.run(host='0.0.0.0', port=5000, debug=True)


if __name__ == '__main__':
    main()
