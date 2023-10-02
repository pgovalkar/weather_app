FROM golang:1.21.0

ADD . /weather
WORKDIR /weather

EXPOSE 8080
CMD ["go", "run", "main.go"]

