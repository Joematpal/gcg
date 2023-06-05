package main

import "testing"

func Test_splitSourceType(t *testing.T) {
	type args struct {
		sourceType string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{
			name: "should pass",
			args: args{
				sourceType: "github.com/hlubek/metaprogramming-go/domain.Product",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := splitSourceType(tt.args.sourceType)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitSourceType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("splitSourceType() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("splitSourceType() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
