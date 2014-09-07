package main

import (
	"encoding/json"
	"fmt"
	"github.com/shutej/elastigo/lib"
	"net/http"
	"os"
	"strings"
)

var es = elastigo.NewConn()

func requestHandler(res http.ResponseWriter, req *http.Request) {
	query := strings.ToLower(strings.TrimLeft(req.RequestURI, "/"))

	if query == "" {
		str := `{"error": "Please specify query path"}`
		res.Write([]byte(str))
		return
	}

	searchJson := `{
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

	searchJson = fmt.Sprintf(searchJson, query)

	out, err := es.Search("places", "place", nil, searchJson)

	if err != nil {
		str := fmt.Sprintf(`{"error": "%s"}`, err.Error())
		res.Write([]byte(str))
		return
	}

	if len(out.Suggestions["place-suggest"][0].Options) == 0 {
		res.Write([]byte("[]"))
		return
	}

	var places []json.RawMessage

	for _, v := range out.Suggestions["place-suggest"][0].Options {
		places = append(places, v.Payload)
	}

	bytes, err := json.Marshal(places)

	if err != nil {
		str := fmt.Sprintf(`{"error": "%s"}`, err)
		res.Write([]byte(str))
		return
	}

	res.Write(bytes)
}

func main() {
	es.Domain = os.Getenv("ELASTICSEARCH_URL")

	http.HandleFunc("/", requestHandler)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
