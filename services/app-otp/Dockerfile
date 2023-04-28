FROM python:3.9.7-alpine

WORKDIR /app
RUN apk  add git gcc openssl-dev

RUN pip install -U \
    pip \
    setuptools \
    wheel

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

CMD ["python", "main.py"]
