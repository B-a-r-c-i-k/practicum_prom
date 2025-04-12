
FROM golang:1.19

WORKDIR /app

COPY . .
CMD ["go", "run", "./logger.go"]

EXPOSE 80