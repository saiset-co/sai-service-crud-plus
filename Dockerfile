# Build stage
FROM golang as BUILD

WORKDIR /src/

COPY ./ /src/

RUN go build -o sai-service-crud-plus-bin -buildvcs=false

FROM ubuntu

WORKDIR /srv

# Copy binary from build stage
COPY --from=BUILD /src/sai-service-crud-plus-bin /srv/sai-service-crud-plus-bin

# Copy other files
COPY ./config.yml /srv/config.yml

RUN chmod +x /srv/sai-service-crud-plus-bin

# Set command to run your binary
CMD /srv/sai-service-crud-plus-bin start

EXPOSE 6080
