### コマンド
```
# image作成
docker-compose build
# appの起動
docker-compose run --rm --service-ports app
# migration up
docker-compose exec app bash -c "sql-migrate up -config=db/dbconfig.yml -env development"
```
