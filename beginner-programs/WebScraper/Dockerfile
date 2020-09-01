FROM golang:1.15.0

RUN go get github.com/gocolly/colly

COPY . . 

# Build the application
RUN go build -o main .

# Command to run the executable
CMD ["./main"]