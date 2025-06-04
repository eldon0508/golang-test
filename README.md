https://go.dev/doc/tutorial/

# Player Management System

- Content-Type: `application/json`

## Create player

- Response: **Success** (201 Created)

### Request Body

```json
{
  "name": "Ben",
  "level_id": "1"
}
```

## Create level

- Response: **Success** (201 Created)

### Request Body

```json
{ "name": "Expert" }
```

# Game Room Management System

- Content-Type: `application/json`

## Create Room

### Request Body

```json
{
  "name": "Quad",
  "description": "Quad Room"
}
```

## Check Reservation

- `room_id`: optional
- `date`: optional
- `limit`: optional

```http
GET /reservations?room_id=1&date=2025-06-01&limit=1
```

## Create Reservation

- Response: **Success** (201 Created)

### Request Body:

```json
{
  "room_id": "2",
  "date": "2025-06-05T00:00:00Z",
  "player_id": "2"
}
```

# Endless Challenge System

## Create Challenge

- Response: **Success** (201 Created)

### Request Body:

```json
{
  "player_id": "2",
  "fee": 20.01
}
```

# Game Log Collector

## Check Log

- `player_id`: optional
- `action`: optional
- `start_time`: optional
- `end_time`: optional

```http
GET /logs?player_id=1&action=login&start_time=09:00&end_time=12:00
```

## Create Log

- Response: **Success** (201 Created)

### Request Body:

```json
{
  "player_id": "2",
  "action": "check-result",
  "start_time": "2025-06-05T00:00:00Z",
  "end_time": "2025-06-05T00:01:00Z"
}
```

# Payment Processing System

## Create Payment

- Response: **Success** (201 Created)

### Request Body:

```json
{
  "method": "Card",
  "amount": 33.33,
  "description": "Card transaction",
  "status": "Failed"
}
```

## Check Payment

```http
GET /payments/1
```
