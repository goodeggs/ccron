FROM alpine:3.2

MAINTAINER Good Eggs <open-source@goodeggs.com>

RUN apk --update add bash curl \
  && curl -q -L https://raw.github.com/kvz/cronlock/master/cronlock -o /usr/bin/cronlock && chmod +x /usr/bin/cronlock \
  && adduser -S ccrond \
  && echo "*/15 * * * * ccrond-sync" | crontab -u root -

COPY ccrond /usr/sbin/
COPY ccrond-sync /usr/sbin/
COPY ccron /usr/bin/

CMD ["ccrond"]
