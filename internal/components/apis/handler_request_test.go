package apis

import (
	"github.com/samber/lo"
	"reflect"
	"testing"
	"time"
)

type testTextUnmarshaler int

func (i *testTextUnmarshaler) UnmarshalText(val []byte) error {
	if string(val) == "a" {
		*i = 1
		return nil
	}
	return nil
}

func Test_mapping_textUnmarshalerBinder(t *testing.T) {
	now := time.Now()

	type args struct {
		obj    any
		values map[string][]string
		tag    string
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "testTextUnmarshaler",
			args: args{
				obj: &struct {
					Val testTextUnmarshaler `form:"testTextUnmarshaler"`
				}{},
				values: map[string][]string{
					"testTextUnmarshaler": {"a"},
				},
				tag: "form",
			},
			want: &struct {
				Val testTextUnmarshaler `form:"testTextUnmarshaler"`
			}{
				Val: 1,
			},
			wantErr: false,
		},
		{
			name: "ptrTestTextUnmarshaler",
			args: args{
				obj: &struct {
					Val *testTextUnmarshaler `form:"testTextUnmarshaler"`
				}{},
				values: map[string][]string{
					"testTextUnmarshaler": {"a"},
				},
				tag: "form",
			},
			want: &struct {
				Val *testTextUnmarshaler `form:"testTextUnmarshaler"`
			}{
				Val: lo.ToPtr(testTextUnmarshaler(1)),
			},
			wantErr: false,
		},
		{
			name: "slice-testTextUnmarshaler",
			args: args{
				obj: &struct {
					Val []testTextUnmarshaler `form:"testTextUnmarshaler"`
				}{},
				values: map[string][]string{
					"testTextUnmarshaler": {"a"},
				},
				tag: "form",
			},
			want: &struct {
				Val []testTextUnmarshaler `form:"testTextUnmarshaler"`
			}{
				Val: []testTextUnmarshaler{1},
			},
			wantErr: false,
		},
		{
			name: "slice-ptrTestTextUnmarshaler",
			args: args{
				obj: &struct {
					Val []*testTextUnmarshaler `form:"testTextUnmarshaler"`
				}{},
				values: map[string][]string{
					"testTextUnmarshaler": {"a"},
				},
				tag: "form",
			},
			want: &struct {
				Val []*testTextUnmarshaler `form:"testTextUnmarshaler"`
			}{
				Val: []*testTextUnmarshaler{lo.ToPtr(testTextUnmarshaler(1))},
			},
			wantErr: false,
		},
		{
			name: "other filed have value",
			args: args{
				obj: &struct {
					Other string                 `form:"other"`
					Val   []*testTextUnmarshaler `form:"testTextUnmarshaler"`
				}{
					Other: "123",
				},
				values: map[string][]string{
					"testTextUnmarshaler": {"a"},
				},
				tag: "form",
			},
			want: &struct {
				Other string                 `form:"other"`
				Val   []*testTextUnmarshaler `form:"testTextUnmarshaler"`
			}{
				Other: "123",
				Val:   []*testTextUnmarshaler{lo.ToPtr(testTextUnmarshaler(1))},
			},
			wantErr: false,
		},
		{
			name: "time",
			args: args{
				obj: &struct {
					Val time.Time `form:"val"`
				}{
					Val: now,
				},
				values: map[string][]string{},
				tag:    "enum",
			},
			want: &struct {
				Val time.Time `form:"val"`
			}{
				Val: now,
			},
			wantErr: false,
		},
		{
			name: "time ptr",
			args: args{
				obj: &struct {
					Val *time.Time `form:"val"`
				}{
					Val: &now,
				},
				values: map[string][]string{},
				tag:    "enum",
			},
			want: &struct {
				Val *time.Time `form:"val"`
			}{
				Val: &now,
			},
			wantErr: false,
		},
		{
			name: "time nil",
			args: args{
				obj: &struct {
					Val *time.Time `form:"val"`
				}{
					Val: nil,
				},
				values: map[string][]string{},
				tag:    "enum",
			},
			want: &struct {
				Val *time.Time `form:"val"`
			}{
				Val: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mapping(tt.args.obj, tt.args.values, tt.args.tag, textUnmarshalerBinder); (err != nil) != tt.wantErr {
				t.Errorf("mapping() error = %+v, wantErr %+v", err, tt.wantErr)
			} else {
				// 比對物件一不一樣
				if !reflect.DeepEqual(tt.args.obj, tt.want) {
					t.Errorf("mapping() = %+v, want %+v", tt.args.obj, tt.want)
				}
			}
		})
	}
}
