# FROM python:3.9.7-alpine

# WORKDIR /app-otp/otp

# COPY requirements.txt requirements.txt
# RUN pip install --no-cache-dir -r requirements.txt

# COPY . .

# CMD ["python3" , "main.py"]

# Use an official Python runtime as a base image

FROM python:3.9.7-alpine

WORKDIR /app

COPY otp/requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY otp .

EXPOSE 3000
CMD ["python", "main.py"]
