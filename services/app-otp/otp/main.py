import time
import hmac
import hashlib
import struct

# Generate a random secret key as a bite string
secret = b'MysecretKey'

# Set the steps and get the unix time
time_step=120
current_time =int(time.time())

# Get the number of time steps
time_steps = int(current_time / time_step)

# Pack the time steps as a big-endian byte string
time_bytes = struct.pack('>Q', time_steps)

# Calculate the HMAC-SHA1 hash of the time steps using the secret key
hash = hmac.new(secret, time_bytes, hashlib.sha1).digest()

# Get the last 4 bits of the hash
offset = hash[-1] & 0x0F

# Extract a 4-byte slice from the hash starting at the offset
otp_bytes = hash[offset:offset+4]

# Convert the bytes to an integer
otp_int = struct.unpack('>I', otp_bytes)[0]

# Generate the OTP as a 6 digit number
otp = otp_int % 10**6

print('Your OTP is :',otp)

# Get the user's input for the OTP
user_otp = input('Enter your OTP: ')

# Verifying the OTP
if int(user_otp) == otp:
    print('OTP is valid')
else:
    try:
        num = int('abc')
    except ValueError:
      print('OTP is invalid')