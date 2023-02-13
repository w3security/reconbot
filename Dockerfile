FROM ubuntu:latest

RUN apt update -y && apt install -y  --no-install-recommends \
    build-essential \
    cmake \
    firefox \
    gcc \
    git \
    libpcap0.8-dev \
    libpcap-dev \
    golang

RUN git clone https://github.com/w3security/reconbot.git \
    && cd reconbot \
    && go build

CMD ["bash"]
