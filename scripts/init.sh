#!/bin/bash

function uninstall_old {
    sudo apt remove docker docker-engine docker.io containerd runc 2>/dev/null
}

function install_prerequisite {
    sudo apt update && \
    sudo apt install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common -y
}

function add_repo_key {
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - && \
    sudo add-apt-repository \
    "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
}

function install_docker_engine {
    sudo apt update && \
    sudo apt install docker-ce docker-ce-cli containerd.io -y
}

uninstall_old
install_prerequisite
add_repo_key
install_docker_engine
