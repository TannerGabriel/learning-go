FROM golang:latest AS builder

RUN go get github.com/gorilla/mux 
RUN go get go.mongodb.org/mongo-driver/bson 
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get github.com/joho/godotenv

COPY . . 

# Build the application
RUN go build -o main .

ENV HOST=mongo

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]