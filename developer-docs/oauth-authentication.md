# OAuth Authentication

The ACM CSUF API uses Discord OAuth2 for authentication and authorization. This ensures that only authorized Discord server members with appropriate roles can access protected API endpoints.

## Overview

The API implements a role-based access control (RBAC) system that:
1. Authenticates users via Discord OAuth2
2. Verifies Discord server membership
3. Checks user roles against required permissions
4. Caches role information to avoid hitting Discord's rate limits

## Architecture

### Server-Side: API Middleware

The API uses Discord OAuth2 middleware defined in [`internal/api/middleware/oauth.go`](../internal/api/middleware/oauth.go) to protect all `/v1` routes.

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
        │    on random port (e.g., :61234)             │
        │                                              │
        │ 2. Open browser with OAuth URL               │
        ├──────────────────────────────────────────>   │
        │   https://discord.com/oauth2/authorize       │
        │   ?client_id=...                             │
        │   &redirect_uri=http://localhost:54321       │
        │   &scope=identify                            │
        │   &response_type=code                        │
        │                                              │
        │          User approves in browser            │
        │                                              │
        │ 3. Discord redirects to callback             │
        │ <──────────────────────────────────────────  │
        │    http://localhost:54321/?code=...          │
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

During development (`ENV=development`), authentication is bypassed:

**API Server:**
```go
// Any request with this header will bypass OAuth
Authorization: Bearer dev-token
```

**CLI Client:**
```go
// Automatically uses dev-token when ENV=development
// No OAuth flow occurs, no tokens are exchanged
```

To test the actual OAuth flow during development, temporarily set `ENV=production` in your `.env` file.

## Testing with OAuth

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

If testing manually with `curl` or `xh`, you need a valid Discord access token:

```bash
# Development mode (no real auth needed)
xh :8080/v1/events Authorization:"Bearer dev-token"

# Production mode (need real Discord token)
xh :8080/v1/events Authorization:"Bearer <discord_access_token>"
```

To get a real Discord access token for testing:
1. Run the CLI once to complete the OAuth flow
2. Extract the token from `~/.config/acmcsuf-cli/token.json`
3. Use that token in your curl/xh commands

### Token Expiry

Discord access tokens expire after a period (typically 1 week). When a token expires:

**CLI**: The OAuth flow automatically re-runs on the next command
**Manual testing**: You'll receive a `401 Unauthorized` response and need a new token

## Security Considerations

1. **Never commit tokens**: Token files (`~/.config/acmcsuf-cli/token.json`) contain sensitive credentials
2. **HTTPS in production**: The API should only run behind HTTPS in production to protect tokens in transit
3. **Client secrets**: Keep `DISCORD_CLIENT_SECRET` secure and never commit to git
4. **Rate limiting**: The middleware caches role information for 5 minutes to respect Discord's rate limits
5. **Token storage**: CLI tokens are stored with `0600` permissions (read/write for owner only)

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
