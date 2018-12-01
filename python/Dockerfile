FROM python:3-slim

RUN apt-get update -y && \
    python3 -m pip install -U discord.py && \
    python3 -m pip install aiohttp websockets 

VOLUME /app
WORKDIR /app

COPY *.py /app

CMD ["python --version"]
