FROM golang:1.20

# Enviroment variable
WORKDIR /usr/src/goscrape

RUN go install github.com/cosmtrek/air@latest

#Copying files to work directory
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .

# Run and expose the server on port 3000
EXPOSE 8080