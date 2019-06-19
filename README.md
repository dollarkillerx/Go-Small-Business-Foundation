# Go-Small-Business-Foundation
使用GO作WEB开发的基础结构

### 目录结构
``` 
.
├── backend  // 后台
│   ├── container // 控制器
│   ├── main.go // 主函数
│   ├── router // 路由层
│   │   └── router.go
│   └── view // 视图层
│       ├── static 
│       │   ├── css
│       │   ├── img
│       │   └── js
│       └── user
├── config // 配置
│   └── config.go
├── datamodels // 模型层
├── datasource // 数据源
│   ├── mysql_conn
│   │   └── mysql.go
│   └── redis_conn
│       └── redis.go
├── defs // 基础定义
│   ├── defs.go
│   └── request.go
├── go.mod
├── LICENSE
├── README.md
├── repositories // dao层
├── request // 返回定义
│   └── request.go
└── services // 业务层
```