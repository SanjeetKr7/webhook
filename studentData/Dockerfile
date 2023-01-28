FROM golang:1.17-alpine
RUN apk upgrade && apk add curl
WORKDIR /app
COPY ./ ./
WORKDIR /app/cmd
RUN go build -o studentData
EXPOSE 9090
CMD [ "./studentData" ]
