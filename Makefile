docker: 
docker compose -f docker/docker-compose.yml up -d
go:
go run .
prisma-go:
go run github.com/steebchen/prisma-client-go migrate dev

go-prisma-db-push:
go run github.com/steebchen/prisma-client-go db push