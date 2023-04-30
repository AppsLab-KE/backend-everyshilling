import uuid
from datetime import datetime

import grpc
import logging
from . import otp
from everyshillingsproto.otp import otp_pb2
from everyshillingsproto.otp import otpserver_pb2_grpc
from everyshillingsproto.otp import otpserver_pb2
from . import cache


class OtpService(otpserver_pb2_grpc.OtpServiceServicer):
    def HealthCheck(self, request, context):
        health = otpserver_pb2.DefaultResponse()
        return health

    def CreateAndSendOtp(self, request, context):
        message = "otp generated successfully"
        status_code = 200
        tracking_uuid = str(uuid.uuid4())

        # check if phone number exists in cache
        phone_number = request.phone_number
        otp_code = cache.get_otp(phone_number)

        if otp_code:
            return otp_pb2.CreateAndSendOtpRes(
                message="Otp already exists. Please wait for 5 minutes before generating a new one",
                status_code=409,
                tracking_uuid="",
            )

        otp_code = otp.generate_otp()
        # Send otp to phone using africa's talking
        #
        otp.send_otp(phone_number, otp_code)

        # save otp to cache
        cache.save_otp(phone_number, tracking_uuid, otp_code)
        return otp_pb2.CreateAndSendOtpRes(
            message=message,
            status_code=status_code,
            tracking_uuid=tracking_uuid,
        )

    def VerifyOtp(self, request, context):
        message = "Otp Validated Successfully"
        status_code = 200

        otp = request.otp_code
        tracking_uuid = request.tracking_uuid

        phone_number = cache.get_phone_number(tracking_uuid)
        if not phone_number:
            return otp_pb2.VerifyOTPRes(
                message="Incorrect/expired otp. Please generate a new one",
                status_code=401
            )

        cached_otp = cache.get_otp(phone_number)
        if not cached_otp:
            return otp_pb2.VerifyOTPRes(
                message="This otp has expired. Please generate a new one",
                status_code=401
            )

        time_stamp = datetime.now().timestamp()
        time_diff = time_stamp - float(cached_otp["time_stamp_unix"])

        if time_diff > 300:
            return otp_pb2.VerifyOTPRes(
                message="This otp has expired. Please generate a new one",
                status_code=401
            )

        if otp != cached_otp["otp"]:
            return otp_pb2.VerifyOTPRes(
                message="Incorrect otp. Please try again",
                status_code=401
            )

        return otp_pb2.VerifyOTPRes(
            message=message,
            status_code=status_code
        )

    def ResendOTP(self, request, context):
        time_stamp = datetime.now().timestamp()
        tracking_uuid = request.tracking_id

        user = cache.get_phone_number(tracking_uuid)
        if not user:
            return otp_pb2.ResendOTPRes(
                message="Otp session expired. Please start again",
                status_code=401,
                tracking_uuid=""
            )

        phone_number = user["phone_number"]
        cached_otp =  user["otp"]

        time_diff = time_stamp - float(user["time_stamp_unix"])

        if time_diff > 600:
            return otp_pb2.ResendOTPRes(
                message="Otp session expired. Please start again",
                status_code=401,
                tracking_uuid=""
            )

        # do not resend otp if it has not been 5 minutes
        if time_diff < 300:
            return otp_pb2.ResendOTPRes(
                message="Please wait for 5 minutes before generating a new otp",
                status_code=409,
                tracking_uuid=""
            )

        # refresh otp cache
        cache.save_otp(phone_number, tracking_uuid, cached_otp)

        otp.send_otp(phone_number, cached_otp)

        return otp_pb2.ResendOTPRes(
            message="Successfully resent otp",
            status_code=200,
            tracking_uuid=tracking_uuid
        )

    async def run(self, port):
        server = grpc.aio.server()
        otpserver_pb2_grpc.add_OtpServiceServicer_to_server(self, server)
        listen_addr = f"[::]:{port}"
        server.add_insecure_port(listen_addr)
        logging.info("starting server")
        await server.start()
        await server.wait_for_termination()
