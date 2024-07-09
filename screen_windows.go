//go:build windows

package go_sikuli

import "syscall"

var (
	user32              = syscall.NewLazyDLL("user32.dll")
	getForegroundWindow = user32.NewProc("GetForegroundWindow")
	enumDisplayMonitors = user32.NewProc("EnumDisplayMonitors")
	monitorFromWindow   = user32.NewProc("MonitorFromWindow")
)

func GetCurrentScreenIndex() int {
	// 获取当前激活的窗口句柄。
	hwnd, _, _ := getForegroundWindow.Call()
	if hwnd == 0 {
		return -1
	}

	// 获取当前窗口所在的监视器句柄。
	var monitorIndex int = -1
	hMonitor, _, _ := monitorFromWindow.Call(hwnd, uintptr(0))
	if hMonitor == 0 {
		return -1
	}

	monitorIndex = 0
	// 使用EnumDisplayMonitors枚举所有监视器，并使用回调函数找到当前窗口所在的监视器的索引
	monitorEnumProc := syscall.NewCallback(func(hMonitorEnum, hdcMonitor, lprcMonitor, lParam uintptr) uintptr {
		if hMonitorEnum == hMonitor {
			return 0 // Stop enumeration, as we found the monitor
		}
		monitorIndex++
		return 1 // Continue enumeration
	})

	enumDisplayMonitors.Call(0, 0, uintptr(monitorEnumProc), 0)

	return monitorIndex
}
