package main

import (
	"reflect"
	"testing"
)

func Example_main() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {
	tests := map[string]struct {
		lang language
		want string
	}{
		"Empty": {
			lang: "",
			want: `unsupported language: ""`,
		},
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"Swahili": {
			lang: "swa",
			want: `unsupported language: "swa"`,
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Greek": {
			lang: "el",
			want: "Χαίρετε Κόσμε",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected %q, got %q", tc.want, got)
			}
		})
	}
}
