package algoexpert

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			"valid input",
			"abcdcba",
			true,
		},
		{
			"invalid input",
			"shdhds",
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if received := IsPalindrome(test.str); received != test.want {
				t.Error("failed test")
			}
		})
	}
}

func TestCaesarCipherEncryptor(t *testing.T) {
	type args struct {
		str string
		key int
	}

	tests := []struct {
		name  string
		input args
		want  string
	}{
		{
			name: "test one",
			input: args{
				"abc",
				0,
			},
			want: "abc",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := CaesarCipherEncryptor(test.input.str, test.input.key); result != test.want {
				t.Error("invalid test ")
			}
		})
	}

}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want [][]string
	}{
		{
			"valid input",
			[]string{"aca", "bba"},
			[][]string{{"aca"}, {"bba"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := GroupAnagrams(test.args); !reflect.DeepEqual(result, test.want) {
				t.Error("failed test")
			}
		})
	}
}
