# golang 博客后台

### Usage

#### set proxy
```bash
export GO111MODULE=on
export GOPROXY=https://goproxy.io

```

#### run

```bash
go get github.com/oxequa/realize
realize start
```

#### init swagger docs

```bash
 go get -u github.com/swaggo/swag/cmd/swag
swag init
```

## TodoList

- [x] realize 热更新
- [x] swagger 文档
- [ ] gomybatis mysql配置
- [ ] JWT、session
- [ ] redis
- [ ] godotenv
- [ ] i18n
- [ ] 定时任务
- [ ] docker

### URL
swagger： http://localhost:3000/swagger/index.html
gin-swagger： https://github.com/swaggo/gin-swagger