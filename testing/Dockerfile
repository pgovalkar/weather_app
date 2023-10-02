FROM centos:centos7

RUN mkdir /testing

# Set execute permission on the script on your host machine
COPY api_testing.sh /testing/
RUN chmod +x /testing/api_testing.sh

WORKDIR /testing

ENTRYPOINT ["sh", "./api_testing.sh",";","tail", "-f", "/dev/null"]
