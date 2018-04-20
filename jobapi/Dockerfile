FROM alpine:3.7
RUN addgroup -S jobnetes && adduser -S -G jobnetes jobnetes
USER jobnetes
COPY jobapi-server /app/jobapi-server
EXPOSE 50051
ENTRYPOINT ["/app/jobapi-server"]