FROM alpine

LABEL maintainer="Harry Bagdi<harry.bagdi@gmail.com>"

COPY rand /bin/rand

ENTRYPOINT /bin/sh
