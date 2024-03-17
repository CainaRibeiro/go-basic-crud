package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DecoderJSON(r *http.Request, i interface{}) (interface{}, error) {
	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	err := dec.Decode(&i)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return i, nil
}
