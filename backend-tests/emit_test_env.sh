echo "Creating mock PostgreSQL database"
docker-compose up -d

echo "Emitting the backend server"
(cd ../backend && make test && make run&)

echo "Running watch tests"
npm run test:watch

docker-compose down
sudo killall backend