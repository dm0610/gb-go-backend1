FROM golang:latest
RUN  adduser \
   --system \
   --shell /bin/bash \
   --disabled-password \
   --home /app goserver && \
	chown -R goserver /app
WORKDIR /app
COPY *.go ./
RUN go mod init goserver
RUN go mod tidy
RUN go build -o goserver .
USER goserver
EXPOSE 8010
CMD ["./goserver"]