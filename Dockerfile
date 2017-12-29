FROM alpine

RUN mkdir -p /opt/resource

ADD check/check /opt/resource/check
ADD in/in /opt/resource/in

CMD ls -laR /opt/resource