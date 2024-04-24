package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return
	}

	fmt.Println(body) // Optional: Print the request body
	errFound := json.Unmarshal([]byte(body), x)
	if errFound != nil {
		panic(errFound)
	}

}
