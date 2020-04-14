package handler

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
)

func CustomUrls(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println(path)
	redirect, err := ParseYaml()
	if err != nil {
		Fallback(w, r)
	}
	if url, ok := redirect[path]; ok {
		http.Redirect(w, r, url, 303)
	} else {
		Fallback(w, r)
	}

}

func Fallback(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)
	w.Write([]byte("Unfortunately we do not offer a redirect for this endpoint"))

}


func ParseYaml() (map[string]string, error){
	yamlFile, err := ioutil.ReadFile("./url-paths.yaml")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result := make(map[string]string)
	err = yaml.Unmarshal(yamlFile, &result)
	fmt.Println(result)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}


