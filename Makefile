DAYS?=365
KEY_SIZE?=2048

.PHONY: cert
key-cert:
	openssl req -x509 \
	-newkey rsa:$(KEY_SIZE) \
	-nodes \
	-keyout server.key \
	-out server.crt \
	-sha256 \
	-days $(DAYS) \
	-subj '/C=US/ST=MA/L=Boston/O=Tufts/OU=DataTeam/CN=localhost'

# Extract public key
.PHONY: pub
pub:
	openssl rsa -in server.key -pubout

clean:
	rm -f *.crt *.key localhost.conf *.pem
