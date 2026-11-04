// Autocomplete: suggest completions for partial user input.
//
// Demonstrates: FindTopN with ASCIIOnly for high-throughput matching,
// comparing different metrics for autocomplete quality.
package main

import (
	"fmt"

	"github.com/jcoruiz/strsim"
)

// Programming languages and frameworks.
var dictionary = []string{
	"go", "golang", "gorilla", "gin", "gorm",
	"python", "pytorch", "pandas", "pip",
	"javascript", "java", "jackson", "jest",
	"typescript", "terraform", "tailwind",
	"rust", "ruby", "rails", "react", "redis",
	"kubernetes", "kafka", "kotlin",
	"docker", "django", "dart", "deno",
	"swift", "svelte", "scala", "spark",
	"vue", "vite", "vagrant", "vim",
	"angular", "ansible", "aws",
	"flask", "fastapi", "flutter",
	"next", "node", "nuxt", "nginx",
	"postgres", "prisma", "prometheus",
}

func main() {
	inputs := []string{"go", "jav", "reat", "typscrip", "kuber", "nde"}

	// Use ASCIIOnly JaroWinkler for maximum speed on ASCII identifiers.
	jw := &strsim.JaroWinklerMetric{
		BoostThreshold: 0.7,
		PrefixSize:     4,
		ASCIIOnly:      true,
	}

	fmt.Print("=== Autocomplete Suggestions ===\n\n")

	for _, input := range inputs {
		top := strsim.FindTopN(input, dictionary, 5, jw)

		fmt.Printf("  %q → ", input)
		for i, m := range top {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s (%.2f)", m.Value, m.Similarity)
		}
		fmt.Println()
	}
}
