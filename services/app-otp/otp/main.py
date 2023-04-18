import time
import hmac
import hashlib
import struct
from Crypto.PublicKey import RSA
from Crypto.Cipher import PKCS1_OAEP
# import africastalking

# Generating RSA keys

key = RSA.generate(2048)

public_key = key.publickey()
private_key = key

# Generate a random secret key as a bite string 
secret = b'MysecretKey'

# Set the steps and get the unix time
time_step=300
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
otp_raw= otp_int % 10**6
otp = '{:06d}'.format(otp_raw)

print('Your OTP is :',otp)

# Encryption of the OTP using the public key
cipher = PKCS1_OAEP.new(public_key)
encrypted_otp = cipher.encrypt(otp.encode())

# print(encrypted_otp)

# Get the user's input for the OTP
user_otp = input('Enter your OTP: ')

# Decryption of the OTP using the private key
cipher = PKCS1_OAEP.new(private_key)
decrypted_otp = cipher.decrypt(encrypted_otp)

# Verifying the OTP
if int(user_otp) == int(decrypted_otp.decode()):
    print('OTP is valid')
else:
    try:
        num = int('abc')
    except ValueError:
      print('OTP is invalid')