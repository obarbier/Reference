## Definition
Hypermedia as the Engine of Application State (HATEOAS) is a constraint of the REST application architecture that distinguishes it from other network application architectures.

With HATEOAS, a client interacts with a network application whose application servers provide information dynamically through hypermedia. A REST client needs little to no prior knowledge about how to interact with an application or server beyond a generic understanding of hypermedia.

By contrast, clients and servers in Common Object Request Broker Architecture (CORBA) interact through a fixed interface shared through documentation or an interface description language (IDL).

The restrictions imposed by HATEOAS decouple client and server. This enables server functionality to evolve independently.

```
GET /accounts/12345 HTTP/1.1
Host: bank.example.com
```

```
HTTP/1.1 200 OK

{
    "account": {
        "account_number": 12345,
        "balance": {
            "currency": "usd",
            "value": 100.00
        },
        "links": {
            "deposits": "/accounts/12345/deposits",
            "withdrawals": "/accounts/12345/withdrawals",
            "transfers": "/accounts/12345/transfers",
            "close-requests": "/accounts/12345/close-requests"
        }
    }
}
```

## Glossary
* **HATEOAS**: Hypermedia as the Engine of Application State

## Reference
* [Wiki Page](https://en.wikipedia.org/wiki/HATEOAS)
* [Richardson Maturity Model](https://restfulapi.net/richardson-maturity-model/)