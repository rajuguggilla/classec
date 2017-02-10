FROM ubuntu:latest
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
#RUN  sshpass -p git1 git clone ssh://git1@110.110.110.164/classec  -b master
ENV GOPATH="/goclassec"
RUN echo $GOPATH
RUN go get "github.com/Azure/azure-sdk-for-go/arm/compute"
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/jinzhu/gorm"
RUN go get "github.com/verdverm/frisby"
RUN go get "github.com/mozillazg/request"
RUN go get "github.com/bitly/go-simplejson"
RUN go get "git.openstack.org/openstack/golang-client.git/openstack"
RUN go get "github.com/dgrijalva/jwt-go"
RUN go get "github.com/vmware/govmomi"
RUN go get "github.com/Azure/go-autorest/autorest/azure"

RUN cp /goclassec/Dockerfiles/ca-certificates.crt /etc/ssl/certs/

#RUN apt-get install docker -y
#RUN echo 110.110.110.5 controller >> /etc/hosts
#RUN echo 110.110.110.5 controller >> my_hosts
#RUN cat  /etc/hosts >> my_hosts
#RUN unshare --mount
EXPOSE 9009
ENTRYPOINT echo "110.110.110.5 controller" >> /etc/hosts &&  go run goclassec/src/gclassec/server/main.go
#RUN go
