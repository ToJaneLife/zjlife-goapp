# zjlife-goapp

至简生活前端+后台

![](./screenshots/QQ20170508-171952@2x.png)

## 项目框架

项目采用beego 生成，需要有一定的golang 基础


```
├── conf
│   └── app.conf
├── controllers
│   ├── weather.go
│   └── default.go
├── main.go
├── models
│   └── models.go
├── static
│   ├── css
│   ├── fonts
│   ├── img
│   └── js
└── views
    ├── index.tpl

```

## 下载安装

通过git 指令，下载项目

```
goi clone https://github.com/ToJaneLife/zjlife-goapp
```

进入根目录

```
bee run
```

通过 [http://127.0.0.1:8082](http://127.0.0.1:8082) 访问

## 衍生项目

- [微信小程序-致简生活](https://github.com/ToJaneLife/zjlife-weapp)

## 一起交流

欢迎提 [issue](https://github.com/ToJaneLife/zjlife-weapp/issues) 或者 [pr](https://github.com/ToJaneLife/zjlife-weapp/pulls) 。