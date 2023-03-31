package util

import "testing"

func TestGenerateHash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name      string
		args      args
		wantEmpty bool
		wantErr   bool
	}{
		{
			name: "Generate Hash",
			args: args{
				password: "pass1234",
			},
			wantErr:   false,
			wantEmpty: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateHash(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateHash() errors = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" != tt.wantEmpty {
				t.Errorf("GenerateHash() got = %v, want %v", got, tt.wantEmpty)
			}
		})
	}
}

func TestCompareHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}

	hash, _ := GenerateHash("pass123")
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Correct password",
			args: args{
				password: "pass123",
				hash:     hash,
			},
			want: true,
		}, {
			name: "InCorrect password",
			args: args{
				password: "pass321",
				hash:     hash,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareHash(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CompareHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
