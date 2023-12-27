
ARG BASE=golang:1.21-alpine
FROM $BASE as base

COPY src /build/

RUN cd /build && go build -o /awktest

