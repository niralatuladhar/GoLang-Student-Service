FROM golang

COPY . $GOPATH/src/github.com/ntuladhar/student-webservice
WORKDIR $GOPATH/src/github.com/ntuladhar/student-webservice

# we need to install dependencies before building
RUN apt-get update && apt-get install -y pkg-config libxml2-dev

RUN go get github.com/codegangsta/martini
RUN go get github.com/jmoiron/sqlx
RUN go get github.com/lib/pq

# build the project
#RUN go build server.go

# open the default port
EXPOSE 3000

# run using defaults, except the XML url, so it points to travel's qa server
#CMD ["go run server.go"]
RUN go run ./server.go

