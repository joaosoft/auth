############################
# STEP 1 build binary
############################
FROM golang:alpine AS builder

LABEL maintainer="João Ribeiro <joaosoft@gmail.com>"

RUN apk update && apk add --no-cache \
	curl \
	bash \
	dep \
	git

WORKDIR /go/src/auth
COPY . .

RUN dep ensure

RUN GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 go build -a -installsuffix cgo -o auth .

RUN chmod +x auth


############################
# STEP 2 run binary
############################
FROM scratch
COPY --from=builder /go/src/auth/auth .

EXPOSE 8000
ENTRYPOINT ["./auth"]