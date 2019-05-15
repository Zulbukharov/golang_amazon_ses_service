![](416aeba28c827174014c487ae9020cec.jpg)
# Confirm service

A contact info confirmation service.

## API

> In this example, Our service will allow to confirm both emails and phones
> route `/email`.

###  QueryString parameters
  - `contact_info`
  - `verification_code`

#### Example
  - [go](https://golang.org/pkg/net/url/#example_URL_Query)
  - [js](https://nodejs.org/api/url.html#url_class_urlsearchparams)


### 1. Requesting the confirmation code to be sent
If just `contact_info` is sent we ask for you to send the confirmation code to the contact info.

#### Expected HTTP Status
  - __`200`__ `Ok`

#### Example:
# for email:
curl localhost:3000/email?contact_info=le.mikmac%40gmail.com
# > HTTP/1.1 200 OK
```


### 2. Sending the confirmation code
If the `verification_code` and `contact_info` are sent, we ask for the comfirmation

#### Expected HTTP Status
  - __`200`__ `Ok`
    > the `contact_info` is confirmed
  - __`401`__ `Unauthorized`
    > the `verification_code` is invalid


#### Example:
```console
# for email:
curl localhost:3000/email?contact_info=le.mikmac%40gmail.com&verification_code=3fJ4Bz
# > HTTP/1.1 200 OK
```

### Additionnal information

```
# export AWS_ACCESS_KEY_ID=""
# export AWS_SECRET_ACCESS_KEY=""
# export SENDER_REGION=""
# export SENDER_EMAIL=""
```
