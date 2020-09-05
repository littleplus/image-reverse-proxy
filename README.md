# 图片反向代理
## 简介
本项目主要适用于，在使用查找图片插件中yandex image、google image、saucenao等以图搜图网站，对某些低分辨率图片进行原图的查找的时候，以图搜图网站因为referer、IP等无法获取到低分辨率图片的问题，开发了一个适用且仅适用于图片的反向代理。

## 主要功能
将目标图片下载，经过转换格式为90质量的jpeg图片格式后，保存到运行本项目的服务器中的/tmp/img目录下，并使得以图搜图网站可以访问此图片。

## 解决问题
本代理通过url获取到图片所在网站的hostname，并在请求时，将其添加到http header的referer里，可以解决大多数限制访问的网站的图片无法被以图搜图网站访问的问题。

## 部署方式
### 下载及编译
```shell script
go get https://github.com/littleplus/image-reverse-proxy
cd $GOPATH/src/github.com/littleplus/image-reverse-proxy
go build -o imageproxy cmd/main.go
```
### 部署
```shell script
通过增加环境变量PORT可以指定程序的监听端口，默认为8080
```
### 使用
**暂时不支持https**\
使用主要是在浏览器的以图搜图插件里面添加一项
```shell script
#反向代理并使用使用saucenao
http://<你的服务器+端口>或<域名>/s?pic=%s&engine=saucenao
#反向代理并使用yandex image
http://<你的服务器+端口>或<域名>/s?pic=%s&engine=yandex
#反向代理并使用google image
http://<你的服务器+端口>或<域名>/s?pic=%s&engine=google
```
### 常见的以图搜图插件
Edge:
```shell script
https://microsoftedge.microsoft.com/addons/detail/ibmfjngadieonblglgamabghhaimfldg
```