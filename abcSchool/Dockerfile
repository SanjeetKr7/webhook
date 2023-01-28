FROM golang:1.17-alpine
WORKDIR /app
COPY ./ ./
WORKDIR /app/cmd
RUN go build -o abcSchool
EXPOSE 9091
CMD [ "./abcSchool" ]