package main

import "testing"

func TestHello(t *testing.T) {
	testCases := []struct {
		desc string
		name string
		lang Language
		want string
	}{
		{
			desc: "it should return 'Hello, {name}!', when name is supplied",
			name: "Erik",
			want: "Hello, Erik!",
		},
		{
			desc: "it should return 'Hello, World!', when empty name is supplied",
			want: "Hello, World!",
		},
		{
			desc: "it should return 'Hola, {name}!', when Spanish is supplied",
			name: "Erik",
			lang: Spanish,
			want: "Hola, Erik!",
		},
		{
			desc: "it should return 'Halo, {name}!', when Bahasa is supplied",
			name: "Erik",
			lang: Bahasa,
			want: "Halo, Erik!",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			got := Hello(testCase.name, testCase.lang)
			assertCorrentMessage(t, got, testCase.want)
		})
	}
}

func assertCorrentMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
