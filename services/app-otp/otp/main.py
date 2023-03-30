import math, random

def generateOTP():
    digits ='0123456789'
    otp =''
    for i in range(6):
        otp+=digits[math.floor(random.random()* 10)]
    return otp    
if __name__=='__main__':
    print('Your OTP is', generateOTP())