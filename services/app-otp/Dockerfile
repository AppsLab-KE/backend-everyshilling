FROM python:3.9.7-alpine

WORKDIR /app-otp/

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . /app-otp/
RUN chmod u+x main.py

CMD ["python" , "otp/main.py"]