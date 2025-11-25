# API Test Fixtures

This directory contains JSON payloads for testing the ACM CSUF API endpoints.

## Organization

Fixtures are named using the pattern: `{resource}_{action}_{variant}.json`

- `resource`: The API resource (event, announcement, officer, tier, position)
- `action`: The operation (create, update, delete)
- `variant`: Optional descriptor for the specific test case

## Events

### Create Operations
- `event_create_basic.json` - Basic event with required fields
- `event_create.json` - Workshop event with specific location
- `event_create_all_day.json` - All-day hackathon event
- `event_create_multiple_hosts.json` - Event with multiple host organizations

### Update Operations
- `event_update_partial.json` - Partial update (location and all_day status)
- `event_update_location.json` - Update only the location field
- `event_update_full.json` - Full update with all fields

## Announcements

### Create Operations
- `announcement_create_basic.json` - Basic announcement with Discord integration
- `announcement_create.json` - Public announcement with full Discord details
- `announcement_create_private.json` - Private announcement for officers
- `announcement_create_no_discord.json` - Website-only announcement without Discord

### Update Operations
- `announcement_update.json` - Update announcement visibility and Discord info

## Board Members

### Officers

#### Create Operations
- `officer_create.json` - Officer with complete profile
- `officer_create_minimal.json` - Officer with only required name field
- `officer_create_github_only.json` - Officer with GitHub handle only

#### Update Operations
- `officer_update.json` - Update officer's full profile

### Tiers

#### Create Operations
- `tier_create.json` - Executive tier (President)
- `tier_create_dev_team.json` - Developer tier

#### Update Operations
- `tier_update.json` - Update tier title and team

### Positions

#### Create Operations
- `position_create.json` - Create position for officer in a semester

#### Update Operations
- `position_update.json` - Update position title and team

#### Delete Operations
- `position_delete.json` - Delete position identifier

## Field Types

### sql.NullString
Fields that are nullable strings use this format:
```json
{
  "string": "value",
  "valid": true
}
```
For null values:
```json
{
  "string": "",
  "valid": false
}
```

### sql.NullInt64
Fields that are nullable integers use this format:
```json
{
  "int64": 123,
  "valid": true
}
```

### sql.NullBool
Fields that are nullable booleans use this format:
```json
{
  "bool": true,
  "valid": true
}
```

## Timestamps

All timestamps are in Unix milliseconds (e.g., `1712851200000` = April 11, 2024).

## Usage Examples

### Using curl

```bash
# Create an event
curl -X POST http://localhost:8080/v1/events \
  -H "Content-Type: application/json" \
  -d @fixtures/event_create.json

# Update an event
curl -X PUT http://localhost:8080/v1/events/workshop-intro-to-git \
  -H "Content-Type: application/json" \
  -d @fixtures/event_update_location.json

# Create an officer
curl -X POST http://localhost:8080/v1/board/officers \
  -H "Content-Type: application/json" \
  -d @fixtures/officer_create.json

# Delete a position
curl -X DELETE http://localhost:8080/v1/board/positions \
  -H "Content-Type: application/json" \
  -d @fixtures/position_delete.json
```

## Notes

- UUIDs should be unique across all resources
- Events require both `location` and `host` fields
- Officers require at minimum the `full_name` field
- Positions are identified by the composite key: `(oid, semester, tier)`
- Update operations use COALESCE, so only provide fields you want to change
