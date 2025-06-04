package main

import (
	"fmt"

	"github.com/Sunhill666/goalex/pkg/core"
)

func main() {
	query := core.QueryParams{
		Filter: map[string]any{
			"institutions.country_code": "fr+gb",
		},
	}
	fmt.Println(query.ToQuery().Encode())
}
