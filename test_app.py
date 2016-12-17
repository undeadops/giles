import urllib2
from flask import Flask
from flask_testing import LiveServerTestCase
import unittest

try:
    from coverage import coverage
    coverage_available = True
except ImportError:
    coverage_available = False

from app import app
# Testing with LiveServer
class Test1(LiveServerTestCase):
  # if the create_app is not implemented NotImplementedError will be raised
  def create_app(self):

    app.config['TESTING'] = True
    return app

  def test_flask_application_is_up_and_running(self):
    response = urllib2.urlopen(self.get_server_url())
    self.assertEqual(response.code, 200)

  def test_envz_url(self):
      response = urllib2.urlopen(self.get_server_url() + "/envz")
      self.assertEqual(response.code, 200)

  def test_healthz_url(self):
      response = urllib2.urlopen(self.get_server_url() + "/healthz")
      self.assertEqual(response.code, 200)


if __name__ == '__main__':
    unittest.main()
