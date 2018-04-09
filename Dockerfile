FROM opensuse

# getting out sourcecode
##############################
RUN zypper in -y go go-doc && zypper in -y make

RUN mkdir -p /go/src/myproj.com/clmgr-coordinator

ENV GOPATH=/go/
ENV GOROOT=/usr/lib64/go/1.9/
ENV PATH=$PATH+:/usr/bin/go

WORKDIR /go/src/myproj.com/clmgr-coordinator
ADD ./ /go/src/myproj.com/clmgr-coordinator/
##############################

# installing corosync
##############################

##############################

CMD ["make"]