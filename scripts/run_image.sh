# run image in the background
docker ps -a --format '{{.Names}}' | grep -Eq "^yappify-api$" && docker start yappify-api || docker run -d -e ENVIRONMENT=development -e DB_URL="postgresql://postgres:postgres@localhost:5432/db?sslmode=disable" -p 8000:8000 --name yappify-api yappify-api