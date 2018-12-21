FROM golang:1.10.4

#WORKDIR /usr/home/Projects/goplace/src/k8s-go-grpc/gcd-service/
COPY gcd-service/gcd .
COPY proto ../proto

#RUN go get -v ./...
#RUN go install -v ./...

EXPOSE 3000

CMD [ "./gcd" ]