FROM golang:1.15

COPY . .

EXPOSE 8080

CMD ["make", "build"]