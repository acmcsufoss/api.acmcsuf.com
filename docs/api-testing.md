# Testing the API

Once the api is running (via `air` or running it manually), you can use `xh` or `curl` to test the api from the command line. We recommend using `xh` over `curl`, and it's provided in the dev shell.  
If you prefer to use a GUI, Postman is a popular option.

## Using OpenAPI/Swagger Documentation

The API includes comprehensive OpenAPI2 (Swagger) documentation that provides an interactive interface for exploring and testing endpoints. When the server is running, you can access the Swagger UI at:

```
http://localhost:8080/docs
```

This is often the easiest way to explore the API and understand what endpoints are available without needing to dive into the source code. You can also use it to make test requests without needing to construct `curl` or `xh` commands manually.

## Using `xh` and `curl`
These examples use `xh`, but you can get the equivalent `curl` command by appending the `--curl` flag to any example command.

### Sending a POST request
This creates a resource in the database that we can later query for with a GET request.  
The OAuth2 middleware expects a `dev-token` passed with the Authorization middleware while in dev mode. See `oauth-authenticaion.md` for details.
```sh
xh post :8080/v1/board/officers -A bearer -a dev-token \
full_name="Bob" \
picture="example.com/picture.webp" \
discord="discord.com/users/bob" \
github="github.com/bob" \
uuid="123"
```

### Sending a GET request
This queries the database for the resource we just created:
```sh
xh :8080/v1/board/officers/123
```

There are lots of other routes, check out the OpenAPI docs for more info.

### Using fixtures
In the `fixtures` directory of this project, there's some JSON payloads that we can use for testing so we don't have to write them out each time. Feel free to create any that might be useful to yourself and/or the team.  
  
Here's an example using one to send a POST request:
```sh
xh post :8080/v1/events -A -a dev-token @fixtures/event_create.json
```

