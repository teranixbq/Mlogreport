include .env

POSTGRESQL_URL := postgres://$(DBUSER):$(DBPASS)@$(DBHOST):$(DBPORT)/$(DBNAME)

dburl:
	export POSTGRESQL_URL='$(POSTGRESQL_URL)'

migratedb:
	migrate -database ${POSTGRESQL_URL} -path app/database/migrations $(filter-out $@,$(MAKECMDGOALS))

up : migratedb 
down: migratedb

run:
	echo "alias run='go run main.go'" | tee -a ~/.bashrc

build:
	GOOS=linux GOARCH=arm64 go build -o bootstrap main.go

zip: build
	zip bootstrap.zip bootstrap