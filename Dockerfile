FROM opensuse

RUN zypper in -y go go-doc

RUN mkdir -p /go/src/clmgr-coordinator

ENV GOPATH=/go/
ENV PATH=$PATH+:/usr/bin/go

WORKDIR /go/src/
COPY ./ /go/src/clmgr-coordinator/

CMD ["go", "run", "./clmgr-coordinator/main.go"]