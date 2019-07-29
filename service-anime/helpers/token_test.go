package helpers

import (
	"github.com/jarcoal/httpmock"
	"gopkg.in/resty.v1"
	"net/http"
	"reflect"
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

func TestTokenVerify(t *testing.T) {
	type args struct {
		accessToken string
		required    bool
		roles       []string
		scopes      []string
	}
	type mock struct {
		json   tokenResponse
		status int
	}
	tests := []struct {
		name    string
		args    args
		mock    mock
		want    Token
		wantErr bool
	}{
		{
			name: "Verify invalid token",
			args: args{
				accessToken: "invalidtoken",
				required:    true,
				roles:       []string{"admin", "superadmin"},
				scopes:      []string{"news", "notif"},
			},
			mock: mock{
				json: tokenResponse{
					Version: "1",
					Token:   Token{},
				},
				// token not founded
				status: http.StatusNotFound,
			},
			wantErr: true,
		},
		{
			name: "Verify token but without roles",
			args: args{
				accessToken: "somesotekn",
				required:    true,
				roles:       []string{"admin", "superadmin"},
				scopes:      []string{"news", "notif"},
			},
			mock: mock{
				json: tokenResponse{
					Version: "1",
					Token: Token{
						Role: "user",
						Permissions: []string{
							"news",
							"notif",
						},
					},
				},
				status: http.StatusOK,
			},
			wantErr: true,
		},
		{
			name: "Verify token but without permissions",
			args: args{
				accessToken: "somesotekn",
				required:    true,
				roles:       []string{"admin", "superadmin"},
				scopes:      []string{"news", "notif"},
			},
			mock: mock{
				json: tokenResponse{
					Version: "1",
					Token: Token{
						Role: "admin",
						Permissions: []string{
							"notif",
							"email",
							"messages",
						},
					},
				},
				status: http.StatusOK,
			},
			wantErr: true,
		},
		{
			name: "Verify token",
			args: args{
				accessToken: "somesotekn",
				required:    true,
				roles:       []string{"admin", "superadmin"},
				scopes:      []string{"news", "notif"},
			},
			mock: mock{
				json: tokenResponse{
					Version: "1",
					Token: Token{
						Role: "admin",
						Permissions: []string{
							"notif",
							"news",
						},
					},
				},
				status: http.StatusOK,
			},
			want: Token{
				UserID: 0,
				Role:   "admin",
				Permissions: []string{
					"notif",
					"news",
				},
			},
			wantErr: false,
		},
		{
			name: "Verify token",
			args: args{
				accessToken: "somesotekn",
				required:    true,
				roles:       []string{"admin", "superadmin"},
				scopes:      []string{"news", "notif"},
			},
			mock: mock{
				json: tokenResponse{
					Version: "1",
					Token: Token{
						UserID: 1,
						Role:   "superadmin",
						Permissions: []string{
							"notif",
							"news",
							"messages",
							"profile",
						},
					},
				},
				status: http.StatusOK,
			},
			want: Token{
				UserID: 1,
				Role:   "superadmin",
				Permissions: []string{
					"notif",
					"news",
					"messages",
					"profile",
				},
			},
			wantErr: false,
		},

		{
			name: "Verify invalid token but token not required",
			args: args{
				accessToken: "invalidtoken",
				required:    false,
				roles:       []string{"admin", "superadmin"},
				scopes:      []string{"news", "notif"},
			},
			mock: mock{
				json: tokenResponse{
					Version: "1",
					Token:   Token{},
				},
				// token not founded
				status: http.StatusNotFound,
			},
			wantErr: false,
		},
		{
			name: "Verify token but without roles but token not required",
			args: args{
				accessToken: "somesotekn",
				required:    false,
				roles:       []string{"admin", "superadmin"},
				scopes:      []string{"news", "notif"},
			},
			mock: mock{
				json: tokenResponse{
					Version: "1",
					Token: Token{
						Role: "user",
						Permissions: []string{
							"news",
							"notif",
						},
					},
				},
				status: http.StatusOK,
			},
			wantErr: true,
		},
		{
			name: "Verify token but without permissions but token not required",
			args: args{
				accessToken: "somesotekn",
				required:    false,
				roles:       []string{"admin", "superadmin"},
				scopes:      []string{"news", "notif"},
			},
			mock: mock{
				json: tokenResponse{
					Version: "1",
					Token: Token{
						Role: "admin",
						Permissions: []string{
							"notif",
							"email",
							"messages",
						},
					},
				},
				status: http.StatusOK,
			},
			wantErr: true,
		},
		{
			name: "Verify token but token not required",
			args: args{
				accessToken: "somesotekn",
				required:    false,
				roles:       []string{"admin", "superadmin"},
				scopes:      []string{"news", "notif"},
			},
			mock: mock{
				json: tokenResponse{
					Version: "1",
					Token: Token{
						Role: "admin",
						Permissions: []string{
							"notif",
							"news",
						},
					},
				},
				status: http.StatusOK,
			},
			want: Token{
				UserID: 0,
				Role:   "admin",
				Permissions: []string{
					"notif",
					"news",
				},
			},
			wantErr: false,
		},
		{
			name: "Verify token but token not required",
			args: args{
				accessToken: "somesotekn",
				required:    false,
				roles:       []string{"admin", "superadmin"},
				scopes:      []string{"news", "notif"},
			},
			mock: mock{
				json: tokenResponse{
					Version: "1",
					Token: Token{
						UserID: 1,
						Role:   "superadmin",
						Permissions: []string{
							"notif",
							"news",
							"messages",
							"profile",
						},
					},
				},
				status: http.StatusOK,
			},
			want: Token{
				UserID: 1,
				Role:   "superadmin",
				Permissions: []string{
					"notif",
					"news",
					"messages",
					"profile",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.ActivateNonDefault(resty.DefaultClient.GetClient())
			fakeURL := GetEnv("SERVICE_AUTH", "http://service-auth:8080") + "/token.check?access_token=" + tt.args.accessToken
			httpmock.RegisterResponder(
				"GET",
				fakeURL,
				httpmock.NewJsonResponderOrPanic(tt.mock.status, tt.mock.json),
			)

			got, err := TokenVerify(tt.args.accessToken, tt.args.required, tt.args.roles, tt.args.scopes)
			if (err != nil) != tt.wantErr {
				t.Errorf("TokenVerify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TokenVerify() = %v, want %v", got, tt.want)
			}
		})
	}
}
