FROM alpine

RUN mkdir -p /opt/resource && \
    apk -Uuv add ca-certificates && \
	update-ca-certificates && \
    rm /var/cache/apk/*

ADD check/check /opt/resource/check
ADD in/in /opt/resource/in
ADD out/out /opt/resource/out

CMD ls -laR /opt/resource