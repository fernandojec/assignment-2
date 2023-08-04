build_api:
	docker build -t 172.16.80.157/hacktiv8-final-project/assignment-2:v0.0.2 -f deploy/docker/api/Dockerfile .

build_web:
	docker build -t 172.16.80.157/hacktiv8-final-project/assignment-2-web:v0.0.2 -f deploy/docker/web/Dockerfile .