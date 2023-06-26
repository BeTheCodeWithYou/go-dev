package main

import "testing"

func ExampleMain() {
	main()
	//Output:
	// hello world
}

func TestGreet_English(t *testing.T) {
	inputLanguage := "en"
	expectedGreeting := "hello world"

	greeting := greet(language(inputLanguage))

	if greeting != expectedGreeting {
		t.Errorf("Expected %q, got %q", expectedGreeting, greeting)
	}

}

func TestGreet_French(t *testing.T) {
	inputLanguage := "fr"
	expectedGreeting := "Bonjour le monde"

	greeting := greet(language(inputLanguage))

	if greeting != expectedGreeting {
		t.Errorf("Expected %q, got %q", expectedGreeting, greeting)
	}
}

func TestGreek(t *testing.T) {
	type scenario struct {
		lang             language
		expectedGreeting string
	}

	var tests = map[string]scenario{
		"english": {
			lang:             "en",
			expectedGreeting: "hello world",
		},
		"french": {
			lang:             "fr",
			expectedGreeting: "Bonjour le monde",
		},
	}

	for name, tc := range tests {

		t.Run(name, func(t *testing.T) {
			greeting := greet(tc.lang)

			if greeting!=tc.expectedGreeting {
				t.Errorf("expected %q, got %q",  tc.expectedGreeting, greeting)
			}
		})

	}
}
