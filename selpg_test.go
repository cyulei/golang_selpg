package main

import "testing"

func TestUsage(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Usage()
		})
	}
}

func Test_process_args(t *testing.T) {

	tests := []struct {
		name string
		args []string
	}{
		// TODO: Add test cases.
		{
			name: "Nil options",
			args: []string{"selpg", "-s1", "-e1", "test.txt"},
		}, {
			name: "Nil",
			args: []string{"selpg", "-s1", "-e2", "test.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			process_args(tt.args)
		})
	}
}

func Test_process_input(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			process_input()
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
