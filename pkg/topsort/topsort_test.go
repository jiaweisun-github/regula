package topsort

import (
	"testing"
)

func Test_topsort(t *testing.T) {
	type Test struct {
		name         string
		entries      []Key
		dependencies [][]Key
		expectError  bool
	}

	tests := []Test{
		{
			name:    "Simple",
			entries: []Key{"one", "two", "three"},
			dependencies: [][]Key{
				{"one", "two"},
				{"two", "three"},
			},
		},
		{
			name:    "Cycle",
			entries: []Key{"one", "two", "three"},
			dependencies: [][]Key{
				{"one", "two"},
				{"two", "three"},
				{"three", "one"},
			},
			expectError: true,
		},
		{
			name:    "Cliques",
			entries: []Key{"one", "two", "three", "four", "five"},
			dependencies: [][]Key{
				{"one", "two"},
				{"two", "three"},
				{"five", "four"},
			},
		},
		{
			name:    "Non-existing dependencies",
			entries: []Key{"one", "two", "three"},
			dependencies: [][]Key{
				{"one", "een"},
				{"two", "twee"},
				{"three", "one"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Build graph
			g := make(Graph)
			for _, k := range tc.entries {
				g[k] = []Key{}
			}
			for _, dependency := range tc.dependencies {
				k := dependency[0]
				g[k] = append(g[k], dependency[1])
			}

			// Run topsort
			sorted, err := Topsort(g)

			// Check error
			if tc.expectError {
				if err == nil {
					t.Fatal("Expected error but succeeded")
				}
			} else {
				// Check that every item appears after all its dependencies
				indices := map[string]int{}
				for i, k := range sorted {
					indices[k] = i
				}
				for _, dependency := range tc.dependencies {
					k := dependency[0]
					l := dependency[1]
					i := indices[k]
					if j, ok := indices[l]; ok {
						if i <= j {
							t.Fatalf("Expected '%s' after '%s' but they have indices %d and %d", k, l, i, j)
						}
					}
				}
			}
		})
	}
}