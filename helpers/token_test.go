package helpers

import (
	"testing"
)

func TestContainsString(t *testing.T) {
	type args struct {
		s []string
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestContainsString", args: args{s: []string{"1", "2"}, v: "2"}, want: true},
		{name: "TestContainsString", args: args{s: []string{"1", "2"}, v: "3"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsString(tt.args.s, tt.args.v); got != tt.want {
				t.Errorf("ContainsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsStrings(t *testing.T) {
	type args struct {
		s []string
		v []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestContainsString", args: args{s: []string{"1", "2"}, v: []string{"1", "2"}}, want: true},
		{name: "TestContainsString", args: args{s: []string{"1", "2", "3"}, v: []string{"1", "3"}}, want: true},
		{name: "TestContainsString", args: args{s: []string{"1", "2"}, v: []string{"1", "3"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsStrings(tt.args.s, tt.args.v); got != tt.want {
				t.Errorf("ContainsStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
