build:
	rm -rf deployer.tgz deployer && \
	./go-executable-build.bash deployer.go && \
	tar czvf deployer.tgz deployer.go-*