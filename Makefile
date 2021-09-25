up-infra:
	docker network create internal_comunication
	docker-compose -f ./BaseA/docker-compose.yml up -d 

down-infra:
	docker-compose -f ./BaseA/docker-compose.yml down
	docker network rm internal_comunication
	
	sudo rm -rf ./BaseA/data