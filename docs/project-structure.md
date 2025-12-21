# Project Structure

This document provides an overview of the project's directory structure and the responsibility of each main package.

- `cmd/`: Main applications for the project.
  - `acmcsuf-api/`: The main entry point for the API server. It handles command-line flags and graceful shutdown.
  - `acmcsuf-cli/`: A command-line interface for interacting with the API.
- `internal/`: Private application and library code.
  - `api/`: Contains all the API server logic.
    - `handlers/`: Gin handlers for processing HTTP requests.
    - `routes/`: API route definitions.
    - `services/`: Business logic for the API.
  - `cli/`: Logic for the CLI application.
  - `db/`: Database-related code, including schema, queries, and models.
- `developer-docs/`: Documentation for developers.
- `docs/`: Swagger/OpenAPI documentation.
- `fixtures/`: JSON payloads for testing.
- `nix/`: Nix dev shell and package definitions.
- `utils/`: Shared utility functions.
