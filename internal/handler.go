package internal

import (
	"encoding/json"
	"io"
)

type Handler struct{}

func (h Handler) ReadJSON(r io.Reader, target interface{}) error {
	jsonDecoder := json.NewDecoder(r)
	jsonDecoder.DisallowUnknownFields()
	return jsonDecoder.Decode(&target)
}

func (h Handler) WriteJSON(w io.Writer, target interface{}) error {
	jsonEncoder := json.NewEncoder(w)
	return jsonEncoder.Encode(target)
}
