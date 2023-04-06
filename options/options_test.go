package options

import (
	"reflect"
	"testing"
)

func TestInitOptions(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want *Options
	}{
		{
			name: "not set option",
			args: args{opts: []Option{}},
			want: &Options{
				optionA: "",
				optionB: 0,
				optionC: false,
			},
		},
		{
			name: "set option a",
			args: args{opts: []Option{
				SetOptionA("test"),
			}},
			want: &Options{
				optionA: "test",
				optionB: 0,
				optionC: false,
			},
		},
		{
			name: "set option b",
			args: args{opts: []Option{
				SetOptionB(1),
			}},
			want: &Options{
				optionA: "",
				optionB: 1,
				optionC: false,
			},
		},
		{
			name: "set option c",
			args: args{opts: []Option{
				SetOptionC(true),
			}},
			want: &Options{
				optionA: "",
				optionB: 0,
				optionC: true,
			},
		},
		{
			name: "set option a and b",
			args: args{opts: []Option{
				SetOptionA("test"),
				SetOptionB(1),
			}},
			want: &Options{
				optionA: "test",
				optionB: 1,
				optionC: false,
			},
		},
		{
			name: "set all option",
			args: args{opts: []Option{
				SetOptionA("test"),
				SetOptionB(1),
				SetOptionC(true),
			}},
			want: &Options{
				optionA: "test",
				optionB: 1,
				optionC: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitOptions(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
