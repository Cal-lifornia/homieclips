FROM ubuntu:latest
LABEL authors="whobson"

ENTRYPOINT ["top", "-b"]