## Card-Validator

### How to run
- Run in Docker:

```bash
make docker_pull
docker compose up
 ```

- Run test:
```bash
make test
 ```

<b>Server supports gRPC and HTTP Gateway.</b> 

gRPC can be tested on `http://0.0.0.0:9090` with `VerifyCard` rpc

Usage of http example:
````
curl --location 'http://0.0.0.0:8080/v1/card/verify' \
--data '{
    "card_number": 5555555555554444,
    "expiration_month": 3,
    "expiration_year": 2024
}'
````