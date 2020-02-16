## 用户服务

### 功能

- 注册用户
- 查询用户
- 数据库迁移

### 技术栈

- [Gorm](https://github.com/jinzhu/gorm)
- [Viper](https://github.com/spf13/viper)
- [mux](https://github.com/gorilla/mux)
- [consul](https://www.consul.io/)


### 安装与部署

### 安装与配置consul

1. 启动consul

```
docker-compose up -d 
```

2. 配置应用的配置

在consul的Key/Value页面下，添加添加目录`config/account-service,dev`, 接下在此目录下配置：
- MYSQL_DB
- MYSQL_HOST
- MYSQL_PORT
- MYSQL_USER
- MYSQL_PASSWD

说明：应用配置以key-value形式存储，存储在`config/application_name,applicaton_env/`目录。比如应用名称demo，项目运行环境是测试环境(test)，则目录名称是`config/demo,test`


### 构建应用

```
make build
```

#### 运行应用

```
make run // 注意consul.host和consul.port参数是否与本地一直
```

