FROM python:2.7-alpine

RUN mkdir -p /app
WORKDIR /app
ADD requirements.txt /app
RUN pip install -r requirements.txt
ADD app.py /app

EXPOSE "5000"
CMD ["python", "app.py"]
