FROM python:3.11-slim

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    DISABLE_MODEL_SOURCE_CHECK=True

WORKDIR /app

COPY requirements.txt .
RUN pip install \
    -r requirements.txt \
    -i https://pypi.tuna.tsinghua.edu.cn/simple \
    --timeout 120 \
    --retries 5

COPY . .

EXPOSE 5000
CMD ["python", "app.py"]
