ENV_FLAGS=--env_variable=AUTHORIZATION_KEY:secret-agent-man \
	--env_variable=AUTH0_CLIENT_SECRET:$(PET_WHISPERER_AUTH0_CLIENT_SECRET) \
	--env_variable=ANIMAL_SERVICE_BASE_PATH:"https://animalservice-dot-pet-whisperer.appspot.com"

LOCAL_ENV_FLAGS=--env_var AUTHORIZATION_KEY="" \
	--env_var AUTH0_CLIENT_SECRET=$(PET_WHISPERER_AUTH0_CLIENT_SECRET) \
	--env_var ANIMAL_SERVICE_BASE_PATH="http://localhost:8081"


serve :
	dev_appserver.py appengine/webapi/app.yaml appengine/animalservice/app.yaml appengine/imageservice/app.yaml ${LOCAL_ENV_FLAGS}
.PHONY: serve

deploy :
	appcfg.py update appengine/webapi/app.yaml appengine/animalservice/app.yaml appengine/imageservice/app.yaml ${ENV_FLAGS}
.PHONY: deploy

rollback :
	appcfg.py rollback appengine/webapi/app.yaml appengine/animalservice/app.yaml appengine/imageservice/app.yaml ${ENV_FLAGS}
.PHONY: rollback
