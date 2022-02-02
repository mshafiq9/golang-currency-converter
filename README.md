# golang-currency-converter
GoLang API converting given amount from one currency to another using live rates saving each transaction

## Installation
Download and install GoLang if you don't have [https://go.dev/doc/install](https://go.dev/doc/install)

## Prerequisite
This project need following two environment variables:

1. `DATABASE_URL` which should follow the patteren `DATABASE_URL=postgres://{user}:{password}@{hostname}:{port}/{database-name}`
2. TraderMade API key which you can after [sign up](https://marketdata.tradermade.com/signup) on TraderMade for free, copy the API key from your [dashboard](https://marketdata.tradermade.com/myAccount).

```sh
# export DATABASE_URL env var in your shell
# DATABASE_URL=postgres://{user}:{password}@{hostname}:{port}/{database-name}

export DATABASE_URL="<your_posgresql_database_url>"

# export the API key env var in your shell
export TRADERMADE_API_KEY="<your_own_api_key>"
```

## Checkout and Run
```sh
$ git clone git@github.com:mshafiq9/golang-currency-converter.git
$ cd golang-currency-converter
$ go run .
```

## Testing
Note: During testing also cross verify console output when hitting API endpoints.

```sh
# Convert Currencies (Read the Output):

curl -i -X POST -H 'Content-Type: application/json' \
-d '{"date": "2022-02-03T01:56:31+04:00", "from": "EUR", "to": "USD", "amount": 10.01}' \
http://localhost:10000/convert

curl -i -X POST -H 'Content-Type: application/json' \
-d '{"date": "2022-02-03T01:56:31+04:00", "from": "EUR", "to": "USD", "amount": 34534.05}' \
http://localhost:10000/convert

curl -i -X POST -H 'Content-Type: application/json' \
-d '{"date": "2022-02-03T01:56:31+04:00", "from": "GBP", "to": "USD", "amount": 29089.09}' \
http://localhost:10000/convert

curl -i -X POST -H 'Content-Type: application/json' \
-d '{"date": "2022-02-03T01:56:31+04:00", "from": "CAD", "to": "USD", "amount": 2342.03}' \
http://localhost:10000/convert
```

```
# Read All Transactions:

http://localhost:10000/transactions

curl -i -X GET http://localhost:10000/transactions
```

## Output
Read the console output. It should be similiar to as below:

#### Converting Curriencies
```sh
curl -i -X POST -H 'Content-Type: application/json' \
-d '{"date": "2022-02-03T01:56:31+04:00", "from": "CAD", "to": "USD", "amount": 2342.01}' \
http://localhost:10000/convert

HTTP/1.1 200 OK
Date: Wed, 02 Feb 2022 22:42:43 GMT
Content-Length: 140
Content-Type: text/plain; charset=utf-8

{"id":7,"date":"2022-02-03T01:56:31+04:00","from":"CAD","to":"USD","amount":2342.01,"rate":0.7890293,"converted_amount":1847.9145108930002}
```

#### Read All Transactions::
```sh
curl -i -X GET http://localhost:10000/transactions

HTTP/1.1 200 OK
Date: Wed, 02 Feb 2022 23:24:02 GMT
Content-Length: 916
Content-Type: text/plain; charset=utf-8

[{"id":1,"date":"2022-02-02T21:47:31Z","from":"EUR","to":"USD","amount":10.01,"rate":1.13048,"converted_amount":11.3161048},{"id":2,"date":"2022-02-02T21:56:32Z","from":"GBP","to":"USD","amount":10.01,"rate":1.3574,"converted_amount":13.587573999999998},{"id":3,"date":"2022-02-02T21:56:31Z","from":"GBP","to":"USD","amount":29089.01,"rate":1.35779,"converted_amount":39496.7668879},{"id":4,"date":"2022-02-02T21:56:31Z","from":"CAD","to":"USD","amount":2342.01,"rate":0.7891227,"converted_amount":1848.133254627},{"id":5,"date":"2022-02-02T21:56:31Z","from":"CAD","to":"USD","amount":2342.01,"rate":0.7891227,"converted_amount":1848.133254627},{"id":6,"date":"2022-02-02T21:56:31Z","from":"CAD","to":"USD","amount":2342.01,"rate":0.7890293,"converted_amount":1847.9145108930002},{"id":7,"date":"2022-02-02T21:56:31Z","from":"CAD","to":"USD","amount":2342.01,"rate":0.7890293,"converted_amount":1847.9145108930002}]
```


## References
https://github.com/mshafiq9/golang-rest-client

https://github.com/mshafiq9/golang-restful-api

https://github.com/mshafiq9/golang-hasura-postgresql
