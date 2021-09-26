up-infra:
	docker network create internal_comunication
	docker-compose -f ./BaseA/docker-compose.yml up -d 
	docker-compose -f ./BaseB/docker-compose.yml up -d 
	docker-compose -f ./BaseC/docker-compose.yml up -d 

down-infra:
	docker-compose -f ./BaseA/docker-compose.yml down --volumes --remove-orphans
	docker-compose -f ./BaseB/docker-compose.yml down --volumes --remove-orphans
	docker-compose -f ./BaseC/docker-compose.yml down --volumes --remove-orphans
	docker network rm internal_comunication


up-dev:
	docker network create internal_comunication_auth_api
	docker-compose -f ./nginx_gateway/docker-compose.yml up -d 
	docker-compose -f ./auth_api/docker-compose.yml up -d 

down:
	docker-compose -f ./nginx_gateway/docker-compose.yml down --volumes --remove-orphans 
	docker-compose -f ./auth_api/docker-compose.yml down --volumes --remove-orphans 
	docker network rm internal_comunication_auth_api

update-docs:
	export PATH=$(go env GOPATH)/bin:$PATH
	swag  init --dir ./auth_api/
	mv docs ./auth_api/docs