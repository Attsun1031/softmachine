FROM alpine:3.7
RUN addgroup -S jobnetes && adduser -S -G jobnetes jobnetes
USER jobnetes
COPY jobnetes-manager /app/jobnetes-manager
ENTRYPOINT ["/app/jobnetes-manager"]