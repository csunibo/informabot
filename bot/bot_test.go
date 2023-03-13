package bot

import "testing"

func TestStartInformaBot(t *testing.T) {
	type args struct {
		token string
		debug bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartInformaBot(tt.args.token, tt.args.debug)
		})
	}
}
