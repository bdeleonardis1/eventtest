package main_test

import (
	"os/exec"
	"strings"
	"testing"

	eventtestapi "github.com/bdeleonardis1/eventtest/api"
	"github.com/bdeleonardis1/eventtest/events"
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
		expectedEvents []*events.Event
	}{
		{
			input: "1",
			expected: "1 is an odd number",
			expectedEvents: []*events.Event{
				events.NewEvent("1Optimization"),
			},
		},
		{
			input: "2",
			expected: "2 is an even number",
			expectedEvents: []*events.Event{
				events.NewEvent("OptimizedSingleDigit"),
			},
		},
		{
			input: "11",
			expected: "11 is an odd number",
			expectedEvents: []*events.Event{
				events.NewEvent("convertToNumber"), events.NewEvent("Modding"),
			},
		},
		{
			input: "-3",
			expected: "-3 is an odd number",
			expectedEvents: []*events.Event{
				events.NewEvent("convertToNumber"), events.NewEvent("OptimizedNegativeSingleDigit"),
			},
		},
		{
			input: "-4",
			expected: "-4 is an even number",
			expectedEvents: []*events.Event{
				events.NewEvent("convertToNumber"), events.NewEvent("OptimizedNegativeSingleDigit"),
			},
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

			eventtestapi.ExpectExactEvents(t, tc.expectedEvents)
		})
	}
}

func TestExpectEventsDemo(t *testing.T) {
	eventtestapi.ClearEvents()

	cmd := exec.Command("./sampleprogram")
	cmd.Stdin = strings.NewReader("19")
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}

	eventtestapi.ExpectEvents(t, []*events.Event{events.NewEvent("convertToNumber"), events.NewEvent("TheVeryEnd")}, eventtestapi.Ordered)
	eventtestapi.ExpectEvents(t, []*events.Event{events.NewEvent("TheVeryEnd"), events.NewEvent("convertToNumber")}, eventtestapi.Unordered)

	eventtestapi.UnexpectedEvents(t, []*events.Event{events.NewEvent("1Optimization"), events.NewEvent("OptimizedNegativeSingleDigit")})

	// should fail.
	eventtestapi.ExpectEvents(t, []*events.Event{events.NewEvent("TheVeryEnd"), events.NewEvent("convertToNumber")}, eventtestapi.Ordered)
}




