package main

import (
	"animeta/lib/core/env"
	"animeta/lib/media"
	"animeta/lib/meta/providers/tvdb"
	"fmt"
)

func main() {
	env.Init()

	tvdbClient, err := tvdb.Authorize()
	if err != nil {
		fmt.Errorf("%w", err)
		return
	}

	media.NewClient(&tvdbClient)
	
	series, err := media.FetchSeries("252322")

	// movie, err := client.FetchMovie("791")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println(series)
	// fmt.Println(movie)
}
