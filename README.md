# Places

Autocomplete API endpoint for geographic places written in Golang with Elasticsearch backend.

## Installation

Create a new index:

```
curl -X PUT localhost:9200/places
```

Add mappings for index:

```
curl -X PUT localhost:9200/places/place/_mapping -d '{
  "place" : {
        "properties" : {
            "description" : { "type" : "string" },
            "suggest" : { "type" : "completion",
                          "index_analyzer" : "simple",
                          "search_analyzer" : "simple",
                          "payloads" : true
            }
        }
    }
}'
```

## Sample Data

Data should be added in the following format.

```
curl -X PUT 'localhost:9200/places/place/1?refresh=true' -d '{
    "description" : "Hollywood, Los Angeles, CA, United States",
    "suggest" : {
        "input": [ "Hollywood", "Los Angeles", "CA", "United States" ],
        "output": "Hollywood, Los Angeles, CA, United States",
        "payload" : { 
          "description" : "Hollywood, Los Angeles, CA, United States",
          "placeId" : "1",
          "placeSource" : "Neighborhood",
          "timeZoneId" : "America/Los_Angeles",
          "timeZoneName" : "Pacific Daylight Time",
          "latitude" : 34.0500,
          "longitude" : 118.2500,
        },
        "weight" : 50
    }
}'
```

## Run the application

Launch the application locally with ENV vars on command line. When deploying these should be set as environment variables.

```
PORT=5001 ELASTICSEARCH_URL=localhost go run places.go
```

[Bam!](http://localhost:5001/hol)
