
# Project: Tasks management system
A REST-API for tasks management system written in Golang.
# Require:
```bash
go >= 1.16
git 
PostgresSQL installed
```

# How to install:
```bash
git clone https://github.com/leehai1107/Task-managent-sys
```
# How to run:
## Excute:
```bash
script.sql
```
## Run
```bash
go run main.go
```

# API Reference

#### Register user

```http
  GET /user/register
```
```bash
{
    "username": "test",
    "password": "mypassword",
    "dob": "2006-07-11",
    "sex": "Male",
    "address": "VN",
    "phone": "0123456789"
}
```

| Parameter  | Type     | Description                |
| :--------- | :------- | :------------------------- |
| `username` | `string` | **Required**. Your username|
| `password` | `string` | **Required**. Your password|
|   `dob`    | `string` | Your dob yyyy-mm-dd        |
|   `sex`    | `string` | Your gender                |
|  `address` | `string` | Your address               |
|  `phone`   | `string` | Your phone                 |

Takes in data and return user_id base on success.
```bash
{
    "status": "Success",
    "data": {
        "user_id": 1
    },
    "error": null
}
```




