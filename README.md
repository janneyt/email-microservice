# email-microservice

# Overview
This is a gateway between a REST API and Twilio's SendGrid API. It is meant as a microservice, with some consumer sending data for an email and then receiving a reply back.

# Install
Make sure Go is installed. You can find directions here: https://go.dev/doc/install

Clone this github: 
`git clone htts://github.com/janneyt/email-microservice.git`

Change directories to the cloned repository:
`cd email-microservice`

In the email-microservice repository, run this command:
`go run email-microservice.go run`

Alternatively, you can build the executable:
`go build email-microservice.go`

And then run the microservice like so:
`./email-microservice run`

# Usage

You must use one of two arguments to run the service. These are the only valid command line arguments for this microservice:

`go run email-microservice.go run` and `go run email-microservice.go help`

##############################
# Communication Contract

Communication with this microservice relies on the following endpoints:

    - /
        -Receives: a GET request with no queries

        -Example call:
            ###
            GET http://localhost:8090 
            Content-Type: application/json 
            Authorization: x-access-token

        -Returns the following status codes:

            Status: 200
            Data: {
                endpoints : {*all endpoints will be here*}
            }

    - /api/contact/send/email
        -Receives: a POST request with the following mandatory fields:
            {
                "sender":"tedjanneyishuman@gmail.com",
                "recipient":"janneyt@oregonstate.edu",
                "message":"Hi Ted, Hope all is well."
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