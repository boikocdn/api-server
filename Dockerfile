FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY .ouput /cdnserver

RUN ["chmod","-R", "777", "/cdnserver"]

CMD [ "sh", "/cdnserver/run.sh"]