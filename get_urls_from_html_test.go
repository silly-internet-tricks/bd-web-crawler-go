package main

import (
	"reflect"
	"testing"
)

const example1 = `
<html>
	<head>
		<title>Hello I am an Example!</title>
	</head>
	<body>
		<div>
			<a href="www.google.com">
				Try googling it!
			</a>
			<a href="https://www.google.com">
				haha psych
			</a>
		</div>
	</body>
</html>
`

const example2 = `
<html>
	<body>
		<a href="https://x.co">
			hello from x
			<a href="https://y.co">
				y is better btw
				<a href="https://z.co">
					z can't be beat!
				</a>
			</a>
		</a>
	</body>
</html>
`

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name       string
		htmlBody   string
		rawBaseURL string
		expected   []string
		err        error
	}{
		{
			name: "appends a path as expected",

			htmlBody:   example1,
			rawBaseURL: "localhost",
			expected:   []string{"localhost/www.google.com", "https://www.google.com"},
			err:        nil,
		},
		{
			name: "absolute and relative URLs",
			htmlBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			rawBaseURL: "https://blog.boot.dev",
			expected:   []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
			err:        nil,
		},
		{
			name:       "anchors nested within one another",
			htmlBody:   example2,
			rawBaseURL: "https://google.com",
			expected:   []string{"https://x.co", "https://y.co", "https://z.co"},
			err:        nil,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.htmlBody, tc.rawBaseURL)
			if tc.err != nil && err == nil {
				t.Errorf("Test %v - '%s' FAIL: missing error: %v", i, tc.name, tc.err)
				return
			}

			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
