FROM ubuntu:latest


RUN apt-get update -y

RUN apt-get install git -y

RUN echo StrictHostKeyChecking no >> /etc/ssh/ssh_config

RUN apt-get install sshpass -y

RUN git config --global user.name "atmecs-classec"
RUN git config --global user.pass "atmecs123"

#clone the git repository

RUN git clone https://atmecs-classec:atmecs123@github.com/bhanu2507/goclassec.git -b dev


#RUN sshpass -p atmecs123 git clone  https://github.com/bhanu2507/goclassec.git -b dev

RUN apt-get install python-software-properties -y

RUN apt-get install curl -y

RUN curl -sL https://deb.nodesource.com/setup_7.x | bash -

RUN apt-get install nodejs -y

RUN npm update
RUN cd /goclassec && npm install -y
RUN cd /goclassec && ls
#RUN cd goclassec/

#RUN ls

RUN cd goclassec && npm install -g bower -y

RUN cd goclassec && bower install --allow-root


RUN cd goclassec && npm install gulp-cli -g -y

RUN cd goclassec && npm install gulp -D -y

RUN cd goclassec && mkdir /scripts

RUN cd goclassec && cp -a node_modules/. scripts
RUN cd goclassec/node_modules && ls
EXPOSE 2200
ENTRYPOINT node /goclassec/index.js
#RUN cd goclassec && node index.js
