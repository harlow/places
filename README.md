# Places

Autocomplete API endpoint for geographic places written in Golang with Elasticsearch backend.

## Elasticsearch Index

Create a new index:

```
curl -X PUT localhost:9200/places
```

Add mappings for index:

```
curl -X PUT localhost:9200/places/place/_mapping -d '{
  "place" : {
    "properties" : {
      "suggest" : {
        "type" : "completion",
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
  "suggest" : {
    "input": [ "Hollywood, Los Angeles, CA, United States" ],
    "payload" : {
      "description" : "Hollywood, Los Angeles, CA, United States",
      "placeId" : "ChIJD79iBh9u5kcRyJsMaMOCCwQ",
      "placeType" : "Neighborhood",
      "timeZoneId" : "America/Los_Angeles",
      "timeZoneName" : "Pacific Daylight Time",
      "latitude" : 34.1000,
      "longitude" : -118.3333
      },
    "weight" : 50
  }
}'

curl -X PUT 'localhost:9200/places/place/2?refresh=true' -d '{
  "suggest" : {
    "input": [ "Hawthorne, Los Angeles, CA, United States" ],
    "payload" : {
      "description" : "Hawthorne, Los Angeles, CA, United States",
      "placeId" : "2",
      "placeType" : "Neighborhood",
      "timeZoneId" : "America/Los_Angeles",
      "timeZoneName" : "Pacific Daylight Time",
      "latitude" : 34.0500,
      "longitude" : 118.2500
    },
    "weight" : 35
  }
}'

curl -X PUT 'localhost:9200/places/place/3?refresh=true' -d '{
  "suggest" : {
    "input": [ "Howard Springs, CA, United States" ],
    "payload" : {
      "description" : "Howard Springs, CA, United States",
      "placeId" : "3",
      "placeType" : "Neighborhood",
      "timeZoneId" : "America/Los_Angeles",
      "timeZoneName" : "Pacific Daylight Time",
      "latitude" : 38.8582349,
      "longitude" : -122.6747093
    },
    "weight" : 35
  }
}'
```

## Run the application

Launch the application locally with ENV vars on command line. When deploying these should be set as environment variables.

```
PORT=5001 ELASTICSEARCH_URL=localhost go run main.go
```

[Bam!](http://localhost:5001/ho)
