#!/usr/bin/env bash
docker run -v /root/ndp/dnsau/bin/:/dnsau/bin/ -v /root/ndp/dnsau/dnsau/src/:/dnsau/src/ -e "CGO_ENABLED=0" dnsau-builder
