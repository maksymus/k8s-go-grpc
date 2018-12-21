FROM golang:1.10.4

#WORKDIR /usr/home/Projects/goplace/src/k8s-go-grpc/api-service
COPY api-service/api .
COPY proto ../proto

#RUN go get -v ./...
#RUN go install -v ./...

EXPOSE 3000

CMD [ "./api" ]