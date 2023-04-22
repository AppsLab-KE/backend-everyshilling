from otp import server
import asyncio
import logging

if __name__ == '__main__':
    otp_server = server.otp_service()
    logging.info("otp server running on port ")
    asyncio.run(otp_server.run(3007))