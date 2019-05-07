FROM scratch

ADD ["bin/dnsau.b", "/"]

CMD ["/dnsau.b"]