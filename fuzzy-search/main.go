// Fuzzy search: find the best matches for a query in a product catalog.
//
// Demonstrates: FindTopN, Metric interface, JaroWinkler vs Levenshtein comparison.
package main

import (
	"fmt"

	"github.com/jcoruiz/strsim"
)

var catalog = []string{
	"Wireless Bluetooth Headphones",
	"USB-C Charging Cable",
	"Mechanical Gaming Keyboard",
	"Portable External SSD 1TB",
	"Noise Cancelling Earbuds",
	"4K Ultra HD Monitor",
	"Ergonomic Office Chair",
	"Laptop Cooling Pad",
	"Smart Home Speaker",
	"Adjustable Standing Desk",
	"Wireless Gaming Mouse",
	"LED Desk Lamp",
	"Webcam HD 1080p",
	"Portable Power Bank 20000mAh",
	"Thunderbolt Docking Station",
}

func main() {
	query := "wireles headphone"

	fmt.Printf("Query: %q\n\n", query)

	// Compare two metrics on the same query.
	metrics := []struct {
		name   string
		metric strsim.Metric
	}{
		{"Jaro-Winkler", strsim.NewJaroWinkler()},
		{"Levenshtein", strsim.NewLevenshtein()},
		{"Dice (bigram)", &strsim.NgramMetric{Size: 2}},
	}

	for _, m := range metrics {
		fmt.Printf("--- %s ---\n", m.name)

		top := strsim.FindTopN(query, catalog, 3, m.metric)
		for i, match := range top {
			fmt.Printf("  #%d  %.4f  %s\n", i+1, match.Similarity, match.Value)
		}

		fmt.Println()
	}
}
