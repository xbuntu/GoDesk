package main

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"syscall"
	"unsafe"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// BrowserOpenURL 默认浏览器打开网址
func (a *App) BrowserOpenURL(url string) {
	runtime.BrowserOpenURL(a.ctx, url)
}

// ClipboardGetText 获取剪切板内容
func (a *App) ClipboardGetText() (string, error) {
	return runtime.ClipboardGetText(a.ctx)
}

// ClipboardSetText 设置剪切板内容
func (a *App) ClipboardSetText(text string) error {
	return runtime.ClipboardSetText(a.ctx, text)
}

// ShellCMD 以shell方式运行cmd命令
func (a *App) ShellCMD(cmdStr string, paramStr string) {
	open, _ := syscall.UTF16PtrFromString("open")
	cmd, _ := syscall.UTF16PtrFromString(cmdStr)
	param, _ := syscall.UTF16PtrFromString(paramStr)
	shell32 := syscall.NewLazyDLL("shell32.dll")
	ShellExecuteW := shell32.NewProc("ShellExecuteW")
	_, _, _ = ShellExecuteW.Call(uintptr(0), uintptr(unsafe.Pointer(open)), uintptr(unsafe.Pointer(cmd)), uintptr(unsafe.Pointer(param)), uintptr(0), uintptr(1))
}
