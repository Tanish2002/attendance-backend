FROM golang:1.17.8-bullseye

RUN apt-get update && apt-get install -y libdlib-dev libblas-dev libatlas-base-dev liblapack-dev libjpeg62-turbo-dev
#RUN apk add jpeg libjpeg libjpeg-turbo jpeg-dev gcc

#ENV LOG_LEVEL 'info'
#ENV READONLY 'false'
#endpoint: http, lambda
#ENV ENDPOINT 'lambda'

WORKDIR /app
COPY . /app

RUN go env -w GOPROXY=direct
#RUN ls /app

#ADD /go.sum /app/
#ADD configuration /app/

#RUN go mod tidy -compat=1.17
#RUN go mod download -x
#RUN go get ...

#ADD /main.go /app/
#ADD /handlers /app/handlers

# RUN go test -v -p 1
RUN GOOS=linux GOARCH=amd64 go build -mod vendor -o /bin/attendance-backend

ENTRYPOINT [ "/bin/attendance-backend"]

#FROM alpine:3.13
#ADD /startup.sh /
#COPY --from=BUILD /bin/attendance-backend /bin/
#CMD [ "/bin/attendance-backend" ]