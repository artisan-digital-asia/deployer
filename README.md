# Usage

### Preparing the executable deployer

Compiles the `deployer.go`,

```
env GOOS=linux GOARCH=amd64 go build deployer.go && tar czf deployer.tgz deployer
```

Copy the `deployer.tgz` and `scripts/init.sh` to Google Drive then make it as a public viewable.

### Setting up the host VM

Download & run the initial script,

```
curl -L -o init.sh \
https://drive.google.com/uc\?export\=download\&id\=1BjMDEYlD2mE2zeYEmEVX23Rg2NrD9TNk && \
chmod 744 init.sh && ./init.sh
```

Download the deployer,

```
curl -L \
https://drive.google.com/uc\?export\=download\&id\=1VBdbaggKbUGTEBucpgcSR5ZEFt4PvDU_ | \
tar xz && chmod 744 deployer
```

Log in before accessing the private GitLab repositories,

```
docker login -u artisandigitalasia
```

Run, `./deployer`

### Setting up the Jenkins

Jenkins Pipeline calls the deployer on the host VM,

```
curl -X POST \$IP_OF_THE_VM:8080/deploy \
-H "secret:deploymenow" \
-H "project:$PROJECT" \
-H "branch:$BRANCH"
```
