import time
import hmac
import hashlib
import struct
from Crypto.PublicKey import RSA
import os


# import africastalking


def generate_otp() -> str:
    # Generating RSA keys
    secret = os.getenv("OTP_SECRET")

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
