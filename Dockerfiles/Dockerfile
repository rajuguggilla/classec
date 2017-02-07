#Base ubuntu image
FROM ubuntu:latest
#Install Basic Commands
#RUN rm /bin/sh && ln -s /bin/bash /bin/sh
RUN apt-get update -y 
RUN apt-get -y upgrade
RUN apt-get install curl -y
RUN curl -O https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz
RUN tar -C /usr/local -xvzf go1.7.linux-amd64.tar.gz
#RUN mkdir -p ~/projects/{bin,pkg,src}

RUN apt-get install vim -y
#RUN mv go /usr/local
RUN exec bash
ENV PATH=$PATH:/usr/local/go/bin
#RUN export PATH=$PATH:/usr/local/go/bin


ENV GOBIN="$HOME/projects/bin"
ENV GOPATH="$HOME/projects/src"

RUN apt-get install -y git
RUN echo StrictHostKeyChecking no >> /etc/ssh/ssh_config
RUN apt-get install sshpass -y
#CMD cat 110.110.110.5 controller >> /etc/hosts
RUN sshpass -p git1 git clone ssh://git1@110.110.110.164/goclassec  -b master
ENV GOPATH="/goclassec"
RUN echo $GOPATH
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/jinzhu/gorm"
RUN go get "github.com/verdverm/frisby"
RUN go get "git.openstack.org/openstack/golang-client.git/openstack"
RUN go get "github.com/dgrijalva/jwt-go"
RUN go get "github.com/vmware/govmomi"
RUN go get "github.com/Azure/go-autorest/autorest/azure"
RUN chmod 755 /goclassec/src/gclassec/server/main.go
#RUN echo 110.110.110.5 controller >> /etc/hosts
EXPOSE 9009
#ENTRYPOINT [echo  "110.110.110.5 controller" >> /etc/hosts && go run goclassec/src/gclassec/server/main.go]
#CMD [ "go run /goclassec/src/gclassec/server/main.go"]
ENTRYPOINT go run /goclassec/src/gclassec/server/main.go
