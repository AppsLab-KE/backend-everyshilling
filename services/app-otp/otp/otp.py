import binascii
import time
import hmac
import hashlib
import struct
# from Crypto.PublicKey import RSA
import os
import africastalking

africastalking_username = os.getenv('AFR_USERNAME')
africastalking_api_key = os.getenv('AFRICASTALKING_API_KEY')

africastalking.initialize(africastalking_username, africastalking_api_key)
sms = africastalking.SMS


def generate_otp() -> str:
    # Generating RSA keys
    secret = os.getenv('OTP_SECRET')

    if not secret:
        raise Exception("otp key missing")

    # Set the steps and get the unix time
    time_step = 300
    current_time = int(time.time())

    # Get the number of time steps
    time_steps = int(current_time / time_step)

    # Pack the time steps as a big-endian byte string
    time_bytes = struct.pack('>Q', time_steps)

    secret = str.encode(secret)

    # Calculate the HMAC-SHA1 hash of the time steps using the secret key

    hash = hmac.new(secret, time_bytes, hashlib.sha1).digest()
    hash_hex = binascii.hexlify(hash).decode('utf-8')
    # Get the last 4 bits of the hash
    last_char = hash_hex[-1]
    if not last_char.isdigit() and not (last_char >= 'a' and last_char <= 'f') and not (
            last_char >= 'A' and last_char <= 'F'):
        raise ValueError("Invalid hexadecimal digit: " + last_char)

    offset = int(last_char, 16) & 0x0F

    # Extract a 4-byte slice from the hash starting at the offset
    otp_bytes = hash[offset:offset + 4]

    # Convert the bytes to an integer
    otp_int = struct.unpack('>I', otp_bytes)[0]

    # Generate the OTP as a 6-digit number
    otp_raw = otp_int % 10 ** 6
    otp = '{:06d}'.format(otp_raw)

    return otp


def send_otp(phone_number: str, otp):
    # Compose the message
    message = f"Your OTP is: {otp}. It expires in 5 minutes. Please do not share it with anyone."

    # Send the message
    try:
        response = sms.send(message, [phone_number])
        if response['SMSMessageData']['Recipients'][0]['status'] == 'Success':
            os.write(2, b"OTP sent successfully\n")
            os.write(2, str(response).encode() + b"\n")
            return True
        else:
            # print response to stderr to avoid buffering
            os.write(2, str(response).encode() + b"\n")
            return False
    except Exception as e:
        # print error to stderr to avoid buffering
        os.write(2, str(e).encode() + b"\n")
        return False
