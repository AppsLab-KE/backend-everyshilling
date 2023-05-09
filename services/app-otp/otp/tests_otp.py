import os
import unittest
from unittest.mock import Mock, patch
import otp

class TestOTPFunctions(unittest.TestCase):
    @patch ('otp.struct')
    @patch ('otp.hmac')
    def test_generate_otp(self, hmac_mock,struct_mock):
        struct_mock.pack.return_value= b'test_bytes'
        hmac_mock.new.return_value= b'test_hash'
        struct_mock.upack.return_value= (123456)
        
        
        generated_otp=otp.generate_otp()
        
        self.assertEqual(generated_otp,'234567')
        
        hmac_mock.new.assert_called_once_with(b"test_secret", b"test_bytes", otp.hashlib.sha1)
        struct_mock.pack.assert_called_once_with(">Q", Mock())
        struct_mock.unpack.assert_called_once_with(">I", b"test_hash"[11:15])
        
    @patch("otp.sms")
    def test_send_otp_success(self, sms_mock):
        
        response = {"SMSMessageData": {"Recipients": [{"status": "Success"}]}}
        sms_mock.send.return_value = response
        
        
        result = otp.send_otp('+254701847888', '123456')
        
        self.assertTrue(result)
        sms_mock.send.assert_called_once_with('Your OTP is: 123456. It expires in 5 minutes. Please do not share it with anyone.', ['+254701847888'])       
        
    @patch("otp.os")
    @patch("otp.sms")
    def test_send_otp_failure(self, sms_mock, os_mock):
        
        response = {"SMSMessageData": {"Recipients": [{"status": "Failed"}]}}
        sms_mock.send.return_value = response
        
        result = otp.send_otp("+254701847888", "123456")
        
        self.assertFalse(result)
        sms_mock.send.assert_called_once_with('Your OTP is: 123456. It expires in 5 minutes. Please do not share it with anyone.', ['+254701847888'])
        os_mock.write.assert_called_once_with(2, b'{"SMSMessageData": {"Recipients": [{"status": "Failed"}]}}\n')    