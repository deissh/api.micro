package helpers

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	type args struct {
		key      string
		fallback string
	}
	type env struct {
		key string
		val string
	}
	tests := []struct {
		args args
		env  env
		want string
	}{
		{
			args: args{
				key:      "TEST_VAR",
				fallback: "test",
			},
			env: env{
				key: "SOME_VAR",
				val: "val",
			},
			want: "test",
		},
		{
			args: args{
				key:      "TEST_VAR",
				fallback: "123",
			},
			env: env{
				key: "SOME_VAR",
				val: "val",
			},
			want: "123",
		},
		{
			args: args{
				key:      "SOME_VAR",
				fallback: "test",
			},
			env: env{
				key: "SOME_VAR",
				val: "val",
			},
			want: "val",
		},
		{
			args: args{
				key:      "SOME_VAR",
				fallback: "test",
			},
			env: env{
				key: "SOME_VAR",
				val: "44",
			},
			want: "44",
		},
	}
	for _, tt := range tests {
		t.Run("GetEnv", func(t *testing.T) {
			_ = os.Setenv(tt.env.key, tt.env.val)

			if got := GetEnv(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvWithPanic(t *testing.T) {
	type args struct {
		key string
	}
	type env struct {
		key string
		val string
	}
	tests := []struct {
		args    args
		env     env
		want    string
		wantErr bool
	}{
		{
			args: args{
				key: "TEST_VAR",
			},
			env: env{
				key: "SOME_VAR",
				val: "val",
			},
			want:    "",
			wantErr: true,
		},
		{
			args: args{
				key: "TEST_VAR",
			},
			env: env{
				key: "SOME_VAR",
				val: "val",
			},
			want:    "",
			wantErr: true,
		},
		{
			args: args{
				key: "SOME_VAR",
			},
			env: env{
				key: "SOME_VAR",
				val: "val",
			},
			want:    "val",
			wantErr: false,
		},
		{
			args: args{
				key: "SOME_VAR",
			},
			env: env{
				key: "SOME_VAR",
				val: "44",
			},
			want:    "44",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run("GetEnvWithPanic", func(t *testing.T) {
			_ = os.Setenv(tt.env.key, tt.env.val)

			defer func() {
				if r := recover(); r == nil && tt.wantErr {
					t.Errorf("GetEnvWithPanic() panic = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			if got := GetEnvWithPanic(tt.args.key); got != tt.want {
				t.Errorf("GetEnvWithPanic() = %v, want %v", got, tt.want)
			}
		})
	}
}
