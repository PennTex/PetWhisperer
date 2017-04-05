ENV_FLAGS=--env_variable=AUTHORIZATION_KEY:"secret-agent-man" \
	--env_variable=AUTH0_CLIENT_SECRET:$(PET_WHISPERER_AUTH0_CLIENT_SECRET) \
	--env_variable=ANIMAL_SERVICE_BASE_PATH:"https://animalservice-dot-pet-whisperer.appspot.com" \
	--env_variable=IMAGE_SERVICE_BASE_PATH:"https://imageservice-dot-pet-whisperer.appspot.com"

LOCAL_ENV_FLAGS=--env_var AUTHORIZATION_KEY="secret-dev" \
	--env_var AUTH0_CLIENT_SECRET=$(PET_WHISPERER_AUTH0_CLIENT_SECRET) \
	--env_var ANIMAL_SERVICE_BASE_PATH="http://localhost:8081" \
	--env_var IMAGE_SERVICE_BASE_PATH="http://localhost:8082" \
	--default_gcs_bucket_name "staging.pet-whisperer.appspot.com"


serve :
	dev_appserver.py appengine/webapi/app.yaml appengine/animalservice/app.yaml appengine/imageservice/app.yaml ${LOCAL_ENV_FLAGS}
.PHONY: serve

deploy :
	appcfg.py update appengine/webapi/app.yaml appengine/animalservice/app.yaml appengine/imageservice/app.yaml ${ENV_FLAGS}
.PHONY: deploy

rollback :
	appcfg.py rollback appengine/webapi/app.yaml appengine/animalservice/app.yaml appengine/imageservice/app.yaml ${ENV_FLAGS}
.PHONY: rollback
