docker compose -f docker/docker-compose.yml up -d
go run .

go run github.com/steebchen/prisma-client-go migrate dev