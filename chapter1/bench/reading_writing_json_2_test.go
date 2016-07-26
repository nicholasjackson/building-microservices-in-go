package bench

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type RequestNoTag struct {
	Name string
}

type RequestWithTag struct {
	Name string `json:"name"`
}

func BenchmarkUnmarshal(b *testing.B) {
	b.ResetTimer()

	data := []byte(`{"name": "World"}`)

	for i := 0; i < b.N; i++ {
		reader := bytes.NewReader(data)

		body, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Println(err)
		}

		var request RequestNoTag
		err = json.Unmarshal(body, &request)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func BenchmarkUnmarshalWithDecoder(b *testing.B) {
	b.ResetTimer()

	data := []byte(`{"name": "World"}`)

	for i := 0; i < b.N; i++ {
		reader := bytes.NewReader(data)
		decoder := json.NewDecoder(reader)

		var request RequestNoTag
		err := decoder.Decode(&request)
		if err != nil {
			fmt.Println(err)
		}
	}
}
