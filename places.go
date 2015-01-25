package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/shutej/elastigo/lib"
)

var es = elastigo.NewConn()

type Place struct {
	Description    string  `json:"description"`
	PlaceId        string  `json:"placeId"`
	PlaceType      string  `json:"placeType"`
	TimeZoneId     string  `json:"timeZoneId"`
	TimeZoneName   string  `json:"timeZoneName"`
	TimeZoneOffset int     `json:"timeZoneOffset"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}

func (p *Place) CalculateTimeZoneOffset() {
	location, err1 := time.LoadLocation(p.TimeZoneId)

	if err1 != nil {
		p.TimeZoneOffset = 0
		return
	}

	_, offset := time.Now().In(location).Zone()

	p.TimeZoneOffset = offset
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(strings.TrimLeft(r.RequestURI, "/"))

	if query == "" {
		http.Error(w, "Please specify location name", http.StatusInternalServerError)
		return
	}

	searchJson := `{
		"from" : 0,
		"size" : 5,
    "suggest" : {
      "place-suggest" : {
        "text" : "%s",
        "completion" : {
          "field" : "suggest",
          "fuzzy" : {
            "fuzziness" : 2
          }
        }
      }
    }
  }`

	qry := fmt.Sprintf(searchJson, query)
	out, err := es.Search("places", "place", nil, qry)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(out.Suggestions["place-suggest"][0].Options) == 0 {
		w.Write([]byte("[]"))
		return
	}

	var places []Place

	for _, v := range out.Suggestions["place-suggest"][0].Options {
		var p Place
		json.Unmarshal(v.Payload, &p)
		p.CalculateTimeZoneOffset()
		places = append(places, p)
	}

	bytes, err := json.Marshal(places)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}

func main() {
	es.Domain = os.Getenv("ELASTICSEARCH_URL")
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	if err != nil {
		panic(err)
	}
}
