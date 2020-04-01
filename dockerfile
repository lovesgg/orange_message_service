FROM hb.test2.yunshanmeicai.com/golang:1.13.1

RUN mkdir -p /data/www/mj_lobster_go_service
RUN mkdir -p /data/logs/mj_lobster_go_service
RUN mkdir /var/log/supervisor

workdir /data/www/mj_lobster_go_service

COPY . .

RUN cp ./conf/app.json.example ./conf/app.json

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN go build -o ./mj_lobster_go_service .

ADD ./bin/supervisord.conf /etc/

RUN mkdir -p /etc/supervisor/conf.d

ADD ./bin/mj_lobster_go_service.supervisor.conf /etc/supervisor/conf.d/

ENTRYPOINT ["supervisord", "--nodaemon", "--configuration", "/etc/supervisord.conf"]
