
createmigrate: #make createmigrate name='name of migration'
	go run ./cmd/migrate  create -ext sql -dir db/migration $(name)

migrateup: 
	go run ./cmd/migrate  up

migratedown: 
	go run ./cmd/migrate  down 1

upload-submodule:
	git submodule init
	git submodule update

run-docker: upload-submodule
	docker-compose up  --build --force-recreate -d  pnl

dev-env:
	docker-compose up -d postgres


unit-tests:
	go test -v ./cmd/... ./config/... ./external/...  ./internal/... ./pkg/...

test: upload-submodule
	docker-compose up --build --force-recreate --remove-orphans --abort-on-container-exit --exit-code-from pnl_test pnl_test

doc:
	docker-compose up --build --force-recreate -d swagger