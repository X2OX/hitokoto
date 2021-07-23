package hitokoto

import (
	"encoding/json"
	"net/http"
	"strings"

	"go.x2ox.com/hitokoto/data"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	newResult().Output(w, strings.ToLower(r.URL.Query().Get("type")))
}

type Result struct {
	Text string `json:"text"`
}

func newResult() Result { return Result{Text: data.Random()} }
func (r Result) Output(w http.ResponseWriter, outputType string) {
	switch outputType {
	case "json":
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		_ = json.NewEncoder(w).Encode(r)
	default:
		_, _ = w.Write([]byte(r.Text))
	}
}
