package jsons

import (
	"encoding/json"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Testタグ付き構造体(t *testing.T) {
	// 実体
	p := Person{
		Name: "yokoyama",
		Age:  32,
	}

	jsonBytes, err := json.Marshal(p)
	if err != nil {
		t.Errorf("json marshal failed")
	}

	t.Log(p)
	t.Log(string(jsonBytes))

	// ポインタ
	pp := &Person{
		Name: "yokoyama",
		Age:  32,
	}

	pjsonBytes, err := json.Marshal(pp)
	if err != nil {
		t.Errorf("json marshal failed")
	}

	t.Log(pp)
	t.Log(string(pjsonBytes))
}

func TestJSON文字列から構造体へ(t *testing.T) {
	bytes := []byte(`{"name":"yokoyama","age":32}`)
	p := &Person{}

	err := json.Unmarshal(bytes, p)
	if err != nil {
		t.Errorf("json unmarshal failed")
	}

	t.Log(p)
}
