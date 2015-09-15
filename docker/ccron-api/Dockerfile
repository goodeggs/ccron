FROM alpine:3.2

MAINTAINER Good Eggs <open-source@goodeggs.com>

EXPOSE 5000

RUN apk --update add bash

COPY ccron-api /usr/sbin/

CMD ["ccron-api"]
