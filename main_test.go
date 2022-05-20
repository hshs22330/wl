package main

import (
	"reflect"
	"testing"
)

func TestRemoveString(t *testing.T) {
	tests := []struct {
		name string
		data []interface{}
		want []int
	}{
		{name: "remove string ang keep numbers", data: []interface{}{1, 3, 'a', 'b', 4, 7}, want: []int{1, 3, 4, 7}},
		{name: "remove string ang keep numbers", data: []interface{}{}, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveString(tt.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected:%#v, got:%#v", tt.want, got)
			}
		})
	}
}

func TestTransferArray(t *testing.T) {
	tests := []struct {
		name string
		data []interface{}
		want []interface{}
	}{
		{name: "transform the array", data: []interface{}{5, "a", "b", "c", "d", "e", "f", "g"}, want: []interface{}{"a", "d", "c", "b", "e"}},
		{name: "transform the array", data: []interface{}{5, 'a', 'b', 'c', 'd', 'e', 'f', 'g'}, want: []interface{}{'a', 'd', 'c', 'b', 'e'}},
		{name: "transform the array", data: []interface{}{}, want: []interface{}{}},
		{name: "transform the array", data: []interface{}{"a", "b", "c", "d", "e", "f", "g"}, want: []interface{}{}},
		{name: "transform the array", data: []interface{}{5, 'a', 'b', 'c', 'd'}, want: []interface{}{'a', 'c', 'b', 'd'}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TransferArray(tt.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected:%#v, got:%#v", tt.want, got)
			}
		})
	}
}
