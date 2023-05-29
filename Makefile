deploy:
	docker build --build-arg IMAGE=simple-auth-v2 -t simple-auth-v2 .
	docker tag simple-auth-v2 hunterjj/simple-auth-v2 
	docker push hunterjj/simple-auth-v2

start:
	kubectl apply -f mongodb-deployment.yaml
	kubectl apply -f backend-deployment.yaml