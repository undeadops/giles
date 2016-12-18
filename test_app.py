from app import app
import os
import json
import unittest

try:
    from coverage import coverage
    coverage_available = True
except ImportError:
    coverage_available = False


class HeathzTests(unittest.TestCase):
    def test_envz_url(self):
        tester = app.test_client(self)
        response = tester.get('/envz', content_type='application/json')
        self.assertEqual(response.status_code, 200)

    def test_healthz_url(self):
        tester = app.test_client(self)
        response = tester.get('/healthz', content_type='application/json')
        self.assertEqual(response.status_code, 200)

    def test_empty_topics(self):
        tester = app.test_client(self)
        response = tester.get("/v1/topics", content_type="application/json")
        self.assertEqual(response.status_code, 404)
        # Check that the result sent is 8: 2+6
        self.assertEqual(json.loads(response.data), {"status": "None"})


if __name__ == '__main__':
    unittest.main()
