package config

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

func NewConfig() error {
	var (
		date, _ = time.Parse(time.RFC822, "14 Mar 10 18:00 UTC")
		buf     = new(bytes.Buffer)
	)
	err := toml.NewEncoder(buf).Encode(map[string]any{
		"date":   date,
		"counts": []int{1, 1, 2, 3, 5, 8},
		"hash": map[string]string{
			"key1": "val1",
			"key2": "val2",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	return nil
}
