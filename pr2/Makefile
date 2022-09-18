start.compose:
	cd deployments && docker compose up --build

start.local.windows:
	cd cmd/myapp && go build -o app.exe && start app.exe

start.local:
	cd cmd/myapp && go build -o app && ./app

start.db:
	cd deployments && docker compose db -d

start.app:
	cd deployments && docker compose app -d

start.web-server:
	cd deployments && docker compose web-server -d

down:
	cd deployments && docker compose down
