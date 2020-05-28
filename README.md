# 開発環境
```shell script
# ビルド
docker-compose -f docker-compose.development.yml build

# 起動（-d オプションなしで、ログが確認できる）
docker-compose -f docker-compose.development.yml up -d

# 停止
docker-compose -f docker-compose.development.yml down

# 起動の確認
docker ps

# sshでログイン(NAMESは、psで確認)
docker exec -it <NAMES> <シェル名(ex:bash)>

# docker-compose -f docker-compose.development.yml start
# docker-compose -f docker-compose.development.yml stop
```

# MySQL接続方法
```shell script
# (dockerコンテナにsshで接続後実行する良いです)
mysql -h localhost --port 3306 --protocol tcp -u root -proot
```
