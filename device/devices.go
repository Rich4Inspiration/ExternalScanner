package device

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
)

// StorageInfo represents the storage information of a device.
type StorageInfo struct {
	Device     string
	Size       uint64
	FileSystem string
}

// GetStorageInfo retrieves the storage information for the given IP address.
func GetStorageInfo(ip string) ([]StorageInfo, error) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve disk partitions: %s", err)
	}

	var storageInfoList []StorageInfo

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Printf("Failed to retrieve disk usage for %s: %s\n", partition.Mountpoint, err)
			continue
		}

		storageInfo := StorageInfo{
			Device:     partition.Device,
			Size:       usage.Total,
			FileSystem: partition.Fstype,
		}

		storageInfoList = append(storageInfoList, storageInfo)
	}

	return storageInfoList, nil
}

/*
将 GetStorageInfo 函数添加到 device 模块中，该函数接受一个 IP 地址作为参数，并返回一个 StorageInfo 列表，其中包含了存储设备的名称、磁盘大小和文件系统类型。
*/
