### docker
```
# image
docker-compose build

# 起動
docker-compose up

# migration適用
docker-compose run app sql-migrate up -config=db/dbconfig.yml -env development
```
