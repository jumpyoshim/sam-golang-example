FROM lambci/lambda:build-go1.x

LABEL LABEL maintainer="jumpyoshim <jumpyoshim@gmail.com>"

ENV GO111MODULE on
ENV GOFLAGS -mod=mod

WORKDIR /app

ADD Makefile \
  go.mod \
  go.sum \
  ./
RUN go mod download

ADD funcs funcs
