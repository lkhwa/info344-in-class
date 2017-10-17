package handlers

	"encoding/json"
	"net/http"
	"strings"

	"github.com/lkhwa/info344-in-class/zipsvr/models"
)

type CityHandler struct {
		http.Error(w, "please provide a city name", http.StatusBadRequest)
		return
	}

	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(accessControlAllowOrigin, "*")
	zips := ch.Index[cityName]
=======
//ServeHTTP handles HTTP requests for the CityHandler. This is a method
//of the CityHandler struct defined above. Methods in Go use a receiver
//parameter defined on the left, which will be an instance of the struct.
//The receiver parameter is exactly like the `this` pointer in Java, just
//more explicitly defined. For more details on receiver parameters, see
//https://drstearns.github.io/tutorials/golang/#secreceivers
func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// URL: /zips/city-name
	//slice off just the city name from the resource path
	//since we know the length of the PathPrefix, we can slice off
	//just the characters that follow that PathPrefix
	cityName := r.URL.Path[len(ch.PathPrefix):]

	//convert the city name to lower case since the ZipIndex map
	//keys are all lower-cased as well
	cityName = strings.ToLower(cityName)

	//if the city name is zero-length respond with an error
	if len(cityName) == 0 {
		//the http.Error() method writes an error message to the response
		//and sets the HTTP status code to the value of the third parameter
		http.Error(w, "please provide a city name", http.StatusBadRequest)
		//since http.Error() writes a response, we should return to
		//stop processing this request.
		return
	}

	//add the header `Content-Type: application/json`
	w.Header().Add(headerContentType, contentTypeJSON)
	//add the CORS header `Access-Control-Allow-Origin: *`
	//see https://drstearns.github.io/tutorials/cors/
	w.Header().Add(headerAccessControlAllowOrigin, "*")

	//get the ZipSlice for the requested city name
	zips := ch.Index[cityName]
	//write that slice to the response, encoded as JSON
>>>>>>> ee296d65393eb82cb7090e69ee848d296c254143
	json.NewEncoder(w).Encode(zips)
}
