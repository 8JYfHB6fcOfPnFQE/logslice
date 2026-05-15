package parser

import (
	"testing"
	"time"
)

type tsCase struct {
	name    string
	line    string
	wantErr bool
	wantTS  string // RFC3339 expected value; empty means skip exact check
}

var tsCases = []tsCase{
	{
		name: "RFC3339",
		line: `2024-03-15T08:22:01Z INFO server started`,
		wantTS: "2024-03-15T08:22:01Z",
	},
	{
		name: "datetime with space",
		line: `2024-03-15 08:22:01 ERROR disk full`,
		wantTS: "2024-03-15T08:22:01Z",
	},
	{
		name: "datetime with milliseconds comma",
		line: `2024-03-15 08:22:01,123 WARN high memory`,
		wantTS: "2024-03-15T08:22:01Z",
	},
	{
		name: "nginx combined log",
		line: `127.0.0.1 - - [15/Mar/2024:08:22:01 +0000] "GET / HTTP/1.1" 200 512`,
		wantTS: "2024-03-15T08:22:01Z",
	},
	{
		name: "syslog format",
		line: `Mar 15 08:22:01 hostname sshd[123]: Accepted`,
		wantErr: false, // year will be 0000 but parse succeeds
	},
	{
		name: "no timestamp",
		line: `this line has no timestamp at all`,
		wantErr: true,
	},
}

func TestExtractTimestamp(t *testing.T) {
	for _, tc := range tsCases {
		t.Run(tc.name, func(t *testing.T) {
			got, offset, err := ExtractTimestamp(tc.line)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got time %v at offset %d", got, offset)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if offset < 0 {
				t.Fatalf("expected non-negative offset, got %d", offset)
			}
			if tc.wantTS != "" {
				want, _ := time.Parse(time.RFC3339, tc.wantTS)
				if !got.UTC().Truncate(time.Second).Equal(want.UTC()) {
					t.Errorf("timestamp mismatch: got %v, want %v", got, want)
				}
			}
		})
	}
}
