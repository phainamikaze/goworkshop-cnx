FROM scratch
# ENV GOROOT /usr/local

COPY app /

EXPOSE 8080

ENTRYPOINT ["/app"]