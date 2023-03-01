# simple-demo

## 极简版抖音

## 一、项目介绍

这里是由青训营后端**面筋队**开发的**极简版抖音**。本项目已实现：视频流功能、视频投稿功能、个人主页展示、喜欢列表功能、用户评论功能、关系列表功能 以及 消息功能。

项目服务地址：http://81.68.118.43:8080

Github 地址：https://github.com/baobaopandas/simple-douyin

运行方法：

```
git clone https://github.com/baobaopandas/simple-douyin
cd simple-douyin
go mod tidy
go build && ./simple-demo
```


## 二、项目分工

| **团队成员** | **主要贡献**                       |
| -------- | ------------------------------ |
| 芦世杰（队长）  | 统筹安排任务，搭建服务器，配置运行环境，负责实现用户关注功能 |
| 包智超      | 主要负责实现视频流功能、用户模块和投稿模块          |
| 王小强      | 主要负责实现用户聊天的消息功能                |
| 查俊荣      | 主要负责实现用户评论功能                   |
| 程宇涛      | 主要负责实现用户点赞功能                   |

  


## 三、项目实现

### 3.1 技术选型与相关开发文档

配置：

1.  [go 1.17](https://go.dev/#)
1.  数据库：[MySQL](https://dev.mysql.com/doc/)
1.  Web框架：[Gin 1.8.2](https://github.com/gin-gonic/gin/)
1.  其他配置：选取了[jwt](https://github.com/golang-jwt/jwt)来实现用户的鉴权, 采取了[Nginx](http://nginx.org/)作为http服务器内容分发, 使用[ffmpeg](https://github.com/FFmpeg/FFmpeg)来进行图像处理, 使用[redis](https://github.com/redis/go-redis)来进行部分模块的缓存实现，添加了[Viper](https://github.com/spf13/viper)进行系统参数的配置统一管理
1.  服务器配置：2核4G，100G云盘，6Mbps带宽

  


### 3.2 架构设计

#### 总体架构：

![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/8f27de18986d4abb975aaaf3370359ff~tplv-k3u1fbpfcp-zoom-1.image)

#### 数据库结构：

![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/568177905715435cadb389e8c1525141~tplv-k3u1fbpfcp-zoom-1.image)

  


### 3.3 项目代码介绍

见[项目代码文档](https://jobn0zzvhc.feishu.cn/docx/Qlw9dE0tsojDXOx1Zekcd4eNn3d#doxcnQKJmzD2iDFtaibZnWqUjQb)

## 四、测试结果

见[项目测试](https://jobn0zzvhc.feishu.cn/docx/Qlw9dE0tsojDXOx1Zekcd4eNn3d#doxcnbWCWhKh88gFfKMomb2O11d)

## 五、Demo 演示视频

见[项目Demo 演示视频](https://jobn0zzvhc.feishu.cn/docx/Qlw9dE0tsojDXOx1Zekcd4eNn3d#doxcnabNXsq4nfmfyKT2szEwNcb)

## 成员

| **团队成员** | **Github地址**                       |
| -------- | ------------------------------ |
| 包智超      | https://github.com/baobaopandas          |
| 芦世杰      | https://github.com/EDGlushijie           |
| 王小强      | https://github.com/XiaoqWang                |
| 查俊荣      | https://github.com/wind9whisper                   |
| 程宇涛      | https://github.com/Tao-Cute                   |
