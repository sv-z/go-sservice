# User register

- **namespace**:
    - user

### Parameters 
- **email**(required) - Correct user's email address
- **password**(required) - User's password
    - _minimum length 4 chars_  
    - _maximum length 20 chars_  
    - _At least one upper case English letter, (?=.*?[A-Z])_
    - _At least one lower case English letter, (?=.*?[a-z])_
    - _At least one digit, (?=.*?[0-9])_
    - _At least one special character, (?=.*?[#?!@$%^&*-])_

Request:
```bash
curl -i -X POST -H "Content-Type: application/json; indent=4" -d '{
    "jsonrpc": "2.0",
    "method": "user.Register",
    "params": {
        "email": "adm@example.com",
        "password": "A!a1"
    },
    "id": "50836230-ea1b-11e7-af6b-0242ac12000f"
}' http://localhost:8080/api/
```

Response:
```json
{
    "jsonrpc": "2.0",
    "result": true,
    "id": "50836230-ea1b-11e7-af6b-0242ac12000f"
}
```

Validation error response:
```json
{
    "jsonrpc": "2.0",
    "error": {
        "code": -32602,
        "message": "Validation error.",
        "data": {
            "email": [
                {
                    "code": "a9fce4d8-7275-4727-956a-df4f21a10004",
                    "data": {
                        "value": "admexample.com"
                    },
                    "message": "This value 'admexample.com' is not a valid email address."
                }
            ]
        }
    },
    "id": "50836230-ea1b-11e7-af6b-0242ac12000f"
}
```