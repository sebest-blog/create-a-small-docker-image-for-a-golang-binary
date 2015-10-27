FROM scratch
ADD zoneinfo.tar.gz /
ADD ca-certificates.crt /etc/ssl/certs/
ADD main /
CMD ["/main"]
