package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/lkhwa/info344-in-class/zipsvr/models"
)

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//URL: /zips/city-name, e.g. /zips/seattle
	//need to grab the last token from the URL + know what's the path prefix
	cityName := r.URL.Path[len(ch.PathPrefix):] //isolates string after path prefix
	cityName = strings.ToLower(cityName)
	if len(cityName) == 0 { //error handling
		http.Error(w, "please provide a city name", http.StatusBadRequest)
		return
	}

	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(accessControlAllowOrigin, "*")
	zips := ch.Index[cityName]
	json.NewEncoder(w).Encode(zips)
}
