package api

import (
	"asher/internal/api"
	"testing"
)

func TestEmitterFile(t *testing.T) {
	e := api.NewEmitterFile("Hello", "App/Transactor", "gibberish", 1)

	if e.Content != "gibberish" || e.FileName != "Hello" || e.Path != "App/Transactor" || e.FileType != 1 {
		t.Errorf("fail harder")
	}


}
