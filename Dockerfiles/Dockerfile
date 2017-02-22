#Base ubuntu image
FROM ubuntu:latest
#Install Basic Commands
#RUN rm /bin/sh && ln -s /bin/bash /bin/sh
RUN     apt-get update -y  && \
        apt-get -y upgrade  && \
        apt-get install curl -y  && \
        curl -O https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz && \
        tar -C /usr/local -xvzf go1.7.linux-amd64.tar.gz && \
        apt-get install vim -y && \
        exec bash

ENV PATH=$PATH:/usr/local/go/bin
#RUN export PATH=$PATH:/usr/local/go/bin
ENV GOBIN="$HOME/projects/bin"
ENV GOPATH="$HOME/projects/src"

RUN      apt-get install -y git && \
         echo StrictHostKeyChecking no >> /etc/ssh/ssh_config && \
         apt-get install sshpass -y --no-install-recommends && \
         sshpass -p git1 git clone ssh://git1@110.110.110.164/goclassec  -b master

ENV GOPATH="/goclassec"

RUN echo $GOPATH && \
    cp /goclassec/src/gclassec/conf/dependencies.env    /dependencies.sh && \
    chmod 755 /dependencies.sh && \
    bash dependencies.sh && \
    chmod 755 /goclassec/src/gclassec/server/main.go



EXPOSE 9009
ENTRYPOINT go run /goclassec/src/gclassec/server/main.go
