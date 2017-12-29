FROM alpine

RUN mkdir -p /opt/resource

ADD check/check /opt/resource/check
ADD in/in /opt/resource/in
ADD out/out /opt/resource/out

CMD ls -laR /opt/resource