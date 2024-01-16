FROM golang:alpine as builder 
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux go build -o golang-modular-software


FROM alpine
WORKDIR /app
COPY --from=builder /app /app/
EXPOSE 9000
CMD [ "./golang-modular-software" ]