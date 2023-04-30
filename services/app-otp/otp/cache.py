import os
from datetime import datetime

import redis

REDIS_HOST = os.getenv("REDIS_HOST")
REDIS_PORT = int(os.getenv("REDIS_PORT"))

r = redis.Redis(host=REDIS_HOST, port=REDIS_PORT, db=0)


def save_otp(phone_number, tracking_uuid, otp):
    """
    Save otp to redis cache
    :param phone_number:  phone number
    :param tracking_uuid:  tracking uuid
    :param otp:  otp
    :return:  None
    """
    tracking_key = f"uuid:{tracking_uuid}_otp"
    phone_key = f"phone:{phone_number}_otp"
    time_stamp = datetime.now().timestamp()

    phone_otp_dict = {
        "phone_number": phone_number,
        "otp": otp,
        "time_stamp_unix": time_stamp
    }

    otp_time_stamp_dict = {
        "otp": otp,
        "time_stamp_unix": time_stamp
    }

    r.hset(tracking_key, mapping=phone_otp_dict)
    r.hset(phone_key, otp_time_stamp_dict)
    r.expire(phone_key, 300)
    r.expire(tracking_key, 60 * 10)


def get_otp(phone_number):
    """
    Get otp from redis cache
    :param phone_number:  phone number
    :return:  otp
    """
    phone_key = f"phone:{phone_number}_otp"
    otp = r.hgetall(phone_key)
    return otp


def get_phone_number(tracking_uuid):
    """
    Get phone number from redis cache
    :param tracking_uuid:  tracking uuid
    :return:  phone number, timestamp and otp
    """
    tracking_key = f"uuid:{tracking_uuid}_otp"
    phone_number = r.hgetall(tracking_key)
    return phone_number


def delete_otp_phone(phone_number, tracking_uuid):
    """
    Delete otp and phone number from redis cache
    :param phone_number:
    :param tracking_uuid:
    :return:  None
    """
    tracking_key = f"uuid:{tracking_uuid}_otp"
    phone_key = f"phone:{phone_number}_otp"

    r.delete(tracking_key)
    r.delete(phone_key)


def delete_otp_tracking(phone_number, tracking_uuid):
    """
    Delete otp and phone number from redis cache
    :param phone_number:
    :param tracking_uuid:
    :return:
    """
    tracking_key = f"uuid:{tracking_uuid}_otp"
    phone_key = f"phone:{phone_number}_otp"

    r.delete(tracking_key)
    r.delete(phone_key)
