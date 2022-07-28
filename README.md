# SSL proxy server

This tool implements a server that listens for TLS traffic and proxies it to 
another service over http. This would be useful for situations when your service
does not support TLS or it is difficult to enable. A concrete example would be 
the jmx prometheus exporter for kafka that doesn't talk TLS. To fix that we would:

- [generate keys and self-signed cert](https://gist.github.com/drio/920e08aee8aa0d2ff549e2c38b2beb22#file-readme-md).
- add a container running this proxy and point it to the jmx kafka metrics service

At this point you can hit the endpoint via TLS but your client will complain because
it doesn't not trust the CA that singed the cert. To solve that:

- Tell your OS to [trust the CA](https://gist.github.com/drio/920e08aee8aa0d2ff549e2c38b2beb22#in-math-we-trust)

## Usage

```
# generates keys and self-signed cert.
$ make
# Run the testing service
$ go run service/service.go
# Now run the server/proxy
$ go run proxy.go -proxy-url http://localhost:8080
# Now, hit the proxy via TLS and see how the request is forwarded
$ curl -k https://localhost
Hello, this is the service.
```

Notice how we use curl's `-k` flag here to ignore the CA trust error.
In a real deployment, you want to tell your OS that you [trust the certificate](https://gist.github.com/drio/920e08aee8aa0d2ff549e2c38b2beb22#in-math-we-trust).
