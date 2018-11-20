FROM ubuntu:16.04

RUN rm -rf /app && mkdir /app && mkdir /kdata
COPY main /app/kask
WORKDIR /app

ENTRYPOINT ["/app/kask"]