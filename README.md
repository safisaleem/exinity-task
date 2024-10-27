

## **Architecture Overview**

The project is a microservice that processes transaction requests of type deposit and withdraw and sends requests to external payment gateways to process the transactions.

To adhere to SOLID principles, I have used the Controller-Service-Repository pattern which makes the code very modular and extendable.

Gateways are using the factory pattern to allow for quickly and easily adding support for new gateways.


-   **Gateway A**: Uses JSON over HTTP.
-   **Gateway B**: Uses SOAP/XML over HTTP.

**Overview of Components:**


![image](https://github.com/safisaleem/exinity-task/blob/main/diagram.png?raw=true)

-   **Controllers**: Handle HTTP requests and responses using the Gin framework.
-   **Services**: Contain business logic for transactions, balances, and webhooks.
-   **Repositories**: Interact with the PostgreSQL database using GORM. The system is built in a way that postgres can be easily swapped out with any database of choice.
-   **Gateways**: Interface with external payment gateways.
-   **Factories**: Provide instances of gateways.
-   **Models**: Define data structures for transactions.
-   **Types**: Contain request and response types for APIs and external integrations.
-   **Helpers**: Contain helper functions.
-  **Constants**: Contains constants to be used across the system.
-   **Retry Mechanism**: A background service that retries held transactions.

**Key Services:**

**1.⁠ ⁠Transaction Service:** This service creates and updates transactions using the transaction repository. 
**2.⁠ ⁠Balance Service:** This is used to keep track of the running total and is key in preventing the user from double spending (spending more than available running total).
**3.⁠ ⁠Webhook Service:** This service asynchronously  receives events from external gateways to update the status of the transactions from PENDING to FAILED or COMPLETE.
**4.⁠ ⁠Retry Service:** This service runs in a separate goroutine and periodically checks HELD transactions, which might be HELD due to unavailability of the external gateway, and then retries them so that they can be reprocessed.

**Transaction Lifecycle:**

**PENDING:** This is when a transaction has been successfully sent to the external payment gateway and is awaiting response via the webhook. The pending withdrawals are special because the running total includes these in balance calculation so that the user can not double spend (create a withdrawal before the webhook from another has arrived).
**HELD:** This is when the payment gateway is down or timing out.
**COMPLETE:** When the webhook confirms that the transaction is complete.
**FAILED:** This is when the webhook confirms that the external gateway couldnt process the transcation.

**Extensibility:**

The code allows for extensibility and new gateways can be added very easily to the system. Here's how:

1.⁠ ⁠Create a new gateway ⁠ `gatewayC.go` ⁠ in ⁠ `pkg/gateway/ `

2.⁠ ⁠Implement the SendDeposit and SendWithdraw functions of the payment gateway interface by adding logic to call external gateway

3.⁠ ⁠Add the new gateway to the gateway_factory.go

4.⁠ ⁠Add a new webhook route and controller to accept events from the new gateway

The system also allows the DB to be very easily swapped out by using GORM. To switch to a new db:

1.⁠ ⁠Create a new connection function in ⁠ `config/database.go` ⁠ to connect to the new db

2.⁠ ⁠Import the drivers for your db of choice with gorm and connect

3.⁠ ⁠Replace db in main


**Improvements and suggestions:**

**1.⁠ ⁠Configs:** Due to shortage of time, I couldn't add .env and use that to connect to db. If I had more time, I would have added a config file which first gets all the config variables, which could come from an env file, a yaml or even the command line and set a config object which would be returned to main. Main would then use this to connect to db, get the gateways ready.

**2.⁠ ⁠Tests:** I could not test the system too well, only a few basic test cases have been added. Expanding on those a very well tested code should be created.

**3.⁠ ⁠Logging:** The code doesn't log events very well. I would like to add a logging package like logrus to log events such as webhook arrival, transaction failed etc.

**4.⁠ ⁠Notifications:** The user should receive email/push notifications etc for transaction updates.

**5.⁠ ⁠Authentication:** The webhook endpoints should not be within the same router, or at least it should be authenticated based on the gateway's docs.

**6. Cloud Deployment:** The Go service has not been dockerized, which would allow for easier deployments on Cloud via Kubernetes.




## **Curl**

Create a Deposit:
```
curl --location 'localhost:8080/transactions' \ --header 'Content-Type: application/json' \ --data '{ "amount": 10.5, "user_id": "1", "type_handle": "DEPOSIT", "provider_handle": "gatewayA" }'
```
Create a Withdrawal:
```
curl --location 'localhost:8080/transactions' \ --header 'Content-Type: application/json' \ --data '{ "amount": 10.5, "user_id": "1", "type_handle": "WITHDRAW", "provider_handle": "gatewayA" }'
```

Simulate Webhook from GatewayA (Uses JSON over HTTP):

```
curl --location 'localhost:8080/webhook/gatewayA' \ --header 'Content-Type: application/json' \ --data '{ "id": "cefc2497-fcc5-4345-9252-708ff642f349", "updated_status": "success" }'
```
Simulate Webhook from GatewayB (Uses SOAP/XML over HTTP):
```
curl --location '[http://localhost:8080/webhook/gatewayB](http://localhost:8080/webhook/gatewayB)' \ --header 'Content-Type: text/xml' \ --data '<?xml version="1.0" encoding="UTF-8"?> <Envelope xmlns="[http://schemas.xmlsoap.org/soap/envelope/](http://schemas.xmlsoap.org/soap/envelope/)"> <Body> <TransactionRequest> <TransactionID>cb5a4fb9-c844-4a88-ad7d-9b0954a7b659</TransactionID> <Status>SUCCESSFULLY_COMPLETED</Status> </TransactionRequest> </Body> </Envelope>'
```
