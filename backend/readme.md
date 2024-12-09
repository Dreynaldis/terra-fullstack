## Setup the backend and database

cd backend

copy .env.example to .env

make migrate-up
sqlc generate

## run the backend service
make run