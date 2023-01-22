package pcloud

import "testing"

func Test_makeURL(t *testing.T) {
	authToken = "test"
	t.Cleanup(func() {
		authToken = ""
	})
	type args struct {
		method string
		params []Param
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				method: "test",
				params: []Param{
					{
						Name: "test",
						Val:  "test",
					},
				},
			},
			want: "https://eapi.pcloud.com/test?auth=test&test=test",
		},
		{
			name: "Test 2",
			args: args{
				method: "test",
				params: []Param{
					{
						Name: "test",
						Val:  "test",
					},
					{
						Name: "test2",
						Val:  "test2",
					},
				},
			},
			want: "https://eapi.pcloud.com/test?auth=test&test=test&test2=test2",
		},
		{
			name: "Test no params",
			args: args{
				method: "test",
				params: []Param{},
			},
			want: "https://eapi.pcloud.com/test?auth=test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeURL(tt.args.method, tt.args.params...)
			if got != tt.want {
				t.Errorf("makeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
