FROM hb.test2.yunshanmeicai.com/golang:1.13.1

RUN mkdir -p /data/www/orange_message_service
RUN mkdir -p /data/logs/orange_message_service
RUN mkdir /var/log/supervisor

workdir /data/www/orange_message_service

COPY . .

RUN cp ./conf/app.json.example ./conf/app.json

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN go build -o ./orange_message_service .

ADD ./bin/supervisord.conf /etc/

RUN mkdir -p /etc/supervisor/conf.d

ADD ./bin/orange_message_service.supervisor.conf /etc/supervisor/conf.d/

ENTRYPOINT ["supervisord", "--nodaemon", "--configuration", "/etc/supervisord.conf"]
