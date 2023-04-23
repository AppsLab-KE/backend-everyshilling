from otp import server
import asyncio
import logging

if __name__ == '__main__':
    otp_server = server.otp_service()
    asyncio.run(otp_server.run(3008))