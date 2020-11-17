FROM lambci/lambda:build-go1.x

LABEL LABEL maintainer="jumpyoshim <jumpyoshim@gmail.com>"

ENV GO111MODULE on
ENV GOFLAGS -mod=mod

WORKDIR /app

RUN curl -sL https://rpm.nodesource.com/setup_12.x | bash - \
  && yum install -y nodejs \
  && yum clean all
RUN npm install npm@latest -g

ADD cdk cdk
RUN cd cdk && npm install

ADD Makefile \
    go.mod \
    go.sum \
    ./
RUN go mod download

ADD funcs funcs
