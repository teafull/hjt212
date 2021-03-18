package utils

import (
	"reflect"
	"testing"
)

func TestGetQN(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "get qn",
			want: []byte("1616080158556"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQN(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQN() = %v, want %v", got, tt.want)
			}
		})
	}
}
