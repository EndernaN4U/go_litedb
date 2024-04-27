package litedb

import (
	"encoding/json"
	"testing"
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
	t.Run("Saving", Saving)
	t.Run("Reading", Reading)
}

func Saving(t *testing.T) {
	json_data, _ := json.Marshal(test_user)
	test_table.NewDoc(json_data)
}

func Reading(t *testing.T) {
	last_index := test_table.MaxIndex()

	user_data := test_table.Doc(last_index)

	var user_json user
	json.Unmarshal(user_data, &user_json)

	if user_json.Nickname != test_user.Nickname {
		t.Error("Wrong data!")
	}
}
