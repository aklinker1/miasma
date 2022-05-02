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
function isSwarmNode {
    if [ "$(docker info 2> /dev/null | grep Swarm | sed 's/Swarm: //g')" == " inactive" ]; then
        echo false;
    else
        echo true;
    fi
}
function isServerRunning {
    if [[ "$(docker container ls | grep miasma)" == "" ]]; then
        echo false
    else
        echo true
    fi
}


# Steps:
# 1. Install docker
# 1. Initialize the swarm
# 1. Start the server

DOCKER_INSTALL_REQUIRED=true
SWARM_INIT_REQUIRED=true
START_REQUIRED=true

echo -e "${C_BOLD}Checking system...${C_RESET}"

if [[ "$(lsb_release -a 2> /dev/null | grep Ubuntu)" == "" ]]; then
    _fail "Installer script only supports Ubuntu"
    exit 1
else
    _done "Supported OS"
fi

if [[ "$(isDockerInstalled)" == "false" ]]; then
    _todo "Install Docker"
    _todo "Initialize the swarm"
    _todo "Start the server"
else
    _done "Install Docker (already installed)"
    DOCKER_INSTALL_REQUIRED=false
    if [[ "$(isSwarmNode)" == "false" ]]; then
        _todo "Initialize the swarm"
    else
        _done "Initialize the swarm"
        SWARM_INIT_REQUIRED=false
    fi

    if [[ "$(isServerRunning)" == "false" ]]; then
        _todo "Start the server"
    else
        _done "Start the server"
        START_REQUIRED=false
    fi
fi

if [[ $DOCKER_INSTALL_REQUIRED == true || $SWARM_INIT_REQUIRED == true || $START_REQUIRED == true ]]; then
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

if [[ $SWARM_INIT_REQUIRED == true ]]; then 
    _todo "Initializing swarm"

    docker swarm init &> /dev/null

    if [[ "$(isSwarmNode)" == "false" ]]; then
        _fail "Swarm could not be initialized"
    else
        _done "Initialized swarm"
    fi
fi

if [[ $START_REQUIRED == true ]]; then
    _todo "Start the server"

    docker pull aklinker1/miasma
    
    docker run -d \
        --restart unless-stopped \
        -p 3000:3000 \
        -v /var/run/docker.sock:/var/run/docker.sock \
        -v $HOME/.miasma:/data/miasma \
        aklinker1/miasma
    if [[ "$(isServerRunning)" == "false" ]]; then
        _fail "Could not start the server"
    else
        _done "Started the server"
    fi
fi

if [[ "$(isSwarmNode)" == "true" ]]; then
    echo ""
    echo -e "${C_BOLD}To add a node to this cluster, SSH into the machine and run this command:${C_RESET}"
    echo ""
    echo -en "$C_CYAN"
    docker swarm join-token worker | grep docker
    echo -en "$C_RESET"
fi

echo ""
