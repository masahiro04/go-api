### コマンド
```
docker-compose build
docker-compose run --rm --service-ports app
docker-compose run db sql-migrate up -config=db/dbconfig.yml -env development
```
