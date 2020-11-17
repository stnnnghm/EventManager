# create an event
curl -X POST \
  http://localhost:8080/api/v1/event \
  -H 'content-type: application/json' \
  -d '{
    "name": "Programmingella",
    "description": "quick brown fox jumps over the lazy dog",
    "slot": {
        "start_time": "2020-12-11T09:00:00+05:30",
        "end_time": "2020-12-11T15:00:00+05:30"
    },
    "website": "https://programmingella.com",
    "address": "Mexnte"
}';

# list one events
curl -X GET 'http://localhost:8080/api/v1/events?limit=1'

# or try getting one using an id (note: your id will be something different)
curl -X GET 'http://localhost:8080/api/v1/event?id=1599425402-0970640120-4671038529'