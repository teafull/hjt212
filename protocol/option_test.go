package protocol

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExtractDataByHeadAndTail(t *testing.T) {
	type args struct {
		dataSegment []byte
		head        []byte
		tail        []byte
	}

	dataSegmentStr := "QN=20200814151600001;ST=22;CN=2011;PW=123456;MN=010000A8900016F000169DC0;Flag=7;"

	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{name: "QN", args: args{dataSegment: []byte(dataSegmentStr), head: []byte("QN="), tail: []byte(";")}, want: []byte("20200814151600001")},
		{name: "ST", args: args{dataSegment: []byte(dataSegmentStr), head: []byte("ST="), tail: []byte(";")}, want: []byte("22")},
		{name: "CN", args: args{dataSegment: []byte(dataSegmentStr), head: []byte("CN="), tail: []byte(";")}, want: []byte("2011")},
		{name: "MN", args: args{dataSegment: []byte(dataSegmentStr), head: []byte("MN="), tail: []byte(";")}, want: []byte("010000A8900016F000169DC0")},
		{name: "Flag", args: args{dataSegment: []byte(dataSegmentStr), head: []byte("Flag="), tail: []byte(";")}, want: []byte("7")},

		// head or tail empty
		{name: "Flag-head and tail empty", args: args{dataSegment: []byte(dataSegmentStr), head: []byte(""), tail: []byte("")}, want: []byte(dataSegmentStr)},
		{name: "Flag-head empty", args: args{dataSegment: []byte(dataSegmentStr), head: []byte(""), tail: []byte(";")}, want: []byte("QN=20200814151600001")},
		{name: "Flag-tail-empty111", args: args{dataSegment: []byte(dataSegmentStr), head: []byte("Flag="), tail: []byte("")}, want: []byte("7;")},
		{name: "Flag-zero-empty", args: args{dataSegment: []byte(dataSegmentStr), head: []byte("Flag="), tail: []byte(";")}, want: []byte("7")},

		// head or tail not find
		{name: "Flag-tail-not-found", args: args{dataSegment: []byte(dataSegmentStr), head: []byte("Flag="), tail: []byte("-")}, want: []byte("7;")},
		{name: "Flag-head-not-found", args: args{dataSegment: []byte(dataSegmentStr), head: []byte("Flag=1"), tail: []byte(";")}, want: []byte("QN=20200814151600001")},
		{name: "Flag-head_and_tail-not-found", args: args{dataSegment: []byte(dataSegmentStr), head: []byte("Flag=1"), tail: []byte(";;")}, want: []byte(dataSegmentStr)},

		// head or tail empty head or tail not find
		{name: "Flag-head_empty_tail-not-found", args: args{dataSegment: []byte(dataSegmentStr), head: []byte(""), tail: []byte(";;")}, want: []byte(dataSegmentStr)},
		{name: "Flag-tail_empty_head-not-found", args: args{dataSegment: []byte(dataSegmentStr), head: []byte("sss"), tail: []byte("")}, want: []byte(dataSegmentStr)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractDataByHeadAndTail(tt.args.dataSegment, tt.args.head, tt.args.tail)
			fmt.Println(string(got))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractDataByHeadAndTail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
