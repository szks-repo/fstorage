package fstorage

import (
	"fmt"
	"github.com/szks-repo/fstorage/test/testutil"
	"log"
	"os"
	"strings"
	"testing"
)

func getTestStoragePath() string {
	wd, _ := os.Getwd()
	return fmt.Sprintf("%s/test/teststorage", wd)
}

func getTestClient() *StorageClient {

	base := fmt.Sprintf(getTestStoragePath())
	c, err := New(base)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func cleanTestStorage() {
	ts := getTestStoragePath()
	entries, _ := os.ReadDir(ts)
	for _, entry := range entries {
		os.RemoveAll(fmt.Sprintf("%s/%s", ts, entry.Name()))
	}
}

func Test_New(t *testing.T) {
	type arg struct {
		base string
	}

	wd, _ := os.Getwd()

	tests := []struct {
		name string
		arg  arg
		ok   bool
	}{
		{name: "fail", arg: arg{base: "test"}, ok: false},
		{name: "fail", arg: arg{base: ""}, ok: false},
		{name: "fail", arg: arg{base: "/"}, ok: false},
		{name: "fail", arg: arg{base: "test/test1"}, ok: false},
		{name: "fail", arg: arg{base: "./client_test.go"}, ok: false},
		{name: "fail", arg: arg{base: testutil.AbsolutePath("/var/www/phantomDir")}, ok: false},
		{name: "fail", arg: arg{base: wd + "/client_test.go"}, ok: false},
		{name: "ok1", arg: arg{base: wd + "/test/teststorage"}, ok: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.arg.base)
			ok := err == nil
			if tt.ok != ok {
				if err != nil {
					fmt.Println(err)
				}
				t.Errorf("wantOk %v, gotOk %v", tt.ok, ok)
			}
		})
	}
}

func Test_Write(t *testing.T) {
	c := getTestClient()
	defer cleanTestStorage()
	type arg struct {
	}

	t.Run("", func(t *testing.T) {
		filename := testutil.RandFileName(".txt")
		err := c.Save(filename, strings.NewReader("test"), nil)
		fmt.Println("Err=>", err)

		err = c.SaveAll("res/"+filename, strings.NewReader("x"), nil)
		fmt.Println("Err=>", err)
	})

}
