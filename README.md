# email-microservice

# Overview
This is a gateway between a REST API and Twilio's SendGrid API. It is meant as a microservice, with some consumer sending data for an email and then receiving a reply back.

# Install
Clone the github: 
`git clone -u www.github.com/janneyt/email-microservice`

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