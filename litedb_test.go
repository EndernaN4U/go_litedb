package litedb

import (
	"testing"
)

// func TestMain(t *testing.T) {
// 	db := LiteDb{dir_path: "./test_db"}
// 	tab := db.Table("test")
// 	//tab.NewDoc([]byte("{\"data\": \"data\"}"))
// 	data := tab.Doc(0)
// 	var buf map[string]string
// 	json.Unmarshal(data, &buf)

// 	if buf["data"] != "data" {
// 		t.Error("data wrong")
// 	}

// 	fmt.Println(buf)
// }

type res struct {
	Data string `json:"data"`
}

func TestReadDir(t *testing.T) {
	db := LiteDb{dir_path: "./test_db"}
	tab := db.Table("test")
	tab.DeleteDoc(0)
}
