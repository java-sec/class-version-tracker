package utils

import "testing"

func TestMD5(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				bytes: []byte("CC11001100"),
			},
			want:    "c1d29fe4ec649cab6916c93f44711bec",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MD5(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("MD5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MD5() got = %v, want %v", got, tt.want)
			}
		})
	}
}
