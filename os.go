package litedb

import "os"

func OpenFile(file_path string) ([]byte, error) {
	data, err := os.ReadFile(file_path)
	return data, err
}

func SaveFile(file_path string, data []byte) {
	os.WriteFile(
		file_path,
		data,
		0644,
	)
}

func ReadDir(path string) []string {
	dirs, _ := os.ReadDir(path)

	dirs_names := []string{}
	for _, dir := range dirs {
		dirs_names = append(dirs_names, dir.Name())
	}

	return dirs_names
}
