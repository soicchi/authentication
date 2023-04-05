FROM golang:1.20-alpine
ENV ROOT /app
WORKDIR $ROOT
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080
CMD ["go", "run", "./..."]
