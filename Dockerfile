
ARG BASE=golang:1.21-alpine
FROM $BASE as base

ENV CGO_ENABLED=0
ENV GOOS=linux

COPY src /build/

RUN cd /build && ls -ltR &&  go build -o /awktest

FROM alpine

COPY --from=base /awktest /bin/awktest

CMD ["awktest"]
