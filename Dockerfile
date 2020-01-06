FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD ./shippy-consignment-service /app/consignment-service

CMD ["./consignment-service"]