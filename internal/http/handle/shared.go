package handle

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type R map[string]interface{}

func writeJson(w http.ResponseWriter, j interface{}) {
	bytes, err := json.Marshal(j)
	if err != nil {
		fmt.Fprint(w, "internal error")
		return
	}

	fmt.Fprint(w, string(bytes))
}
