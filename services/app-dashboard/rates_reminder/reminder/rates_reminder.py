import africastalking
import datetime
import os


africastalking_username = os.getenv('AFR_USERNAME')
africastalking_api_key = os.getenv('AFRICASTALKING_API_KEY')

print('USName', africastalking_username)
print('Key', africastalking_api_key)
africastalking.initialize(africastalking_username, africastalking_api_key)
sms = africastalking.SMS


def send_reminder(admin_phone):
    # Get the current date and time
    current_time = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")

    # Compose the message
    message = f"Hello, Kindly  update the exchange rates. Current time is: {current_time}"

    # Send the SMS
    try:
        response = sms.send(message, [admin_phone])
        print("Reminder sent successfully!")
    except Exception as e:
        print(f"An error occurred while sending the reminder: {str(e)}")


admin_phone_number = "+254701847888"
send_reminder(admin_phone_number)
