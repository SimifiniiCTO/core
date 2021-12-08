# setup database (postgres)
cleanup_test_env:
	docker-compose -f compose/docker-compose.postgres.yaml down

setup_test_env:
	docker-compose -f compose/docker-compose.postgres.yaml up --build --force-recreate --no-deps -d
	sh -c "echo waiting 10s to let database start... && sleep 10"

run_test:
	go test -v -race -cover -coverprofile=prof.out ./...

test: setup_test_env run_test cleanup_test_env

coverage: setup_test_env run_test
	go tool cover -html=prof.out
	make cleanup_test_env
