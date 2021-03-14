package protocol

import (
	"reflect"
	"testing"
)

func TestHjtDecoder_Decoder(t *testing.T) {
	type args struct {
		dadaPack []byte
	}
	data := "##0367QN=20200814151600001;ST=22;CN=2011;PW=123456;MN=010000A8900016F000169DC0;Flag=7;CP=&&DataTime=20200814151600;LA-td=50.1,LA-Flag=N;a34004-Rtd=207,a34004-Flag=N;a34002-Rtd=295,a34002-Flag=N;a01001-Rtd=12.6,a01001-Flag=N;a01002-Rtd=32,a01002-Flag=N;a01006-Rtd=101.02,a01006-Flag=N;a01007-Rtd=2.1,a01007-Flag=N;a01008-Rtd=120,a01008-Flag=N;a34001-Rtd=217,a34001-Flag=N;&&5CC0\r\n"
	tests := []struct {
		name    string
		args    args
		want    Hjt212Package
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "normal-1",
			args: args{
				dadaPack: []byte(data),
			},
			want: Hjt212Package{
				QN:      []byte("20200814151600001"),
				ST:      22,
				CN:      2011,
				PW:      []byte("123456"),
				MN:      []byte("010000A8900016F000169DC0"),
				Flag:    7,
				PNum:    1,
				Package: map[int]map[string]string{1: {"DataTime": "20200814151600"}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HjtDecoder{}
			got, err := h.Decoder(tt.args.dadaPack)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
