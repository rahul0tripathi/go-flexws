FROM alpine:3.14

RUN apk add --no-cache python python-dev gcc build-base musl-dev libffi-dev openssl-dev && \
	python -m ensurepip && \
	pip install --upgrade pip && \
	pip install autobahntestsuite

VOLUME /config
VOLUME /report

CMD ["/usr/bin/wstest", "--mode", "fuzzingclient", "--spec", "/config/fuzzingclient.json"]
