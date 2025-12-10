# API Versioning

This document explains the API versioning strategy.

All API endpoints are versioned to allow for future changes without breaking existing clients. The version is specified in the URL, for example, `/v1/events`.

When adding new endpoints, they should be added to the latest version group in `internal/api/routes/v1.go`.

If a breaking change is required, a new version of the API should be created. This would involve:

1. Creating a new routes file, e.g., `internal/api/routes/v2.go`.
2. Creating a new router group, e.g., `v2 := router.Group("/v2")`.
3. Adding the new and modified endpoints to the `v2` group.
