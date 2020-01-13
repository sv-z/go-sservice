# User POST

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
    "email": "adm@example.com",
    "password": "A!a1"
}' http://localhost:8080/users
```

Response:
```bash
HTTP/1.1 201 Created
Date: Mon, 13 Jan 2020 01:40:53 GMT
Content-Length: 0
```

Validation error response:
```bash
HTTP/1.1 400 Bad Request
Date: Mon, 13 Jan 2020 01:36:35 GMT
Content-Length: 174
Content-Type: text/plain; charset=utf-8

{"error":{"email":[{"code":"a9fce4d8-7275-4727-956a-df4f21a10004","data":{"value":"admexample.com"},"message":"This value 'admexample.com' is not a valid email address."}]}}
```