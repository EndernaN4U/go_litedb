package litedb

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type user struct {
	First_Name  string `json:"first_name"`
	Second_Name string `json:"second_name"`
	Nickname    string `json:"nickname"`
}

var (
	test_db    = LiteDb{dir_path: "./test_db"}
	test_table = test_db.Table("test")
	test_user  = user{
		First_Name:  "Ben",
		Second_Name: "Ten",
		Nickname:    "Ben10",
	}
)

func TestMain(t *testing.T) {
	// t.Run("Saving", Saving)
	// t.Run("Reading", Reading)
	// t.Run("Filtering", Filter)
	// t.Run("Crypto", Crypto)
}

func Saving(t *testing.T) {
	json_data, _ := json.Marshal(test_user)
	test_table.NewDoc(json_data)
}

func Reading(t *testing.T) {
	last_index := test_table.LastIndex()

	user_data := test_table.Doc(UnID(last_index))

	var user_json user
	json.Unmarshal(user_data, &user_json)

	if user_json.Nickname != test_user.Nickname {
		t.Error("Wrong data!")
	}
}

func Filter(t *testing.T) {
	criteria := []Criterion{
		{Name: "first_name", Value: "Ben"},
		{Name: "age", Value: float64(21)},
	}

	fmt.Println(test_table.Filter(criteria))
}

func Crypto(t *testing.T) {
	for i := 1; i < 1000; i++ {
		random_str := genRandomString(i)
		fmt.Println(i, random_str, len(random_str))
	}
}

func TestCompare(t *testing.T) {
	user_data, _ := json.Marshal(test_user)

	ai_start := time.Now().UnixMilli()
	table_ai := test_db.Table("table_ai")

	for i := 0; i < 100; i++ {
		table_ai.NewDoc(user_data)
	}

	ai_end := time.Now().UnixMilli()

	fmt.Println("ai:", ai_end-ai_start)

	cr_start := time.Now().UnixMilli()
	table_cr := test_db.Table("table_cr")

	for i := 0; i < 100; i++ {
		table_cr.NewDocCrypto(user_data)
	}

	cr_end := time.Now().UnixMilli()

	fmt.Println("cr:", cr_end-cr_start)
}
