# Configuration via Environment Variables

See [config.go](../internal/api/config/config.go) for more information/default
values.

- `GIN_MODE`: One of `debug`, `release`, or `test`. This affects what and how
many HTTP middleware logs appear. Separate from manual logging.
    - Default: `debug` (change in production)
- `ENV`: One of `production` or `development`. While in development mode,
authentication with Discord OAuth2 is bypassed. You may need to change this value to
`production` if testing Auth.
    - Default: `development` (change in production)
- `PORT`: Port to run on.
    - Default: `8080`
- `DATABASE_URL`: The path to the SQLite database file. Takes the format
`file:path/to/database.db`. Can further configure using query params. Example:
`mode=rwc` means read-write and create if doesn't exist. `cache=shared` is more
memory efficient than the default of `cache=private`.
    - Default: `file:dev.db?cache=shared&mode=rwc`
- `TRUSTED_PROXIES`: Used in [server.go](../internal/api/server.go) in Gin's
`SetTrustedProxies` setting. Security mechanism to prevent IP spoofing when the
API sits behind a reverse proxy in production.
    - Default: `127.0.0.1/32` (change in production)
- `ALLOWED_ORIGINS`: To be used with CORS middleware. Controls which web origins
  are allowed to make cross-origin requests to the API.
    - Default: `*` (change in production)
