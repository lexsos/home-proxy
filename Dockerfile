FROM alpine:latest
COPY ./build/hproxy /bin/hproxy

ENV APP_ADDR=":8080"
ENV APP_LOG_LEVEL="info"
ENV APP_LOG_FORMAT="text"
ENV APP_CONF_FILE="/config.json"

CMD /bin/hproxy -proxy-addr $APP_ADDR -log-level $APP_LOG_LEVEL -log-format $APP_LOG_FORMAT -auth-file $APP_CONF_FILE
