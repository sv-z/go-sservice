# Users GET

### Parameters 
- **id**(required) - Correct user's id
Request:
```bash
curl -i -H "Content-Type: application/json; indent=4" http://localhost:8080/users/{id}
```

Response:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 17 Jan 2020 00:09:03 GMT
Content-Length: 35

{"id":5,"email":"adm5@example.com"}
```