# stripe-apis-test
sample project to test a few stripe APIs

## start up instructions
start the server by entering the following command in the terminal

```golang
    go run main.go
```
### Endpoints

    - /api/v1/create_charge
    - /api/v1/capture_charge/{chargeId}
    - /api/v1/create_refund/{chargeId}
    - /api/v1/get_charges

    {chargeId} is to replaced with the an actual chargeId ex. ch_3KJ9Ti2Ri1hXKaRW0yZo8ufy