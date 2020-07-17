build:
	rm -rf deployer.tgz deployer && \
	env GOOS=linux GOARCH=amd64 go build deployer.go && \
	tar czf deployer.tgz deployer
