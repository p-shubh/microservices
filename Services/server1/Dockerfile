FROM alpine:3.18

# under the same container it creates a working directory of appServer1 
WORKDIR /appServer1

# copy appServer1's that Specifies the module path and the dependency requirements for this Go project
COPY go.mod .

# copy appServer1's that Contains the checksums for the exact contents of each dependency at the time it is added to your module
COPY go.sum .

# Downloads the dependencies defined in the go. mod file and stores them in the module cache
RUN go mod download

# it copies all the source code in the current directory that is /appServer1
COPY . .

EXPOSE 90901

CMD [ "./appServer1" ]