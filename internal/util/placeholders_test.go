package util

import "testing"

func Test_GetPlaceholders(t *testing.T) {
    type args struct {
        n   int
        pos []int
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {name: "zero", args: args{n: 0}, want: "()"},
        {name: "one", args: args{n: 1}, want: "(?)"},
        {name: "two", args: args{n: 2}, want: "(?,?)"},
        {name: "three", args: args{n: 3}, want: "(?,?,?)"},
        {name: "four", args: args{n: 4}, want: "(?,?,?,?)"},
        {name: "four_with_invalid_bin", args: args{n: 4, pos: []int{6}}, want: "(?,?,?,?)"},
        {name: "four_with_bin", args: args{n: 4, pos: []int{1}}, want: "(?,UUID_TO_BIN(?),?,?)"},
        {name: "five_with_two_bin", args: args{n: 5, pos: []int{0, 2}}, want: "(UUID_TO_BIN(?),?,UUID_TO_BIN(?),?,?)"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := GetPlaceholders(tt.args.n, tt.args.pos...); got != tt.want {
                t.Errorf("getPlaceholders() = %v, want %v", got, tt.want)
            }
        })
    }
}
