package file

import "os"

// DirExists 检查目录是否存在
func DirExists(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		return false
	}

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

// FileExists 检查文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !os.IsNotExist(err)
}
