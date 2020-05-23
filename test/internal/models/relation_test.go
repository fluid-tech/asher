package models

import (
	"asher/internal/models"
	"reflect"
	"testing"
)

func TestRelation(t *testing.T) {
	arr := []string{
		"order", "product",
	}
	brr := []string{
		"asorder", "product",
	}
	var relationTests = []struct {
		in  [][]string
		out models.Relation
	}{
		{in: [][]string{arr, nil}, out: models.Relation{HasMany: arr, HasOne: nil}},
		{in: [][]string{nil, brr}, out: models.Relation{HasMany: nil, HasOne: brr}},
		{in: [][]string{arr, brr}, out: models.Relation{HasMany: arr, HasOne: brr}},
	}

	for _, element := range relationTests {
		if !reflect.DeepEqual(element.in[0], element.out.HasMany) {
			t.Error("unexpected error while comparing HasMany arrays")
		}
		if !reflect.DeepEqual(element.in[1], element.out.HasOne) {
			t.Error("unexpected error while comparing HasOne arrays")
		}
	}
}
