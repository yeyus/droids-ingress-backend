FROM golang

ADD . /go/src/github.com/yeyus/droids-ingress-backend

RUN go install github.com/yeyus/droids-ingress-backend

ENTRYPOINT /go/bin/droids-ingress-backend

ENV HTTP_SERVE_PORT 80
ENV HTTP_BASE_DIR /go/src/github.com/yeyus/droids-ingress-backend/

EXPOSE 80