import time
import hmac
import hashlib
import struct
from Crypto.PublicKey import RSA
import os
import africastalking

africastalking_username =os.getenv("AFR_USERNAME")
africastalking_api_key = os.getenv('AFRICASTALKING_API_KEY')

africastalking.initialize(africastalking_username, africastalking_api_key)
sms = africastalking.SMS

def generate_otp() -> str:
    # Generating RSA keys
    secret = os.getenv('OTP_SECRET')

    if not secret:
        raise Exception("otp key missing")

    key = RSA.generate(2048)
    public_key = key.publickey()
    private_key = key
    # Generate a random secret key as a bite string

    # Set the steps and get the unix time
    time_step = 120
    current_time = int(time.time())

    # Get the number of time steps
    time_steps = int(current_time / time_step)

    # Pack the time steps as a big-endian byte string
    time_bytes = struct.pack('>Q', time_steps)

    secret = str.encode(secret)

    # Calculate the HMAC-SHA1 hash of the time steps using the secret key
    hash = hmac.new(secret, time_bytes, hashlib.sha1).digest()

    # Get the last 4 bits of the hash
    offset = hash[-1] & 0x0F

    # Extract a 4-byte slice from the hash starting at the offset
    otp_bytes = hash[offset:offset + 4]

    # Convert the bytes to an integer
    otp_int = struct.unpack('>I', otp_bytes)[0]

    # Generate the OTP as a 6 digit number
    otp_raw = otp_int % 10 ** 6
    otp = '{:06d}'.format(otp_raw)

    return otp

def send_otp(phone_number, otp):
    # Compose the message
    message = f"Your OTP is: {otp}. Please don't share it with anyone."

    # Send the message
    try:
        response = sms.send(message, [phone_number])
        print(response)
    except Exception as e:
        print(f"Encountered an error while sending SMS: {str(e)}")
        
phone_number = '+254738847827'

if not phone_number:
    raise Exception("Phone number missing")

# Generate the OTP
otp = generate_otp()

# Send the OTP via SMS
send_otp(phone_number, otp)        