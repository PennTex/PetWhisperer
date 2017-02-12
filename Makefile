serve :
	goapp serve appengine/web-app/app.yaml appengine/activity-service/app.yaml appengine/animal-service/app.yaml
.PHONY: serve

deploy :
	goapp deploy appengine/web-app/app.yaml appengine/activity-service/app.yaml appengine/animal-service/app.yaml
.PHONY: deploy

