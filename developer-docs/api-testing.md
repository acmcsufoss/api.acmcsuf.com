# Testing the API

Once the api is running (via `air` or running it manually), you can use `xh` or `curl` to test the api from the command line. You probably already have `curl` installed, but `xh` is provided in the dev shell.
We recommend using `xh` as it is the more modern option.
If you prefer to use a GUI instead, Postman is a popular option.

## Using Swagger Documentation

The API includes comprehensive Swagger documentation that provides an interactive interface for exploring and testing endpoints. When the server is running, you can access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

The Swagger documentation includes:
- Complete API endpoint documentation for all available routes
- Interactive forms to test API calls directly from the browser
- Request/response schemas and examples
- Parameter descriptions and validation rules

This is often the easiest way to explore the API and understand what endpoints are available without needing to dive into the source code. You can also use it to make test requests without needing to construct `curl` or `xh` commands manually.

## Using `xh` and `curl`

### Sending a GET request
The following requests will request an event called "event1" from the APIs `events` endpoint. You'll probably have to create this event yourself if you want to use this exact command (see the POST section below).
```sh
# assuming the server is running on localhost:8080
xh :8080/events/event1

# or with curl:
curl localhost:8080/events/event1
```
Both `xh` and `curl` send a GET request by default, so you don't need to specify which method to use.  
Note: `curl` won't give you formatted output by default, so it might be helpful to pipe to `jq` (also provided in dev shell) to make it look pretty.
```sh
curl localhost:8080/events/event1 | jq
```


### Sending a POST request
```sh
xh POST :8080/events \
uuid="event1" \
location="tsu" \
start_at="1712851200000" \
end_at="1712851200000" \
is_all_day:=false \
host="ACM"

# same thing with curl:
curl -X POST localhost:8080/events \
-H 'content-type: application/json' \
-H 'accept: application/json' \
-d '{"uuid":"event1","location":"tsu","start_at":"1712851200000","end_at":"1712851200000","is_all_day":false,"host":"ACM"}'
```

As you can see, `xh` provides a more concise way to do the same thing.  

### Using fixtures
In the `fixtures` directory of this project, there's some JSON payloads that we can use for testing so we don't have to write them out each time. Feel free to create any that might be useful to yourself and/or the team.  
  
Here's an example using one to send a POST request:
```sh
xh POST :8080/events @fixtures/event.json

# same thing with curl:
curl -X POST localhost:8080/events \
-H "Content-Type: application/json" \
-d @payload.json
```

