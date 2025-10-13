FROM alpine:latest

RUN apk --no-cache add curl netcat-openbsd

RUN adduser default --disabled-password
USER default

COPY --from=gsn_manager_service_base /go-app /go-app

EXPOSE 8080

CMD ["/go-app"]
