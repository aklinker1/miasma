#!/bin/bash
echo ""

C_RESET="\x1b[0m"
C_BOLD="\x1b[1m"
C_DIM="\x1b[2m"
C_RED="\x1b[91m"
C_ORANGE="\x1b[93m"
C_GREEN="\x1b[92m"
C_CYAN="\x1b[96m"

function _todo {
    echo -e "$C_ORANGE$BOLD ● $1$C_RESET"
}
function _done {
    echo -e "$C_GREEN$BOLD ✔ $1$C_RESET"
}
function _fail {
    echo -e "$C_RED$BOLD ✖ $1$C_RESET"
}

function isDockerInstalled {
    if [[ "$(which docker)" == "" ]]; then
        echo false
    else
        echo true
    fi
}


# Steps:
# 1. Install docker

DOCKER_INSTALL_REQUIRED=true

echo -e "${C_BOLD}Checking system...${C_RESET}"

if [[ "$(lsb_release -a 2> /dev/null | grep Ubuntu)" == "" ]]; then
    _fail "Installer script only supports Ubuntu"
    exit 1
else
    _done "Supported OS"
fi

if [[ "$(isDockerInstalled)" == "false" ]]; then
    _todo "Install Docker"
else
    _done "Install Docker (already installed)"
    DOCKER_INSTALL_REQUIRED=false
fi

if [[ $DOCKER_INSTALL_REQUIRED == true ]]; then
    echo ""
    echo -e "${C_BOLD}Installing required components:${C_RESET}"
fi

if [[ $DOCKER_INSTALL_REQUIRED == true ]]; then 
    _todo "Installing docker"
    
    sudo apt remove docker docker.io containerd runc
    sudo apt install -qy \
        apt-transport-https \
        ca-certificates \
        curl \
        gnupg-agent \
        software-properties-common
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
    sudo add-apt-repository \
        "deb [arch=arm64] https://download.docker.com/linux/ubuntu \
        $(lsb_release -cs) \
        stable"
    sudo apt-get update
    sudo apt-get -y install docker-ce docker-ce-cli containerd.io
    sudo usermod -aG docker $USER

    if [[ "$(isDockerInstalled)" == "false" ]]; then
        _fail "Docker was not installed"
    else
        _done "Installed docker"
    fi
fi

echo ""
