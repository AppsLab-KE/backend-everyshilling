import os 
import unittest
from datetime import datetime, time 
import redis
from unittest.mock import patch 
import cache

REDIS_HOST = os.getenv("REDIS_HOST", "localhost")
REDIS_PORT = os.getenv("REDIS_PORT", "6379")


r = redis.Redis(host=REDIS_HOST, port=REDIS_PORT, db=0)


class TestOTP(unittest.TestCase):
    
    @classmethod
    def setUpClass(cls):
        '''
        Setting up Redis environment variables for testing.
        '''
        os.environ['REDIS_HOST'] ='Localhost'
        os.environ['REDIS_PORT'] ='6379'
        
    def setUp(self):
        '''
        This flush redis database before each test is executed.
        '''
        r.flushdb()
        
    def test_save_and_get_otp(self):

        phone_number = '+254701847888'
        tracking_uuid = 'abcdc123'
        otp = '123456'
        save_otp (phone_number, tracking_uuid, otp) 
        result = get_otp(phone_number)
        expected = {"otp": otp.encode(), "time_stamp_unix": str(datetime.now().timestamp()).encode()}
        self.assertEqual(result, expected)
        
    def test_get_phone_number(self):
        phone_number = '+254701847888'
        tracking_uuid = 'abcdc123'
        otp = '123456'
        save_otp(phone_number, tracking_uuid, otp)
        result = get_phone_number(tracking_uuid) 
        expected = {
            "phone_number": phone_number.encode(),
            "otp": otp.encode(),
            "time_stamp_unix": str(datetime.now().timestamp()).encode(),
        }
        self.assertEqual(result, expected)
        
    def test_delete_otp_phone(self):
        
        phone_number = '+254701847888'
        tracking_uuid = 'abcdc123'
        otp = '123456'
        save_otp(phone_number, tracking_uuid, otp)     
        delete_otp_phone(phone_number, tracking_uuid)
        result = get_otp(phone_number)
        self.assertFalse(result)
        
    def test_delete_otp_tracking(self):
        
        phone_number = '+254701847888'
        tracking_uuid = 'abcdc123'
        otp = '123456'
        save_otp(phone_number, tracking_uuid, otp)     
        delete_otp_tracking(phone_number, tracking_uuid)
        result = get_otp(tracking_uuid)
        self.assertFalse(result)
        
    @patch.object(datetime, "now")
    def test_otp_expiry(self, mock_now):
         
        phone_number = '+254701847888'
        tracking_uuid = 'abcdc123'
        otp = '123456'
        save_otp(phone_number, tracking_uuid, otp)     
        mock_now.return_value = datetime.now() + timedelta(minutes=5) 
        
        result = get_otp(phone_number)
        self.assertFalse(result)
        