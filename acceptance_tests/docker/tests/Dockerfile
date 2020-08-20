FROM python:3.8.5-alpine

WORKDIR /app

RUN apk add --no-cache vim curl bash

COPY requirements.txt .

RUN pip install -r requirements

CMD tail -F /dev/null