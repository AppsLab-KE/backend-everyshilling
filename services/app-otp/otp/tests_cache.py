import os
import unittest
from datetime import datetime, timedelta
import redis
from unittest.mock import patch
from cache import save_otp, get_otp, get_phone_number, delete_otp_phone, delete_otp_tracking

REDIS_HOST = os.getenv("REDIS_HOST", "localhost")
REDIS_PORT = int(os.getenv("REDIS_PORT", '6379'))

r = redis.Redis(host=REDIS_HOST, port=REDIS_PORT, db=0)


def get_mock_time():
    # Return a mock time 5 minutes ahead
    return datetime.now() + timedelta(minutes=5)


class TestOTP(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        """
        Setting up Redis environment variables for testing.
        """
        os.environ['REDIS_HOST'] = 'Localhost'
        os.environ['REDIS_PORT'] = '6379'

    def setUp(self):
        """
        This flush redis database before each test is executed.
        """
        r.flushdb()

    def test_save_and_get_otp(self):
        phone_number = '+254701847888'
        tracking_uuid = 'abc-dc123'
        otp = '123456'
        save_otp(phone_number, tracking_uuid, otp)
        result = get_otp(phone_number)
        expected = {
            b'otp': otp.encode(),
            b'time_stamp_unix': str(round(datetime.now().timestamp(), 2)).encode()
        }
        self.assertEqual(result, expected)

    def test_get_phone_number(self):
        phone_number = '+254701847888'
        tracking_uuid = 'abc-dc123'
        otp = '123456'
        save_otp(phone_number, tracking_uuid, otp)
        result = get_phone_number(tracking_uuid)
        expected = {
            b"phone_number": phone_number.encode(),
            b"otp": otp.encode(),
            b"time_stamp_unix": str(round(datetime.now().timestamp(), 2)).encode(),
        }
        self.assertEqual(result, expected)

    def test_delete_otp_phone(self):
        phone_number = '+254701847888'
        tracking_uuid = 'abc-dc123'
        otp = '123456'
        save_otp(phone_number, tracking_uuid, otp)
        delete_otp_phone(phone_number, tracking_uuid)
        result = get_otp(phone_number)
        self.assertFalse(result)

    def test_delete_otp_tracking(self):
        phone_number = '+254701847888'
        tracking_uuid = 'abc-dc123'
        otp = '123456'
        save_otp(phone_number, tracking_uuid, otp)
        delete_otp_tracking(phone_number, tracking_uuid)
        result = get_otp(tracking_uuid)
        self.assertFalse(result)

    def test_otp_expiry(self):
        phone_number = '+254701847888'
        tracking_uuid = 'abc-dc123'
        otp = '123456'
        save_time = datetime.now()
        save_otp(phone_number, tracking_uuid, otp)

        with patch('cache.datetime') as mock_datetime:
            mock_datetime.now.side_effect = lambda: get_mock_time(save_time)
            result = get_otp(phone_number)

        expected = {
            b'otp': otp.encode(),
            b'time_stamp_unix': str(round(get_mock_time().timestamp(), 2)).encode()
        }
        # Check if the OTP is expired
        current_time = datetime.now().timestamp()
        expiration_time = float(result.get(b'time_stamp_unix', 0))
        self.assertTrue(current_time > expiration_time)


if __name__ == '__main__':
    unittest.main()
