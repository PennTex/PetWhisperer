ENV_FLAGS=--env_variable=AUTH0_DOMAIN:$(PET_WHISPERER_AUTH0_DOMAIN) \
		--env_variable=AUTH0_CLIENT_ID:$(PET_WHISPERER_AUTH0_CLIENT_ID) \
		--env_variable=AUTH0_CLIENT_SECRET:$(PET_WHISPERER_AUTH0_CLIENT_SECRET) \
		--env_variable=AUTH0_CALLBACK_URL:$(PET_WHISPERER_AUTH0_CALLBACK_URL)

LOCAL_ENV_FLAGS=--env_var AUTH0_DOMAIN=$(PET_WHISPERER_AUTH0_DOMAIN) \
		--env_var AUTH0_CLIENT_ID=$(PET_WHISPERER_AUTH0_CLIENT_ID) \
		--env_var AUTH0_CLIENT_SECRET=$(PET_WHISPERER_AUTH0_CLIENT_SECRET) \
		--env_var AUTH0_CALLBACK_URL="http://localhost:8080/callback"


serve :
	dev_appserver.py appengine/webapp/app.yaml appengine/activityservice/app.yaml appengine/animalservice/app.yaml ${LOCAL_ENV_FLAGS}
.PHONY: serve

deploy :
	appcfg.py update appengine/webapp/app.yaml appengine/activityservice/app.yaml appengine/animalservice/app.yaml ${ENV_FLAGS}
.PHONY: deploy

rollback :
	appcfg.py rollback appengine/webapp/app.yaml appengine/activityservice/app.yaml appengine/animalservice/app.yaml ${ENV_FLAGS}
.PHONY: rollback
