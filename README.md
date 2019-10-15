# Go Mailjet API Relay

A simple API relay to send mail with mailjet.
The idea behind this project is to have a ready to work API hidding your mailjet key on which you can plug your front app without worrying about implementing a backend.

[![Build Status](https://travis-ci.com/adelplace/go-mailjet-api-relay.svg?branch=master)](https://travis-ci.com/adelplace/go-mailjet-api-relay)

## Getting Started

You can deploy this app everywhere you want using docker. 

### Prerequisites

This project use google recaptcha service to protect you from robots and spammers so you will need :
- A google recaptcha account
- A Mailjet account
- Docker installed on your server

### Installing

You will need some mandatory parameters to launch the app

| Parameter | Description |
| --------- | ----------- |
| RECAPTCHA_SECRET | Your recpatcha secret from https://www.google.com/recaptcha/admin |
| MAILJET_PRIVATE_KEY | Your mailjet private key from https://app.mailjet.com/account/api_keys |
| MAILJET_PUBLIC_KEY | Your mailjet public key from https://app.mailjet.com/account/api_keys |
| EMAIL | The email that will send the mail (must be configured in mailjet). Usually your own email |

#### Using docker

```
docker run --rm -d \
  -p 8080:80
  -e RECAPTCHA_SECRET="..." \
  -e MAILJET_PRIVATE_KEY="..." \
  -e MAILJET_PUBLIC_KEY="..." \
  -e EMAIL="..." \
  alexandredelplace/go-mailjet-api-relay
```

#### Using docker-compose

Create a docker-compose.yml file

```docker-compose
version: '3.7'
services:
  go-mailjet-api-relay:
    image: alexandredelplace/go-mailjet-api-relay
    ports:
      - "8080:80"
    environment:
      RECAPTCHA_SECRET: "..."
      MAILJET_PRIVATE_KEY: "..."
      MAILJET_PUBLIC_KEY: "..."
      EMAIL: "..."
```

Then run the container

```
docker-compose up -d
```

#### Using binaries

First you need to clone and build the project. Be sure to have a go version installed and up to date on your computer

```
git clone ...
go build ./...
```

Then fill the required env vars and run the application on the port of your wish

```
export RECAPTCHA_SECRET="..."
export MAILJET_PRIVATE_KEY="..."
export MAILJET_PUBLIC_KEY="..."
export EMAIL="..."

./application --addr=80
```

### Usage

You need to send few parameters to send your mail. 
The only one that is required is "g-recaptcha-response". If you don't send the other parameters, they will be replaced by an empty strings.

| Parameter | Description |
| --------- | ----------- |
| g-recaptcha-response | The token from the google recaptcha response |
| email | Email of the contact |
| name | Name of the contact |
| subject | Subject of the contact |
| message | Message of the contact |

So a curl request to send an email with the API would be something like

```
curl --request POST \
  --url http://localhost:8080 \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data 'g-recaptcha-response=my_captcha_response&name=my_name&subject=my_subject&message=my_message&email=my_email'
```

#### Return codes

The API return a json object that contain the status of your request

```json
{
  "success": false|true,
  "message": "Example message",
  "code": "example_error_code"
}
```

Here are the different codes that you can get from the API

| Status | Code    | Description |
| ------ | ------- | ----------- |
| 200    | success | The mail has been successfully delivred to Mailjet |
| 405    | method_not_allowed | Only POST requests are allowed |
| 400    | invalid_data | The sent data are invalid |
| 400    | no_captcha | The "g-recaptcha-response" parameter is missing |
| 400    | invalid_captcha | The captcha response is invalid/expired |
| 400    | mailjet_error | Mailjet returned an error |

## Running the tests

```
go test ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
