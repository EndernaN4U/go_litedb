package litedb

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Cache struct {
	Last_Index int    `json:"last_index"`
	Table_Name string `json:"table_name"`
}

const (
	cache_file_name = ".cache"
)

func contCacheFile(table_path string) bool {
	_, err := os.Stat(filepath.Join(table_path, cache_file_name))
	return err == nil
}

func createCacheFile(table_path string, tab_name string) {
	cache_path := filepath.Join(table_path, cache_file_name)

	default_cache := Cache{
		Last_Index: 0,
		Table_Name: tab_name,
	}

	json_cache, _ := json.Marshal(default_cache)

	SaveFile(cache_path, json_cache)
}

func cacheData(tab table) Cache {
	cache_path := filepath.Join(tab.path, cache_file_name)

	cache_data, _ := OpenFile(cache_path)

	return BytesToJson[Cache](cache_data)
}

func cacheUpdate(tab table, data Cache) {
	cache_path := filepath.Join(tab.path, cache_file_name)
	json_cache, _ := json.Marshal(data)

	SaveFile(cache_path, json_cache)
}
