FROM python:2.7-alpine

ADD requirements.txt /code
WORKDIR /code
RUN pip install -r requirements.txt
ADD . /code/

EXPOSE "5000"
CMD ["python", "app.py"]
