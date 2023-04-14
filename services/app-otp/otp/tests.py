import unittest
from io import StringIO
from unittest.mock import patch
import hmac
import hashlib
import struct
import time

class TestOTP(unittest.TestCase):
    
    def setUp(self):
        self.secret = b'MysecretKey'
        self.time_step = 120 
        self.current_time = 1620396920 #sample unix timestamp
        
    def test_otp_generation(self):
        
        time_steps = int(self.current_time / self.time_step)
        time_bytes = struct.pack('>Q', time_steps)
        hash = hmac.new(self.secret, time_bytes, hashlib.sha1).digest()
        offset = hash[-1] & 0x0F
        otp_bytes = hash[offset:offset+4]
        otp_int = struct.unpack('>I', otp_bytes)[0]
        expected_otp = otp_int % 10**6
        
        # Act
        with patch('builtins.input', return_value=expected_otp):
            with patch('sys.stdout', new=StringIO()) as fake_output:
                exec(open('./main.py').read())
                actual_output = fake_output.getvalue().strip()

        # Assert
        self.assertEqual(actual_output, f'Your OTP is : {expected_otp}\nEnter your OTP: OTP is valid')

    
    def test_invalid_otp(self):
        # Arrange
        time_steps = int(self.current_time / self.time_step)
        time_bytes = struct.pack('>Q', time_steps)
        hash = hmac.new(self.secret, time_bytes, hashlib.sha1).digest()
        offset = hash[-1] & 0x0F
        otp_bytes = hash[offset:offset+4]
        otp_int = struct.unpack('>I', otp_bytes)[0]
        expected_otp = otp_int % 10**6

        # Act
        with patch('builtins.input', return_value='000000'):
            with patch('sys.stdout', new=StringIO()) as fake_output:
                exec(open("./main.py").read())
                actual_output = fake_output.getvalue().strip()

        # Assert
        self.assertEqual(actual_output, f"Your OTP is: {expected_otp}\nEnter your OTP: OTP is invalid")

if __name__ == '__main__':
    unittest.main()