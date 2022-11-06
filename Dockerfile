# Parent image = docker apline version
# https://docs.docker.com/develop/develop-images/baseimages/
FROM golang:alpine

#RUN apk add git
#RUN git clone https://github.com/yuuwe-n/RateMyProfessors.help.git

COPY . /app
# subsequent commands are at /rmp directory
# use / to refer to root
WORKDIR /app

RUN go mod download
RUN go mod vendor
RUN go mod verify
# Don't do go build -o /rmp, creates the binary at root
RUN go build 

EXPOSE 8080

# CMD needs to be full path or smth
CMD [ "/app/rmp" ]
