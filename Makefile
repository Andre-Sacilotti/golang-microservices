up-infra:
	docker network create internal_comunication
	docker-compose -f ./BaseA/docker-compose.yml up -d 
	docker-compose -f ./BaseB/docker-compose.yml up -d 

down-infra:
	docker-compose -f ./BaseA/docker-compose.yml down --volumes
	docker-compose -f ./BaseB/docker-compose.yml down --volumes
	docker network rm internal_comunication
