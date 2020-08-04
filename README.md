# Usage

### Preparing the executable deployer

Compile the `deployer.go`,

```
make build
```

Copy the `deployer.tgz` and `scripts/init.sh` to Google Drive then make it a public viewable link.

### Setting up the host VM

Download & run the initial script,

```
curl -L -o init.sh \
https://drive.google.com/uc\?export\=download\&id\=$FILE_ID && \
chmod 744 init.sh && ./init.sh
```

Download the deployer,

```
curl -L \
https://drive.google.com/uc\?export\=download\&id\=$FILE_ID | \
tar xz && chmod 744 deployer
```

Log in to Docker Hub before accessing private repositories,

```
sudo docker login -u artisandigitalasia
```

Run, `./deployer`

### Setting up the Jenkins

Jenkins Build Pipeline calls the deployer on the host VM,

```
curl -X POST \$IP_OF_THE_VM:8080/deploy \
-H "secret:deploymenow" \
-H "project:$PROJECT" \
-H "branch:$BRANCH"
```


### Usange
```
DOCKERHUB_ACCOUNT=youraccount PORT=999 SECRET=yoursecret /path/to/deployer
```
