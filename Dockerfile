FROM golang:1.10.0

##Install main rest Golang application
RUN apt-get update -y 
ADD ./install_dep.sh /install.sh
RUN bash /install.sh

ADD ./app /go/src/bitbucket.org/lyledean/pong/app
ADD ./Gopkg.toml /go/src/bitbucket.org/lyledean/pong/app
#COPY ./Gopkg.lock /go/src/bitbucket.org/lyledean/pong/app/Gopkg.lock

#WORKDIR /$GOPATH/src
#RUN go get k8s.io/client-go/...
#RUN go get k8s.io/client-go/...
WORKDIR /go/src/bitbucket.org/lyledean/pong/app
RUN dep ensure 
#WORKDIR $GOPATH/src/k8s.io/client-go
#RUN git checkout v7.0.0 # replace v9.0.0 with the required version
# cd 1.5 # only necessary with 1.5 and 1.4 clients.
#RUN godep restore ./...
WORKDIR /go/src/bitbucket.org/lyledean/pong/app
#RUN go get -u k8s.io/apimachinery/...
RUN go get "github.com/golang/glog"
RUN go get "github.com/google/gofuzz"
RUN go get "github.com/gorilla/mux"
RUN go build -o main .
CMD ["/go/src/bitbucket.org/lyledean/pong/app/main"]