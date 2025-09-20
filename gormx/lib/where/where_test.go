package wh_test

import (
	"testing"

	ope "github.com/skyrocket-qy/gox/gormx/lib/operator"
	wh "github.com/skyrocket-qy/gox/gormx/lib/where"
)

func TestB(t *testing.T) {
	type args struct {
		field string
		oper  string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "equal",
			args: args{field: "name", oper: ope.Eq},
			want: "name = ?",
		},
		{
			name: "between",
			args: args{field: "age", oper: ope.Bt},
			want: "age BETWEEN ? AND ?",
		},
		{
			name: "not between",
			args: args{field: "created_at", oper: ope.NBt},
			want: "created_at NOT BETWEEN ? AND ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wh.B(tt.args.field, tt.args.oper); got != tt.want {
				t.Errorf("B() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSub(t *testing.T) {
	type args struct {
		field string
		oper  string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "in",
			args: args{field: "id", oper: ope.In},
			want: "id IN (?)",
		},
		{
			name: "between",
			args: args{field: "price", oper: ope.Bt},
			want: "price BETWEEN (?) AND (?)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wh.BSub(tt.args.field, tt.args.oper); got != tt.want {
				t.Errorf("BSub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDesc(t *testing.T) {
	type args struct {
		column string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{column: "name"},
			want: "name DESC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wh.Desc(tt.args.column); got != tt.want {
				t.Errorf("Desc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAsc(t *testing.T) {
	type args struct {
		column string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{column: "name"},
			want: "name ASC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wh.Asc(tt.args.column); got != tt.want {
				t.Errorf("Asc() = %v, want %v", got, tt.want)
			}
		})
	}
}
