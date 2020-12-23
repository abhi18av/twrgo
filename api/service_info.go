package api

import (
	"fmt"
	"github.com/imroc/req"
	"log"
)

func ServiceInfo(towerApiEndpoint string) {
	ENDPOINT := towerApiEndpoint + "service-info/"
	r, err := req.Get(ENDPOINT)
	if err != nil {
		log.Fatal(err)
	}

	// use req package to initiate request.
	fmt.Println( r)
}
