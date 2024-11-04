# jwt-auth-server
sns-archive 認証サーバー

## DBに接続する
```bash
$ mysql -u root -p -h 127.0.0.1 -P 3307 sns_archive_jwt
```

## migration
```bash
$ migrate --path db/migrations --database 'mysql://root:root@tcp(127.0.0.1:3307)/sns_archive_jwt' -verbose up
```
