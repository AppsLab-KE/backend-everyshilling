FROM python:3.9.7-alpine

WORKDIR /app-otp

COPY requirements.txt /app-otp/
RUN pip install --no-cache-dir -r requirements.txt

COPY . /app-otp/

CMD ["python" , "otp/main.py"]