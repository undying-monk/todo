build-model:
	fastText-0.9.2/fasttext supervised -input backend/internal/classify/data.txt -output model

frontend:
	npm run frontend/dev

backend:
	go run backend/cmd/server/main.go