## 存储相关逻辑，包含持久化和缓存等


### ent 数据库实体模型使用指南

实体初始化命令：

```shell
entc init [ModelName]    
```

模型目录:

    ent->schema->[ModelName]


关于软删除：

用法：

    s.User.Update().SoftDelete().Save(ctx)
    s.User.Query().SoftDelete().Where()