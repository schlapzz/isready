FROM registry.puzzle.ch/bit/golang-alpine:1.14-alpine AS build-stage
WORKDIR /src/

COPY go.mod go.sum ./
COPY . .
RUN go build -o isready ./main.go

#------------------

FROM scratch
COPY --from=build-stage /src/isready .
ENTRYPOINT [ "./isready" ]
