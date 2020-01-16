# Users GET

### Parameters 

Request:
```bash
curl -i -H "Content-Type: application/json; indent=4" http://localhost:8080/users
```

Response:
```bash
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 16 Jan 2020 23:25:59 GMT
Content-Length: 50

[
  {
    "id": 1,
    "email": "adm@example.com"
  },
  {
    "id": 15,
    "email": "adm1@example.com"
  }
}
```