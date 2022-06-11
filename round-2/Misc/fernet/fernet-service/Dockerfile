FROM python:3.8-slim-buster
COPY . /opt/app
WORKDIR /opt/app
RUN  pip3 install -r requirements.txt
CMD [ "python3","server.py"]
