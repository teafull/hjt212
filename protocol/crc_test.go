package protocol

import (
	"testing"
)

func TestCalCrc(t *testing.T) {
	type args struct {
		puchMsg []byte
	}

	hjt212data := "QN=20160801085857223;ST=32;CN=1062;PW=100000;MN=010000A8900016F000169DC0;Flag=5;CP=&&RtdInterval=30&&"
	data2 := "QN=20200814151600001;ST=22;CN=2011;PW=123456;MN=010000A8900016F000169DC0;Flag=7;CP=&&DataTime=20200814151600;LA-td=50.1,LA-Flag=N;a34004-Rtd=207,a34004-Flag=N;a34002-Rtd=295,a34002-Flag=N;a01001-Rtd=12.6,a01001-Flag=N;a01002-Rtd=32,a01002-Flag=N;a01006-Rtd=101.02,a01006-Flag=N;a01007-Rtd=2.1,a01007-Flag=N;a01008-Rtd=120,a01008-Flag=N;a34001-Rtd=217,a34001-Flag=N;&&"
	data3 := "QN=20210210073530000;ST=22;CN=2011;MN=aqms-1100-2020110011;PW=123456;Flag=4;CP=&&DataTime=20210210073530;a34004-Rtd=33,a34004-Flag=N;a34002-Rtd=42,a34002-Flag=N;pms_rules-Rtd=7,pms_rules-Flag=N;a01001-Rtd=10,a01001-Flag=N;a01002-Rtd=73,a01002-Flag=N;a01006-Rtd=102.3,a01006-Flag=N;lng-Rtd=120.32879,lng-Flag=N;lat-Rtd=31.17347,lat-Flag=N;a21026-Rtd=2,a21026-Flag=N;a21004-Rtd=28,a21004-Flag=N;a05024-Rtd=58,a05024-Flag=N;a21005-Rtd=0.5,a21005-Flag=N;a99054-Rtd=0.137,a99054-Flag=N;hot_t-Rtd=39.3,hot_t-Flag=N;a01007-Rtd=2.2,a01007-Flag=N;a01008-Rtd=41,a01008-Flag=N;vn-Rtd=1.04,vn-Flag=N&&"

	tests := []struct {
		name string
		args args
		want uint
	}{
		// TODO: Add test cases.
		{name: "cal crc of data package of hjt212data", args: args{puchMsg: []byte(hjt212data)}, want: 7296},
		{name: "cal crc of data package of data2", args: args{puchMsg: []byte(data2)}, want: 0x5CC0},
		{name: "cal crc of data package of data3", args: args{puchMsg: []byte(data3)}, want: 0x5D80},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := CalCrc(tt.args.puchMsg); got != tt.want {
				t.Errorf("CalCrc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyCrc(t *testing.T) {
	type args struct {
		dataSegment []byte
		crc         uint
	}

	hjt212data := "QN=20160801085857223;ST=32;CN=1062;PW=100000;MN=010000A8900016F000169DC0;Flag=5;CP=&&RtdInterval=30&&"
	data2 := "QN=20200814151600001;ST=22;CN=2011;PW=123456;MN=010000A8900016F000169DC0;Flag=7;CP=&&DataTime=20200814151600;LA-td=50.1,LA-Flag=N;a34004-Rtd=207,a34004-Flag=N;a34002-Rtd=295,a34002-Flag=N;a01001-Rtd=12.6,a01001-Flag=N;a01002-Rtd=32,a01002-Flag=N;a01006-Rtd=101.02,a01006-Flag=N;a01007-Rtd=2.1,a01007-Flag=N;a01008-Rtd=120,a01008-Flag=N;a34001-Rtd=217,a34001-Flag=N;&&"
	data3 := "QN=20210210073530000;ST=22;CN=2011;MN=aqms-1100-2020110011;PW=123456;Flag=4;CP=&&DataTime=20210210073530;a34004-Rtd=33,a34004-Flag=N;a34002-Rtd=42,a34002-Flag=N;pms_rules-Rtd=7,pms_rules-Flag=N;a01001-Rtd=10,a01001-Flag=N;a01002-Rtd=73,a01002-Flag=N;a01006-Rtd=102.3,a01006-Flag=N;lng-Rtd=120.32879,lng-Flag=N;lat-Rtd=31.17347,lat-Flag=N;a21026-Rtd=2,a21026-Flag=N;a21004-Rtd=28,a21004-Flag=N;a05024-Rtd=58,a05024-Flag=N;a21005-Rtd=0.5,a21005-Flag=N;a99054-Rtd=0.137,a99054-Flag=N;hot_t-Rtd=39.3,hot_t-Flag=N;a01007-Rtd=2.2,a01007-Flag=N;a01008-Rtd=41,a01008-Flag=N;vn-Rtd=1.04,vn-Flag=N&&"

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "cal crc of data package of hjt212data", args: args{dataSegment: []byte(hjt212data), crc: 7296}, want: true, wantErr: false},
		{name: "cal crc of data package of data2", args: args{dataSegment: []byte(data2), crc: 0x5CC0}, want: true, wantErr: false},
		{name: "cal crc of data package of data2 verify failed", args: args{dataSegment: []byte(data2), crc: 0x5CC1}, want: false, wantErr: false},
		{name: "cal crc of data package of data3", args: args{dataSegment: []byte(data3), crc: 0x5D80}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyCrc(tt.args.dataSegment, tt.args.crc)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyCrc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyCrc() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCalCrc(b *testing.B) {
	hjt212data := "QN=20160801085857223;ST=32;CN=1062;PW=100000;MN=010000A8900016F000169DC0;Flag=5;CP=&&RtdInterval=30&&"

	for i := 0; i < b.N; i++ {
		CalCrc([]byte(hjt212data))
	}
}
