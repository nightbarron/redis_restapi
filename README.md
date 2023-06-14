# Redis REST API

## Run
```bash
docker run -d -p 8080:8080 redis-rest:latest
```

## Structure
```go
type RedisDataModel struct {
Key    string `json:"key"`
Value  string    `json:"value"`
Expire int    `json:"expire" default:"0"`
}
```
## Details
| Key  | Types  | Description    | Value        | Default        |
|------|--------|----------------|--------------|----------------|
| key  | string | Key            | string       | *              |
| Value | string | Value          | string       | *              |
| Expire | int | Time for Expired | int (second) | 0: non expired |

_* : required_
## Get Version

| API Endpoint            |
| ----------------------- |
| GET /v1/version |

##### Request

`curl -L -X GET 'localhost:8080/v1/version'`

##### Respond

```json
{
  "success": true,
  "message": "Version 1.0.0"
}
```

## GET KEY

| API Endpoint      |
|-------------------|
| GET /v1/keys/:key |

##### Request

`curl -L -X GET 'localhost:8080/v1/keys/:keyexample'`

##### Respond

```json
{
  "success": true,
  "message": "Get key from Redis server",
  "data": {
    "key": "keyexample",
    "value": "1",
    "expire": 0
  }
}
```

## PUT KEY

| API Endpoint |
|--------------|
| PUT /v1/keys |

##### Request

```jsunicoderegexp
curl -L -X PUT 'localhost:8080/v1/keys' 
--header 'Content-Type: application/json' 
--data '{
"key": "keyexample",
"value": 1,
"expire": 10
}'
```

##### Respond

```json
{
  "success": true,
  "message": "Put key to Redis server",
  "data": {
    "key": "keyexample",
    "value": 1,
    "expire": 10
  }
}
```

## DELETE KEY

| API Endpoint               |
|----------------------------|
| DELETE /v1/keys/keyexample |

##### Request

```
`curl -L -X DELETE 'localhost:8080/v1/keys/:keyexample'`
```

##### Respond

```json
{
  "success": true,
  "message": "Delete key from Redis server"
}
```