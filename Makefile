setup:
	go install github.com/joho/godotenv
	go install github.com/spf13/viper
	go install github.com/sirupsen/logrus
	cp .env-sample .env

init:
	docker-compose up -d --remove-orphans

destroy:
	docker-compose down

migrate:
	sql-migrate up -config=dbconfig.yml -env=development

consumer:
	go run cmd/consumer/main.go



