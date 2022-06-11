FROM python:3.8

COPY requirements.txt ./requirements.txt
COPY app.py ./app.py
COPY templates/login.html ./templates/login.html

RUN pip install -r requirements.txt

CMD ["python3", "./app.py"]
