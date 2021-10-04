### local
```
# localのposgreでDB作成
createdb golang_api
# migration実行
mike migrate
# サーバー起動
mike run
```

### docker
```
# image作成
docker-compose build
# appの起動
docker-compose run --rm --service-ports app

# migration up
docker-compose up -d # まずはコンテナ起動
docker-compose exec app bash -c "sql-migrate up -config=db/dbconfig.yml -env development"
```
