package litedb

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type cache struct {
	Last_Index int    `json:"last_index"`
	Table_Name string `json:"table_name"`
}

const (
	cache_file_name = ".cache"
)

func contCacheFile(dir_path string) bool {
	_, err := os.Stat(filepath.Join(dir_path, cache_file_name))
	return err != nil
}

func createCacheFile(dir_path string, tab_name string) {
	cache_path := filepath.Join(dir_path, cache_file_name)
	default_cache := cache{
		Last_Index: 0,
		Table_Name: tab_name,
	}

	json_cache, _ := json.Marshal(default_cache)

	SaveFile(cache_path, json_cache)
}
