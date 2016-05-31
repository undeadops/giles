from flask import Flask, request
from flask_restful import Resource, Api
import requests
import pymongo
import random
from time import sleep
import os

NAME = 'giles'
VERSION = '0.1'

# Define default MONGO_URI - Can be modified through environment var
MONGO_URI = 'mongodb://mongo:27017'

app = Flask(__name__)
api = Api(app)

def _get_random_num():
    """
    Return Random number between 5-30
    """
    return random.randint(5,30)


def _get_mongodb_conn(mongo_host):
    """
    Get MongoDB connection
    """
    try:
        print "Connecting to: %s" % mongo_host
        mongo_conn = pymongo.MongoClient(mongo_host, connect=True, connectTimeoutMS=10000)
        mongo_conn.server_info()
        return mongo_conn
    except pymongo.errors.ConnectionFailure, e:
        print "Error Connecting to: %s" % e
        return False

def _connect_mongodb(mongo_host):
    """
    Connect to mongodb, return connection
    """
    print "Starting MongoDB connection"
    mongo_conn = _get_mongodb_conn(mongo_host)
    while not mongo_conn:
            sleep_time = _get_random_num()
            print "Problem Connecting to MongoDB, Waiting %s before trying again" % sleep_time
            sleep(sleep_time)
            mongo_conn = _get_mongodb_conn(mongo_host)
    return mongo_conn

class GilesPosts(Resource):
    def post(self):
        

    def get(self):
        queryText = request.form['text']
        result = TextBlob(queryText)
        return '{ "sentinment": { "polarity": %s, "subjectivity": %s } }' % (result.sentiment.polarity, result.sentiment.subjectivity)

api.add_resource(JudgeDredd, '/api/v1/')


def main():
    print "%s (%s) starting up..." % (NAME, VERSION)
    # maybe print some environment vars here...
    # Testing connectiong to "mongo"
    mongo_host = os.getenv('MONGO_URI', MONGO_URI)
    print "Trying Connecting to MONGO_URI: %s" % mongo_host
    mongo_conn = _connect_mongodb(mongo_host)



if __name__ == '__main__':
    main()
