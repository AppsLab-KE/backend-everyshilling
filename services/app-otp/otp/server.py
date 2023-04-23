import uuid
import grpc
import logging
from . import otp
from everyshillingsproto.otp import otpserver_pb2_grpc
from everyshillingsproto.otp import otpserver_pb2


class otp_service(otpserver_pb2_grpc.OtpServiceServicer):
    def HealthCheck(self, request, context):
        health = otpserver_pb2.DefaultResponse()
        return health

    def CreateAndSendOtp(self, request, context):
        message = "otp generated successfully"
        status_code = 200
        tracking_uuid = str(uuid.uuid4())

        otp_code = otp.generate_otp()

        # Send otp to phone using africa's talking

        return otpserver_pb2.CreateAndSendOtpRes(
            message=message,
            status_code=status_code,
            tracking_uuid=tracking_uuid,
        )

    def VerifyOtp(self, request, context):
        message = "otp valid"
        status_code = 200

        return otpserver_pb2.VerifyOTPRes(
            message=message,
            status_code=status_code
        )

    def ResendOTP(self, request, context):
        return super().ResendOTP(request, context)


    async def run(self, port):
        server = grpc.aio.server()
        otpserver_pb2_grpc.add_OtpServiceServicer_to_server(self, server)
        listen_addr = f"[::]:{port}"
        server.add_insecure_port(listen_addr)
        logging.info("starting server")
        await server.start()
        await server.wait_for_termination()
