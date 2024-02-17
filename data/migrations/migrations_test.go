package migrations

import (
	"fmt"
	"jw/utils"
	"testing"
)

func TestAdd(t *testing.T) {
	u, d := Add("test-migartions-add")
	defer func(u string, d string) {
		fmt.Println("Deleting files", u, d)
		_ = utils.DeleteFile(u)
		_ = utils.DeleteFile(d)
	}(u, d)

	if !utils.CheckFile(u) {
		t.Error("File does not exist")
	}
	if !utils.CheckFile(d) {
		t.Error("File does not exist")
	}
}
