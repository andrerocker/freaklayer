FROM asemenov/squeeze
MAINTAINER Freaklayer Team: 0.1

RUN echo "deb http://mirror-debian.locaweb.com.br/debian/ squeeze main non-free" > /etc/apt/sources.list
RUN echo "deb http://mirror-debian.locaweb.com.br/debian-security/ squeeze/updates main" >> /etc/apt/sources.list

RUN apt-get update
RUN apt-get install -y --force-yes --no-install-recommends ruby1.9.1 build-essential ca-certificates curl
RUN apt-get install -y --force-yes --no-install-recommends mysql-client libmysqlclient-dev

RUN ln -s /usr/bin/ruby1.9.1 /usr/bin/ruby
