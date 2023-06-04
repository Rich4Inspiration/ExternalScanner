package ExternalScan

import (
	"ExternalScan/device"
	"ExternalScan/scanner"
	"ExternalScan/ui"
)

func main() {
	// 启动 UI 模块
	go ui.RunUI()

	// 等待 UI 模块完成初始化
	// 这里可以添加适当的等待时间或机制，确保 UI 模块初始化完成

	// 获取 UI 模块中的起始地址和结束地址
	startIP, endIP := ui.GetIPRange()

	// 使用 Scanner 模块扫描在线设备
	onlineDevices := scanner.ScanDevices(startIP, endIP)

	// 使用 Device 模块获取存储设备信息
	storageData := device.GetStorageData(onlineDevices)

	// 更新 UI 模块中的在线终端信息和存储设备信息
	ui.UpdateTerminalList(onlineDevices)
	ui.UpdateStorageData(storageData)

	// 在这里可以继续执行其他操作或等待程序结束
	// ...
}
