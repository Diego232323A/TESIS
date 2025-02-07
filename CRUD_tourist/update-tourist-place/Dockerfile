FROM python:3.10

WORKDIR /app

COPY . /app

RUN pip install --no-cache-dir -r requirements.txt

EXPOSE 5002

CMD ["uvicorn", "src.main:app", "--host", "0.0.0.0", "--port", "5002", "--reload"]
