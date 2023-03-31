import math, random
from config import TWILIO_ACCOUNT_SID, TWILIO_AUTH_TOKEN
# from twilio.rest import Client


def generateOTP(): # this function generates otp
    digits ='0123456789'
    otp =''
    for i in range(6):
        otp+=digits[math.floor(random.random()* 10)]
    
    return otp
    
    # # Twilio accounts credentials
    # account_sid = TWILIO_ACCOUNT_SID
    # auth_token = TWILIO_AUTH_TOKEN
    
    # # The phone number you want to send the SMS to
    # to_number = '+254738847827'
    
    # # The Twilio phone number that will be used to send the SMS
    # from_number = '+15855586532'
    
    # # Create a Twilio client
    # client = Client(account_sid, auth_token)
    
    # # Send the SMS with the OTP
    # message = client.messages.create(
    #     body=f'Your OTP is: {otp}',
    #     from_=from_number,
    #     to=to_number
    # )
    
if __name__=='__main__':
    otp =generateOTP()
    print('Your OTP is', generateOTP())  
    inputOTP = input("Enter your OTP >>: ")
    print(inputOTP)
    if  inputOTP == otp:
            print("OTP is Verified")
    else:
        print("Please Check your OTP again")
