
#base image to be used
FROM golang:1.18.3-alpine as builder
#select app folder
WORKDIR /UserService_BeMyPet
#copy go file to install packages
#install dependencies
COPY go.* ./
RUN go mod download
# Copy local code to the container image.
COPY . ./

#copy source code to image, ignore node_modules because we already installed them
#COPY . /react-frontend

# Build the binary.
RUN go build -o main .
#set environment
ENV port=8080
#expose port so we can access the app
EXPOSE 8080
#command to start the app , "PORT=0.0.0.0:8080"
CMD ["go", "run", "main.go"]