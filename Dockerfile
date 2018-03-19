FROM opensuse

RUN zypper in -y go go-doc

RUN mkdir -p /go/src/vkr-clmgr

ENV GOPATH=/go/src/
ENV PATH=$PATH+:/usr/bin/go

WORKDIR /go/src/
#COPY ./ /go/src/vkr-clmgr/

CMD ["go", "run", "./vkr-clmgr/cmd/vkr_clmgr/main.go"]