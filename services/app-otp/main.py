import os

from otp import server
import asyncio

if __name__ == '__main__':
    # write to stderr to avoid buffering
    os.write(2, b"Starting otp server...\n")
    otp_server = server.OtpService()
    asyncio.run(otp_server.run(3008))
