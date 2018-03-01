FROM opensuse

RUN zypper in -y go go-doc

RUN mkdir -p /go/src/

COPY . /go/src/
ENV GOPATH=/go/src/

WORKDIR /go/src/vkr_clmgr