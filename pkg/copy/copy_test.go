package copy

import (
	"reflect"
	"testing"
)

type Source struct {
	Field1 int
	Field2 string
	Field4 *Destination
	Field5 string
}

type Destination struct {
	Field1 int
	Field2 string
	Field3 bool
	Filed5 int64
}

func TestAssignStruct(t *testing.T) {
	type args struct {
		src interface{}
		dst interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "test1",
			args: args{
				src: &Source{Field1: 10, Field2: "test"},
				dst: &Destination{Field3: true},
			},
			want: &Destination{Field1: 10, Field2: "test", Field3: true},
		},
		{
			name: "test2",
			args: args{
				src: &Source{
					Field1: 10,
					Field2: "test",
					Field4: &Destination{Field1: 1, Field2: "test2", Field3: false}},
				dst: &Destination{Field3: true},
			},
			want: &Destination{Field1: 10, Field2: "test", Field3: true},
		},

		{
			name: "test3",
			args: args{
				src: &Source{
					Field1: 10,
					Field2: "test",
					Field4: &Destination{Field1: 1, Field2: "test2", Field3: false},
					Field5: "hello",
				},
				dst: &Destination{Field3: true, Filed5: 12},
			},
			want: &Destination{Field1: 10, Field2: "test", Field3: true, Filed5: 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssignStruct(tt.args.src, tt.args.dst)
			if !reflect.DeepEqual(tt.args.dst, tt.want) {
				t.Errorf("AssignStruct() = %v, want %v", tt.args.dst, tt.want)
			}
		})
	}
}

func TestDeepCopy(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "test1",
			args: args{
				value: Source{
					Field1: 123,
					Field2: "hello",
					Field4: &Destination{
						Field1: 456,
						Field2: "world",
						Field3: true},
				},
			},
			want: Source{
				Field1: 123,
				Field2: "hello",
				Field4: &Destination{
					Field1: 456,
					Field2: "world",
					Field3: true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeepCopy(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}
