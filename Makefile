dockerup:
	docker-compose --project-name="simple-go-microservices" up -d

dockerstart:
	docker-compose --project-name="simple-go-microservices" start

dockerstop:
	docker-compose --project-name="simple-go-microservices" stop

dockerdown:
	docker-compose --project-name="simple-go-microservices" down