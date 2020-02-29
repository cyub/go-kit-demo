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


### 安装与配置prometheus和grafana

1. 启动prometheus和grafana

```
docker-compose up -d
```
2. prometheus配置：

编辑`prometheus.yml`

```
- job_name: 'go-kit-demo-account'
    scrape_interval: 5s
    static_configs:
      - targets: ['192.168.33.10:8888']
        labels:
          group: 'demo'
```

3. prometheus查询语句

qps:

```
sum by (path) (rate(go_kit_demo_account_http_requests_total[5m]))
```

响应速度：

```
sum(rate(go_kit_demo_account_http_request_duration_seconds_sum[5m])/rate(go_kit_demo_account_http_request_duration_seconds_count[5m]))
```

响应时间中位百分比：

```
histogram_quantile(
  0.5,
  sum by (le) (rate(go_kit_demo_account_http_request_duration_seconds_bucket[5m]))
)
```

### 构建应用

```
make build
```

#### 运行应用

```
make run // 注意consul.host和consul.port参数是否与本地一直
```

