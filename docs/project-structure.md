# Project Structure

This document provides an overview of the project's directory structure and the responsibility of each main package.

- `cmd/`: Main applications for the project.
   - `acmcsuf-api/`: The main entry point for the API server. It handles command-line flags and graceful shutdown.
   - `acmcsuf-cli/`: A command-line interface for interacting with the API.
- `internal/`: Private application and library code.
   - `api/`: Contains all the API server logic.
      - `server.go`: Initializes and starts server.
      - `routes/`: API route definitions.
      - `handlers/`: Gin handlers for processing HTTP requests.
      - `services/`: Business logic for the API.
      - `middleware/`: Contains middleware like the rate limiter and OAuth implementation.
      - `dbmodels/`: sqlc generated database models. Do not edit manually.
      - `config/`: Loads environment vars and defines defaults.
   - `cli/`: Logic for the CLI application.
- `fixtures/`: JSON payloads for testing.
- `nix/`: Nix dev shell and package definitions.
- `utils/`: Shared utility functions.
