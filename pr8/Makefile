start.compose:
	cd deployments && docker compose up --build

compose.up:
	cd deployments && docker compose up

start.db:
	cd deployments && docker compose up db -d

start.app:
	cd deployments && docker compose up app -d

start.web-server:
	cd deployments && docker compose up web-server -d

start.local:
	cd cmd/myapp && go build -o app && \
	DB_URL="postgres://postgres:mydbpassword@localhost:5432/postgres?sslmode=disable&" ./app

down:
	cd deployments && docker compose down
