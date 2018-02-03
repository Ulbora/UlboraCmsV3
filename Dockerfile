FROM ubuntu

RUN apt-get update  
RUN apt-get install -y ca-certificates
VOLUME /static/templates
ADD main /main
ADD entrypoint.sh /entrypoint.sh
ADD static /static
WORKDIR /

EXPOSE 8090
ENTRYPOINT ["/entrypoint.sh"]

