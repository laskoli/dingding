# fastwego/dingding

A fast [dingding](https://ding-doc.dingtalk.com/) development sdk written in Golang

[![GoDoc](https://pkg.go.dev/badge/github.com/fastwego/dingding?status.svg)](https://pkg.go.dev/github.com/fastwego/dingding?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/fastwego/dingding)](https://goreportcard.com/report/github.com/fastwego/dingding)

## 快速开始 & demo

```shell script
go get github.com/fastwego/dingding
```
```go
// 创建应用实例
app := dingding.New(dingding.Config{
    AppKey:         viper.GetString("AppKey"),
    AppSecret:      viper.GetString("AppSecret"),
})

// 调用 api
resp, err := microapp.List(App, []byte(``))
fmt.Println(string(resp), err)
```

完整的演示项目：

[https://github.com/fastwego/dingding-demo](https://github.com/fastwego/dingding-demo)

接口列表：

[doc/apilist.md](doc/apilist.md)

## 架构设计

![sdk](./doc/img/sdk.jpg)

## 框架特点

### 快速

「快」作为框架设计的核心理念，体现在方方面面：

- 使用 Go 语言，开发快、编译快、部署快、运行快，轻松服务海量用户
- 丰富的[文档](https://pkg.go.dev/github.com/fastwego/dingding) 和 [演示代码](https://github.com/fastwego/dingding-demo) ，快速上手，5 分钟即可搭建一套完整的钉钉服务
- 独立清晰的模块划分，快速熟悉整个框架，没有意外，一切都是你期望的样子
- 甚至连框架自身的大部分代码也是自动生成的，维护更新快到超乎想象

### 符合直觉

作为第三方开发框架，尽可能贴合官方文档和设计，不引入新的概念，不给开发者添加学习负担

### 简洁而不过度封装

作为具体业务和钉钉之间的中间层，专注于通道的角色：帮业务把配置/材料投递到，将响应/推送透传回业务

至于 AccessToken 管理 和 消息加解密处理，框架内部完成得干净利落，开发者甚至觉察不到存在

### 官方文档就是最好的文档

每个接口的注释都附带官方文档的链接，让你随时翻阅，省时省心

### 完备的单元测试

100% 覆盖每一个接口，让你每一次调用都信心满满

### 详细的日志

每个关键环节都为你完整记录，Debug 倍轻松，你可以自由定义日志输出，甚至可以关闭日志

### 活跃的开发者社区

FastWeGo 是一套丰富的 Go 开发框架，钉钉、飞书、微信等服务，拥有庞大的开发者用户群体

你遇到的所有问题几乎都可以在社区找到解决方案


## 参与贡献

欢迎提交 pull request / issue / 文档，一起让钉钉开发更快更好！

Faster we go together!

[加入开发者交流群](https://github.com/fastwego/fastwego.dev#%E5%BC%80%E5%8F%91%E8%80%85%E4%BA%A4%E6%B5%81%E7%BE%A4)
