FROM golang:1.22

WORKDIR /slijterij

COPY go.* ./

RUN go mod download

COPY . .

EXPOSE 8080

CMD ["go", "run", "."]
