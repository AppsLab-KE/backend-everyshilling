package util

import (
	"github.com/google/uuid"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		userId        string
		secretKey     string
		expiryMinutes int
	}
	tests := []struct {
		name      string
		args      args
		wantEmpty bool
		wantErr   bool
	}{
		{
			name: "Generate token",
			args: args{
				userId:        uuid.New().String(),
				secretKey:     "SUPER_SECRET_KEY",
				expiryMinutes: 7000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.userId, tt.args.secretKey, tt.args.expiryMinutes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" != tt.wantEmpty {
				t.Errorf("GenerateToken() got = %v, want %v", got, err)
			}
		})
	}
}

func TestVerifyToken(t *testing.T) {
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODA2NTg1NTQyNzcyNTUxODYsInV1aWQiOiIzZjc0NzQ5NS04NDk1LTQ3MTgtOGVjYy03NDczNDRlNWYzOWUifQ.aBj8Cv5ncXgpKHaRovvil0gLTwd652eyEV99enob1_A"
	userId := "3f747495-8495-4718-8ecc-747344e5f39e"

	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODAyMzkzOTczNTgzMjA2OTUsInV1aWQiOiI1Y2NmN2ExOC0yMjE4LTQxNTItOWRiYS04ODExZDZhNDU1ZTMifQ.BLIgdkZrZg2h7jnmyl3c--ff0PT24ndwVZMJx5-dikU"

	type args struct {
		jwtToken string
	}
	tests := []struct {
		name       string
		args       args
		wantUserId string
		wantErr    bool
	}{
		{
			name: "VerifyToken",
			args: args{
				jwtToken: jwtToken,
			},
			wantUserId: userId,
			wantErr:    false,
		}, {
			name: "VerifyExpired",
			args: args{
				jwtToken: expiredToken,
			},
			wantUserId: "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserId, err := VerifyToken(tt.args.jwtToken, "SUPER_SECRET_KEY")
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserId != tt.wantUserId {
				t.Errorf("VerifyToken() gotUserId = %v, want %v", gotUserId, tt.wantUserId)
			}
		})
	}
}
