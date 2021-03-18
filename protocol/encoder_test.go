package protocol

import (
	"reflect"
	"testing"
)

func TestHjtEncoder_makeCmdCp(t *testing.T) {
	type args struct {
		params map[string]string
	}
	tests := []struct {
		name  string
		args  args
		want1 []byte
		want2 []byte
	}{
		// TODO: Add test cases.
		{
			name: "empty param",
			args: args{
				params: map[string]string{},
			},
			want1: []byte("CP=&&&&"),
			want2: []byte("CP=&&&&"),
		},
		{
			name: "only one param",
			args: args{
				params: map[string]string{"SystemTime": "20040516010101"},
			},
			want1: []byte("CP=&&SystemTime=20040516010101&&"),
			want2: []byte("CP=&&SystemTime=20040516010101&&"),
		},
		{
			name: "multi params",
			args: args{
				params: map[string]string{"BeginTime": "20040506111000", "EndTime": "20040506151000"},
			},
			want1: []byte("CP=&&BeginTime=20040506111000,EndTime=20040506151000&&"),
			want2: []byte("CP=&&EndTime=20040506151000,BeginTime=20040506111000&&"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HjtEncoder{}
			got := h.makeCmdCp(tt.args.params)

			if !reflect.DeepEqual(got, tt.want1) && !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("makeCmdCp() = %v, want1 %v, want2 %v", string(got), string(tt.want1), string(tt.want2))
			}
		})
	}
}

func TestHjtEncoder_makeCmdDataPkg(t *testing.T) {
	type args struct {
		hjt212Cmd Hjt212Cmd
	}
	tests := []struct {
		name  string
		args  args
		want1 []byte
		want2 []byte
	}{
		// TODO: Add test cases.
		{
			name: "make cmd data package",
			args: args{
				hjt212Cmd: Hjt212Cmd{
					QN:     []byte("20040516010101001"),
					ST:     32,
					CN:     2051,
					MN:     []byte("88888880000001"),
					PW:     []byte("123456"),
					Flag:   3,
					Params: map[string]string{"BeginTime": "20040506111000", "EndTime": "20040506151000"},
				},
			},
			want1: []byte("QN=20040516010101001;ST=32;CN=2051;PW=123456;MN=88888880000001;Flag=3;CP=&&BeginTime=20040506111000,EndTime=20040506151000&&"),
			want2: []byte("QN=20040516010101001;ST=32;CN=2051;PW=123456;MN=88888880000001;Flag=3;CP=&&EndTime=20040506151000,BeginTime=20040506111000&&"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HjtEncoder{}
			if got := h.makeCmdDataPkg(tt.args.hjt212Cmd); !reflect.DeepEqual(got, tt.want1) && !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("makeCmdDataPkg() = %v, want1 %v, want2 %v", string(got), string(tt.want1), string(tt.want2))
			}
		})
	}
}

func TestHjtEncoder_Encoder(t *testing.T) {
	type args struct {
		hjt212Cmd Hjt212Cmd
	}
	tests := []struct {
		name    string
		args    args
		want1   []byte
		want2   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "hjt212 encoder",
			args: args{
				hjt212Cmd: Hjt212Cmd{
					QN:     []byte("20040516010101001"),
					ST:     32,
					CN:     2051,
					MN:     []byte("88888880000001"),
					PW:     []byte("123456"),
					Flag:   3,
					Params: map[string]string{"BeginTime": "20040506111000", "EndTime": "20040506151000"},
				},
			},
			want1: []byte("##0124QN=20040516010101001;ST=32;CN=2051;PW=123456;MN=88888880000001;Flag=3;CP=&&BeginTime=20040506111000,EndTime=20040506151000&&6B80\r\n"),
			want2: []byte("##0124QN=20040516010101001;ST=32;CN=2051;PW=123456;MN=88888880000001;Flag=3;CP=&&EndTime=20040506151000,BeginTime=20040506111000&&EF41\r\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HjtEncoder{}
			got, err := h.Encoder(tt.args.hjt212Cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want1) && !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("Encoder() got = %v, want1 %v, want2 %v", string(got), string(tt.want1), string(tt.want2))
			}
		})
	}
}

func BenchmarkHjtEncoder_Encoder(b *testing.B) {

	hjt212Cmd := Hjt212Cmd{
		QN:     []byte("20040516010101001"),
		ST:     32,
		CN:     2051,
		MN:     []byte("88888880000001"),
		PW:     []byte("123456"),
		Flag:   3,
		Params: map[string]string{"BeginTime": "20040506111000", "EndTime": "20040506151000"},
	}

	h := &HjtEncoder{}
	for i := 0; i < b.N; i++ {
		h.Encoder(hjt212Cmd)
	}
}
