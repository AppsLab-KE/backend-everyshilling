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