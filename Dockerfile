FROM golang:alpine
WORKDIR .
COPY main .
COPY conf/app.ini ./conf/
COPY templates/*.html ./templates/
EXPOSE 8088
CMD ["./main"]