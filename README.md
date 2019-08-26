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
- [x] ~~gomybatis~~ mysql配置
- [x] JWT、~~session~~
- [ ] redis
- [ ] godotenv
- [ ] i18n
- [ ] 定时任务
- [ ] docker
- [ ] 用户管理
- [ ] 角色管理
- [ ] 权限管理
- [ ] 内容管理
- [ ] github 第三方登录

### URL
swagger： http://localhost:3000/swagger/index.html
gin-swagger： https://github.com/swaggo/gin-swagger