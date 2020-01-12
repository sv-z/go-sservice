# Find user by params

- **namespace**:
    - user

### Parameters 

Request:
```bash
curl -i -X POST -H "Content-Type: application/json; indent=4" -d '{
    "jsonrpc": "2.0",
    "method": "user.Find",
    "params": [],
    "id": "50836230-ea1b-11e7-af6b-0242ac12000f"
}' http://localhost:8080/api/
```

Response:
```json
{
    "jsonrpc": "2.0",
    "result": [
        {
            "email": "test1@example.com",
            "id": 1
        },
        {
            "email": "test2@example.com",
            "id": 10
        }
    ],
    "id": "50836230-ea1b-11e7-af6b-0242ac12000f"
}
```