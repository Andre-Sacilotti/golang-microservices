# Exemplos de infra e microserviços


## Como usar?

Primeiramente precisaremos instalar o make, para utilizar o Makefile

```
sudo apt-get install build-essential
```

Em seguida, precisamos instalar o docker e docker-compsoe

```
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

sudo add-apt-repository  "deb [arch=amd64] https://download.docker.com/linux/ubuntu  $(lsb_release -cs) stable"

sudo apt-get update

sudo apt-get -y install docker-ce

sudo curl -L https://github.com/docker/compose/releases/download/1.17.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose

sudo chmod +x /usr/local/bin/docker-compose
```

Feito isso, para subirmos a infraestrutura é simples:

```
make up-dev
```

Agora podemos acessar as documentações das apis e realizar as requisições:

http://0.0.0.0:80/docs/auth   -> Serviço de autenticação
http://0.0.0.0:80/docs/citizen  -> Serviço da base A


## Testes

Se quiser rodar os testes, basta em cada pasta das apis, rodar o seguinte comando:


```
go test -v ./...
```
