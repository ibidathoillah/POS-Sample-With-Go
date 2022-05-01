#!/bin/bash -x

COMPOSE_VERSION="1.23.2"
COMPOSE_URL="https://github.com/docker/compose/releases/download/$COMPOSE_VERSION/docker-compose-$(uname -s)-$(uname -m)"


# Docker installation
install_docker() {
  curl -fsSL https://get.docker.com/ | bash
  sudo bash <<EOS
usermod -a -G docker $USER
curl -L "$COMPOSE_URL" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
EOS
}

install_mysql() {
  bash <<EOS
docker-compose -f $(pwd)/bin/mysql-majoo.yml up -d
EOS
}

install_postgres() {
  bash <<EOS
docker-compose -f $(pwd)/bin/postgres-majoo.yml up -d
EOS
}


if [ -x "$(command -v docker)" ]; then
    echo "Docker already installed"
else
    echo "Installing docker"
    install_docker
fi

install_mysql