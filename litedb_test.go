package litedb

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
)

func TestMain(t *testing.T) {

	db := LiteDb{dir_path: "./test_db"}

	for i := 0; i < 100; i++ {
		random_num := fmt.Sprint(rand.Intn(100))
		map_data := map[string]string{"data": random_num}

		json_value, _ := json.Marshal(map_data)
		db.Save("test.json", json_value)

		data, err := db.Open("test.json")

		var buffor map[string]interface{}
		json.Unmarshal(data, &buffor)

		if err != nil {
			t.Error("Error reading")
		}

		if buffor["data"] != random_num {
			t.Errorf("Error: expected %s. get %s", random_num, buffor["data"])

			colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, fmt.Sprintf("Test %d failed", i+1))
			fmt.Println(colored)
			continue
		}
		colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, fmt.Sprintf("Test %d passed", i+1))
		fmt.Println(colored)
	}
}
