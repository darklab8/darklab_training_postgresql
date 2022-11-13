FROM python:3.10

WORKDIR /install

COPY ./requirements.txt ./
RUN pip install -r requirements.txt

USER root

WORKDIR /code

COPY . .