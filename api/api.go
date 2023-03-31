package api

import (
	"fmt"
	"log"
	"net/http"
	"refes/api/ds"
)

func Init(address *string, port *int) error {
	http.HandleFunc("/", ds.HandleRequest)

	log.Printf("INFO: server starting on %s:%d\n", *address, *port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", *address, *port), nil)
	if err != nil {
		return err
	}

	return nil
}
