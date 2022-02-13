### docker
```
# image
docker-compose build

# db作成
docker-compose run db createdb -h db -U postgresql api
docker-compose run test-db createdb -h test-db -U postgresql test-api

# migration適用
docker-compose run make migrate
docker-compose run make test_migrate

# 起動
docker-compose up
```
