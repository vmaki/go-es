```shell
# 生成数据模型
go run main.go make model 模型名称
go run main.go make migration add_users_table // 生成迁移文件 add_users_table
```

### docker

```shell
docker exec -it mysql

mysql -uroot -p
$ use mysql;
$ update user set host='%' where user='root';
$ FLUSH PRIVILEGES;
```
