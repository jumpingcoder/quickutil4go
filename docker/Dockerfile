FROM ubuntu:16.04

# 安装依赖
RUN apt update -y \
    && apt install -y locales \
    && rm -rf /var/lib/apt/lists/* \
    && localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8 \
    && apt update -y \
    && apt install -y fonts-wqy-zenhei fonts-wqy-microhei \

ENV LANG en_US.utf8

COPY lib/time/zoneinfo.zip /usr/local/go/lib/time/

COPY ./quickutil4go /opt/quickutil4go/
COPY ./quickutil4go/views /opt/quickutil4go/

RUN chmod 0777 -R /opt/quickutil4go/quickutil4go

WORKDIR /opt/quickutil4go

CMD [ "./quickutil4go" ]