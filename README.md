# goa_microservice_sample
microservice architecture sample with goa

# usage

### in terminal1
```sh
$ cd user_service
$ go run *.go
```

### in terminal2
```sh
$ cd gateway
$ go run *.go
```

### in terminal3
```sh
$ curl -XGET localhost:8082/users |jq .
[
  {
    "created_at": "2018-04-24T02:30:37.89602+09:00",
    "id": 1,
    "name": "Andrew"
  },
  {
    "created_at": "2018-04-24T02:30:37.896022+09:00",
    "id": 2,
    "name": "Bob"
  },
  {
    "created_at": "2018-04-24T02:30:37.896023+09:00",
    "id": 3,
    "name": "Charlse"
  }
]

$ curl -XGET localhost:8081/gateway |jq .
{
  "users": [
    {
      "created_at": "2018-04-24T02:31:12.014015+09:00",
      "id": 1,
      "name": "Andrew"
    },
    {
      "created_at": "2018-04-24T02:31:12.014017+09:00",
      "id": 2,
      "name": "Bob"
    },
    {
      "created_at": "2018-04-24T02:31:12.014018+09:00",
      "id": 3,
      "name": "Charlse"
    }
  ]
}
```
