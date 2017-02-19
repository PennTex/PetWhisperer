ENV_FLAGS=--env_variable=AUTHORIZATION_KEY:secret-agent-man \
	--env_variable=AUTH0_DOMAIN:$(PET_WHISPERER_AUTH0_DOMAIN) \
	--env_variable=AUTH0_CLIENT_ID:$(PET_WHISPERER_AUTH0_CLIENT_ID) \
	--env_variable=AUTH0_CLIENT_SECRET:$(PET_WHISPERER_AUTH0_CLIENT_SECRET) \
	--env_variable=AUTH0_CALLBACK_URL:$(PET_WHISPERER_AUTH0_CALLBACK_URL)

LOCAL_ENV_FLAGS=--env_var AUTHORIZATION_KEY="" \
	--env_var AUTH0_DOMAIN=$(PET_WHISPERER_AUTH0_DOMAIN) \
	--env_var AUTH0_CLIENT_ID=$(PET_WHISPERER_AUTH0_CLIENT_ID) \
	--env_var AUTH0_CLIENT_SECRET=$(PET_WHISPERER_AUTH0_CLIENT_SECRET) \
	--env_var AUTH0_CALLBACK_URL="http://localhost:8080/callback"


serve :
	dev_appserver.py appengine/webapi/app.yaml appengine/animalservice/app.yaml ${LOCAL_ENV_FLAGS}
.PHONY: serve

deploy :
	appcfg.py update appengine/webapi/app.yaml appengine/animalservice/app.yaml ${ENV_FLAGS}
.PHONY: deploy

rollback :
	appcfg.py rollback appengine/webapi/app.yaml appengine/animalservice/app.yaml ${ENV_FLAGS}
.PHONY: rollback
