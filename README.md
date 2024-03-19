# GoDesk 介绍
>GoDesk 基于golang+wails+vue实现跨平台桌面应用程序开发基础框架，集成了vue路由、基本的登录、element UI
  
下载体验版：https://github.com/xbuntu/GoDesk/releases  
码云体验版：https://gitee.com/xbuntu/go-desk/releases/tag/v1.0-beta  

# QQ交流群  
![](preview/GoDesk.png)

# 预览
![](preview/1.jpg)
![](preview/2.jpg)
![](preview/3.jpg)
![](preview/4.jpg)
![](preview/5.jpg)
![](preview/6.jpg)
![](preview/7.jpg)
![](preview/8.jpg)

# 开发环境搭建
- 1、安装 [go](https://studygolang.com/dl)  
- 2、配置 go 环境变量（参考）

```shell
#/opt/go 替换成 go 安装目录
export GOROOT=/opt/go
export GOPATH=/opt/go/path
export GOBIN=$GOPATH/bin
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export GOPRIVATE=*.gitlab.com,*.gitee.com,*.github.com
export GOSUMDB=sum.golang.google.cn
export PATH=$PATH:$GOROOT/bin:$GOBIN
```

- 3、安装 gcc（推荐安装 [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/) ）  
- 4、安装 [wails](https://wails.io/zh-Hans/docs/gettingstarted/installation/)  
```shell
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```
- 5、安装 [npm](https://nodejs.org/) 

# 运行命令
```shell
#如果首次打包提示 dist目录不存在，需先执行下面两条命令
cd frontend
npm i

#检测
wails doctor

#调试
wails dev

#打包 .exe
wails build 

#打包 .exe 带安装步骤
wails build -nsis -upx
```

# 参考网站
- wails  
>使用 Go 构建漂亮的跨平台应用程序    
https://wails.io/zh-Hans/docs/gettingstarted/installation/

- npm  
>Vue包检索网站    
https://www.npmjs.com/

- vue
>用于构建用户界面的 JavaScript 框架   
https://cn.vuejs.org/

- Element UI
>Element，一套为开发者、设计师和产品经理准备的基于 Vue 2.0 的桌面端组件库    
https://element.eleme.io/#/zh-CN/component/installation

- upx  
>用于压缩您的应用程序    
https://upx.github.io/

- NSIS  
>用于生成 Windows 安装程序   
https://nsis.sourceforge.io/Download

- Gorm  
>orm框架  
https://gorm.io/zh_CN/docs/query.html
 
- GoFrame  
>一款模块化、高性能、企业级的Go基础开发框架。    
https://goframe.org/display/gf
  
- Gin
>Go最快的全功能web框架  
https://gin-gonic.com/

- golang中文网
>go学习交流平台  
https://studygolang.com/dl

- TDM-GCC
>windows系统 gcc 编译环境    
https://jmeubank.github.io/tdm-gcc/download/

# QQ交流群  
![](preview/GoDesk.png)

