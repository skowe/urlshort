package jsonshort

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type handler struct {
	data map[string]string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.RequestURI()
	log.Println(uri)
	if i, ok := h.data[uri]; ok {
		c := i
		fmt.Println(c)
		http.Redirect(w, r, c, http.StatusSeeOther)
	} else {
		w.Write([]byte("Hi there!" + uri))
	}
}

func GetHandler(path string) (http.Handler, error) {
	data, err := jsonToMap(path)
	log.Println(data)
	if err != nil {
		return nil, err
	}
	red := handler{data: data}
	return &red, nil
}

func jsonToMap(path string) (map[string]string, error) {
	fileBytes, err := readJson(path)
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)

	dec := json.NewDecoder(fileBytes)
	err = dec.Decode(&data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func readJson(path string) (*bytes.Buffer, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	d, err := io.ReadAll(f)

	if err != nil {
		return nil, err
	}
	data := bytes.NewBuffer(d)
	return data, nil
}
