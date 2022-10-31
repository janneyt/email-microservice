# email-microservice

# Overview
This is a gateway between a REST API and Twilio's SendGrid API. It is meant as a microservice, with some consumer sending data for an email and then receiving a reply back.

# Install
Make sure Go is installed. You can find directions here: https://go.dev/doc/install

Clone this github: 
`git clone htts://github.com/janneyt/email-microservice.git`

Change directories to the cloned repository:
`cd email-microservice`

Initialize the mod system
`go mod init email-microservice.com/m/v2`

Tidy the mod (install dependencies, among many other useful features)
`go mod tidy`

Now you can build the executable:
`go build email-microservice.go`

And then run the microservice like so:
`./email-microservice run`
or on Windows:
`email-microservice run`

# Usage

You must use one of two arguments to run the service. These are the only valid command line arguments for this microservice:

`go run email-microservice.go run` and `go run email-microservice.go help`

##############################
# Communication Contract

Communication with this microservice relies on the following endpoints:

    - Endpoint: /
        -Receives: a GET request with no queries

        -Example call:
            ###
            GET http://localhost:8090 
            Content-Type: application/json 
            Authorization: x-access-token

        -Returns the following status codes:

            Status: 200
            Data: 
                [{
                    "Name":"All",
                    "Method":"GET",
                    "Description":"Returns all endpoints currently enabled",
                    "Example":"/api/contact/email/all",
                    "Required":
                        {"Sender":"Not required",
                        "SenderName":"Not required",
                        "Recipient":"Not required",
                        "RecipientName":"Not required",
                        "Subject":"Not required"}},
                {
                    "Name":"send_email",
                    "Method":"POST",
                    "Description":"Sends a email based on the passed properties.","Example":"/api/contact/email/send_email",
                    "Required":
                        {"Sender":"Not required",
                        "SenderName":"Not required",
                        "Recipient":"Not required",
                        "RecipientName":"Not required",
                        "Subject":"Not required"}}]

            Status: 400-Bad Request
            Data: None, this indicates a malformed request on the calling agent's side

            Status: 404-Not Found
            Data: None, this indicates a processing error but a graceful exit and reply to the end user.

    - Endpoint: /api/contact/send/email
        -Receives: a POST request with the following mandatory fields:
            {
                "sender":"tedjanneyishuman@gmail.com",
                "recipient":"janneyt@oregonstate.edu",
                "message":"Hi Ted, Hope all is well."
                "sendername":"Ted"
                "recipientname":"Ted"
            }
        
        -Example call:
            ###
            POST http://localhost:8090/api/contact/email/send_email HTTP/1.1
            Content-Type: application/json
            Authorization: x-access-token

            {
                "sender":"tedjanneyishuman@gmail.com",
                "recipient":"janneyt@oregonstate.edu",
                "message":"Hi Ted, Hope all is well."
            }
            
            ![Email Contact Form Microservice vpd](https://user-images.githubusercontent.com/70920801/199119972-65b86ab7-8d11-448e-aeec-e3ca9d03e2cb.jpg)
file:///home/castimir/Downloads/Email%20Contact%20Form%20Microservice.vpd.jpg![image](https://user-images.githubusercontent.com/70920801/199120089-a10fae54-fcd7-4afc-a488-ea8386b0e3ff.png)
