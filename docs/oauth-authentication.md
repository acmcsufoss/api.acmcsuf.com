# OAuth Authentication

## Overview

The API implements a role-based access control (RBAC) system that:
1. Authenticates users via Discord OAuth2
2. Verifies Discord server membership
3. Checks user roles against required permissions
4. Caches role information to avoid hitting Discord's rate limits

## Architecture

### Server-Side: API Middleware

The API uses Discord OAuth2 middleware defined in [`internal/api/middleware/oauth.go`](../internal/api/middleware/oauth.go) to protect state-changing `/v1` routes.

**How it works:**

1. **Authorization Header**: The middleware expects requests to include:
   ```
   Authorization: Bearer <discord_access_token>
   ```

2. **Token Validation**: The middleware validates the token by:
   - Making a request to Discord's API to fetch user information
   - Verifying the user is a member of the configured Discord guild/server
   - Checking if the user has the required role

3. **Role-Based Access Control**:
   - Roles are mapped in `RoleMap` (e.g., Discord role ID `1445971950584205554` → `"Board"`)
   - The middleware checks if the user has the required role for the endpoint
   - Role hierarchy: `President` role also grants `Board` access

4. **Caching**:
   - Role information is cached for 5 minutes to prevent rate limit issues
   - Cache key is the Authorization header value
   - Expired cache entries are automatically removed

5. **Development Mode**:
   - When `ENV=development`, the special token `Bearer dev-token` bypasses authentication
   - This allows local testing without setting up OAuth

### Client-Side: CLI OAuth Flow

The CLI client (defined in [`utils/requests/request_with_auth.go`](../utils/requests/request_with_auth.go)) implements the OAuth2 authorization code flow with a local callback server.

**How it works:**

1. **Token Persistence**:
   - Tokens are stored in `~/.config/acmcsuf-cli/token.json` on Unix systems
   - The file contains: `access_token`, `refresh_token`, and `expiry` timestamp
   - Tokens are automatically loaded on subsequent CLI runs

2. **OAuth Flow** (when no valid token exists):
   ```
   ┌─────────┐                                  ┌─────────────┐
   │   CLI   │                                  │   Discord   │
   └────┬────┘                                  └──────┬──────┘
        │                                              │
        │ 1. Start local callback server               │
        │    on hardcoded port (:61234)                │
        │                                              │
        │ 2. Open browser with OAuth URL               │
        ├──────────────────────────────────────────>   │
        │   https://discord.com/oauth2/authorize       │
        │   ?client_id=...                             │
        │   &redirect_uri=http://localhost:61234       │
        │   &scope=identify                            │
        │   &response_type=code                        │
        │                                              │
        │          User approves in browser            │
        │                                              │
        │ 3. Discord redirects to callback             │
        │ <──────────────────────────────────────────  │
        │    http://localhost:61234/?code=...          │
        │                                              │
        │ 4. Exchange code for token                   │
        ├──────────────────────────────────────────>   │
        │    POST /oauth2/token                        │
        │                                              │
        │ 5. Receive access token                      │
        │ <──────────────────────────────────────────  │
        │    { access_token, refresh_token, ... }      │
        │                                              │
        │ 6. Save token to ~/.config/acmcsuf-cli/      │
        │                                              │
        │ 7. Make authenticated API request            │
        │    Authorization: Bearer <access_token>      │
        └──────────────────────────────────────────────┘
   ```

4. **Token Exchange**:
   - The callback server receives the authorization code from Discord
   - Exchanges the code for an access token via `POST https://discord.com/api/oauth2/token`
   - Stores the token with expiry information for future use

## Environment Variables

See [`developer-docs/env-vars.md`](./env-vars.md) for the complete list, but OAuth-specific variables include:

- `ENV`: Set to `production` to enable OAuth (default: `development`)
- `DISCORD_BOT_TOKEN`: Bot token for server-side API authentication
- `GUILD_ID`: Discord server/guild ID to verify membership
- `DISCORD_CLIENT_ID`: OAuth2 application client ID (for CLI)
- `DISCORD_CLIENT_SECRET`: OAuth2 application client secret (for CLI)

## Development Mode

During development (`ENV=development`), authentication is bypassed by supplying `dev-token` to the `Authorization header`. The CLI will do this automatically.

### Using the CLI

The CLI handles OAuth automatically:

```bash
# First run will trigger OAuth flow
./api.acmcsuf.com events get event-id

# Browser opens for authentication
# Token is saved to ~/.config/acmcsuf-cli/token.json
# Subsequent runs use the cached token
```

### Using curl/xh with OAuth

You need to manually pass the token when using a standard http client:
```bash
# Development mode (no real auth needed)
xh post :8080/v1/events --auth-type bearer --auth dev-token
# Or with short flags:
xh post :8080/v1/events -A bearer -a dev-token
# Or defining the header manually:
xh :8080/v1/events Authorization:'Bearer dev-token'
```

To get a real Discord access token for testing, run the OAuth flow with the CLI and extract the token from `~/.config/acmcsuf-cli/token.json`

### Token Expiry

Discord access tokens expire after a period (typically 1 week). When a token expires:

**CLI**: The OAuth flow automatically re-runs on the next command
**Manual testing**: You'll receive a `401 Unauthorized` response and need a new token

## Role Configuration

To add or modify roles, edit the `RoleMap` in [`internal/api/middleware/oauth.go`](../internal/api/middleware/oauth.go):

```go
var RoleMap = map[string]string{
    "1445971950584205554": "Board",     // Discord role ID -> Role name
    "another-role-id":     "Developer",
}
```

To find Discord role IDs:
1. Enable Developer Mode in Discord settings
2. Right-click a role in Server Settings → Roles
3. Click "Copy ID"
