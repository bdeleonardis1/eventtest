package main_test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	eventtestapi "github.com/bdeleonardis1/eventtest/api"
	"github.com/bdeleonardis1/eventtest/server"
)

const (
	expectedBase = "Enter a number 1-5: "
)

func TestParity(t *testing.T) {
	go server.Serve()

	testCases := []struct{
		input string
		expected string
	}{
		{
			input: "1",
			expected: "1 is an odd number",
		},
		{
			input: "2",
			expected: "2 is an even number",
		},
		{
			input: "11",
			expected: "11 is an odd number",
		},
		{
			input: "-3",
			expected: "-3 is an odd number",
		},
		{
			input: "-4",
			expected: "-4 is an even number",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			eventtestapi.ClearEvents()

			cmd := exec.Command("./sampleprogram")
			cmd.Stdin = strings.NewReader(tc.input)
			out, err := cmd.Output()
			if err != nil {
				t.Fatal(err)
			}
			outString := strings.TrimSpace(string(out))

			if outString != expectedBase + tc.expected {
				t.Errorf("expected '%v', but got '%v'", expectedBase + tc.expected, outString)
			}

			events, err := eventtestapi.GetEvents()
			if err != nil {
				t.Errorf("error getting events: %v", err)
			}
			eventsStr := ""
			for _, event := range events {
				eventsStr += event.Name + ", "
			}
			eventsStr = eventsStr[:len(eventsStr)-2]
			fmt.Printf("events (%v): %v \n", tc.input, eventsStr)
		})
	}
}


