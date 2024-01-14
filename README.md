# Curl

## Chapter 1

- `-0`/`--http1.0`
- `--get --data-urlencode <keyword>` -> send keyword as query paraeter. escaping according to [RFC3986](https://datatracker.ietf.org/doc/html/rfc3986)
- `-v` -> show verbose info
- `-H "header-name: value"` -> specify request header
- `-A` -> user agent
- `-X <method>`
- `--head`/`-I`
- `-L` -> redirect to location
- `-d`/`--data`/`--data-ascii` -> specify request body. defaulting to POST and application/x-www-form-urlencoded. no escaping.

## Chapter 2

- `-F <name>=<content>` -> multipart/form-data.
- `-c <file>` -> save cookie to file
- `-b <file>` -> load cookie from file, or specify cookie as `name=value`
- `--basic -u <user:pass>` -> use basic auth (add Authorization header)
- `--digest -u <user:pass>` -> use digest auth. send request again with Authorization header after 401 is received
- `-x <host:port>` -> use proxy
