# ---------- build stage ----------
FROM python:3.11-slim AS builder

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    PIP_NO_CACHE_DIR=1 \
    PIP_DISABLE_PIP_VERSION_CHECK=1

WORKDIR /wheels

COPY requirements.txt .
RUN pip wheel --no-deps --wheel-dir /wheels -r requirements.txt

# ---------- runtime stage ----------
FROM python:3.11-alpine

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1

# 安装运行时依赖 + 时区数据
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

# 创建非 root 用户
RUN addgroup -g 1000 appgroup && \
    adduser -D -u 1000 -G appgroup appuser

WORKDIR /app

# 把 wheels 和代码拷进来
COPY --from=builder /wheels /wheels
RUN pip install --no-cache /wheels/* && \
    rm -rf /wheels

COPY --chown=appuser:appgroup . .

USER appuser

# 默认端口
EXPOSE 5000

# gunicorn 启动（推荐生产用）
# 如果入口文件是 app.py，application 实例叫 app，则命令如下：
CMD ["gunicorn", "-b", "0.0.0.0:5000", "--access-logfile", "-", "app:app"]

# 如果只想用 Flask 自带开发服务器，把上面 CMD 换成：
# CMD ["python", "app.py"]