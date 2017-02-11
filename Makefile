serve-activity: 
	go run services/activity/main.go services/activity/handlers.go
.PHONY: serve-activity

serve-animal: 
	go run services/animal/main.go services/animal/handlers.go
.PHONY: serve-animal