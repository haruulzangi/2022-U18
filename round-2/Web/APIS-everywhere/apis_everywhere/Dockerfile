FROM python:3.8

COPY requirements.txt ./requirements.txt
COPY api.py ./api.py
COPY todo.db ./todo.db
COPY __pycache__ ./__pycache__

RUN pip install -r requirements.txt

EXPOSE 5000

CMD ["python3", "./api.py"]
