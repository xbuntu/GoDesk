package controller

import (
	"context"
	"godesk/app/util"
	"gorm.io/gorm"
)

// Base 控制器基类
type Base struct {
	ctx context.Context
}

const (
	SuccessCode = 200
	ErrorCode   = 101
	NoLoginCode = 501
)

// Res 返回结果
type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// res 返回
func (b *Base) res(code int, message string, data interface{}) Res {
	return Res{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// success 返回成功
func (b *Base) success(data interface{}) Res {
	return b.res(SuccessCode, "操作成功", data)
}

// error 返回错误
func (b *Base) error(message string) Res {
	return b.res(ErrorCode, message, nil)
}

// setCtx 设置上下文对象
func (b *Base) setCtx(ctx context.Context) {
	b.ctx = ctx
}

// db 获取数据库操作对象和数据库初始化
func (b *Base) db() *gorm.DB {
	return util.NewUtil().Db()
}

// log 增加日志记录
func (b *Base) log(content string) *Base {
	util.NewUtil().Log(content)
	return b
}

// schema 根据model自动建立数据表
func (b *Base) schema(dst ...interface{}) {
	util.NewUtil().Schema(dst...)
}

// getAppPath 获取应用主目录
func (b *Base) getAppPath() string {
	return util.NewUtil().GetAppPath()
}

// pathExist 判断文件目录是否存在，不存在创建
func (b *Base) pathExist(path string) string {
	return util.NewUtil().PathExist(path)
}

// shellCMD 以shell方式运行cmd命令
func (b *Base) shellCMD(cmdStr string, paramStr string) {
	util.NewUtil().ShellCMD(cmdStr, paramStr)
}

// setSystemBackground 设置系统壁纸
func (b *Base) setSystemBackground(path string) {
	util.NewUtil().SetSystemBackground(path)
}

// httpGet get请求
func (b *Base) httpGet(url string) []byte {
	return util.NewUtil().HttpGet(url)
}

// openDir 打开指定目录
func (b *Base) openDir(path string) {
	util.NewUtil().OpenDir(path)
}

// pathConvert 路径转换
func (b *Base) pathConvert(path string) string {
	return util.NewUtil().PathConvert(path)
}
