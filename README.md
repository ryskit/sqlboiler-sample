# sqlboiler-sample

## 手順

1. PostgreSQLコンテナを起動
   ```shell
    docker run --name postgres \
           -e POSTGRES_PASSWORD=password \
           -e POSTGRES_INITDB_ARGS="--encoding=UTF8 --no-locale" \
           -e TZ=Asia/Tokyo \
           -v postgresdb:/var/lib/postgresql/data \
           -p 5432:5432 \
           -d postgres
    ```
2. スキーマからテーブルを生成
   - schema.sqlをインポートする
3. モデルを生成
   以下を実行することで`models/`にファイルが生成される
   ```shell
   go generate
   ```

