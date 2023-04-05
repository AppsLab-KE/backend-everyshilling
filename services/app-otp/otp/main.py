import math, random



def generateOTP(): # this function generates otp
    digits ='0123456789'
    otp =''
    for i in range(6):
        otp+=digits[math.floor(random.random()* 10)]
    
    return otp
    
    
if __name__=='__main__':
    otp =generateOTP()
    print('Your OTP is', generateOTP())  
    inputOTP = input("Enter your OTP >>: ")
    print(inputOTP)
    if  inputOTP == otp:
            print("OTP is Verified")
    else:
        print("Please Check your OTP again")
