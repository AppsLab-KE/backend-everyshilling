import os
import unittest
from otp import generate_otp, send_otp


class TestOTP(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        # Set up the environment variables
        os.environ['OTP_SECRET'] = 'my_secret_key'
        os.environ['AFR_USERNAME'] = 'my_africastalking_username'
        os.environ['AFRICASTALKING_API_KEY'] = 'my_africastalking_api_key'

    def test_generate_otp(self):
        # Test that an OTP is generated successfully
        otp = generate_otp()
        self.assertIsInstance(otp, str)
        self.assertEqual(len(otp), 6)

    def test_send_otp(self):
        # Test that an OTP is sent successfully
        phone_number = '+254701847888'
        otp = generate_otp()
        result = send_otp(phone_number, otp)
        self.assertTrue(result)


if __name__ == '__main__':
    unittest.main()