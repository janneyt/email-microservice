
package main

import (
    "fmt"
    "os"
    "github.com/gorilla/mux"
    "net/http"
    "log"
    "encoding/json"
    "io/ioutil"
    "strings"
    "github.com/joho/godotenv"

    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type required_post_properties struct {
    Sender string                               `json:"Sender"`
    SenderName string                           `json:"SenderName"`
    Recipient  string                           `json:"Recipient"`
    RecipientName string                        `json:"RecipientName"`
    Subject      string                         `json:"Subject"`
}

type endpoint_information struct {
    Name string                                 `'json:"Name"`
    Method string                               `json:"Method"`
    Description string                          `json:"Description"`
    Example string                              `json:"Example"`
    Required required_post_properties           `json:"Required"`
}

type return_endpoints struct {
    Endpoints []endpoint_information            `json:"Endpoints"`
}

type incoming_email struct {
    Sender string                               `json:"sender"`
    Recipient string                            `json:"recipient"`
    Message string                              `json:"message"`
    Subject string                              `json:"subject"`
    SenderName string                           `json:"sendername"`
    RecipientName string                        `json:"recipientname"`
}

func show_endpoints(w http.ResponseWriter, r *http.Request) {
    _ , err := ioutil.ReadAll(r.Body)
    if err == nil {
        
        /* Register all routes here */

        required_properties := required_post_properties{
            Sender:"Not required",
            SenderName:"Not required",
            Recipient:"Not required",
            RecipientName:"Not required",
            Subject:"Not required",
        }
        all := endpoint_information{
            Name: "All",
            Method: "GET",
            Description: "Returns all endpoints currently enabled",
            Example: "/api/contact/email/all",
            Required: required_properties,
        }

        returnables := []endpoint_information{
            all,
        }

        /* Setup send_email with its required post properties data member */
        required_properties1 := required_post_properties{
            Sender:"Not required",
            SenderName:"Not required",
            Recipient:"Not required",
            RecipientName:"Not required",
            Subject:"Not required",
        }
        send_email := endpoint_information{
            Name:"send_email",
            Method:"POST",
            Description: "Sends a email based on the passed properties.",
            Example: "/api/contact/email/send_email",
            Required:  required_properties1,
        }
        returnables = append(returnables, send_email)
        returnable, err := json.Marshal(returnables)
        if err != nil {
            /* Mangled data returns a 404*/
            w.WriteHeader(http.StatusNotFound)
        } else {
            w.WriteHeader(http.StatusOK)
            fmt.Fprintf(w, "%+v", returnable)
        }
    } else {
        w.WriteHeader(http.StatusBadRequest)
    }  
}
 
func send_email(w http.ResponseWriter, r *http.Request) {
    err := godotenv.Load("sendgrid.env")
    reqBody, err := ioutil.ReadAll(r.Body)
    if err == nil {
        var incoming incoming_email 
        err := json.Unmarshal(reqBody, &incoming)
        if err != nil{
            w.WriteHeader(http.StatusNotFound)
        }
        from := mail.NewEmail(incoming.SenderName, incoming.Sender)
        to := mail.NewEmail(incoming.RecipientName, incoming.Recipient)
        subject := incoming.Subject
        plainTextContent := incoming.Message
        htmlContent := incoming.Message
        message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
        client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
        response, err := client.Send(message)
        if err != nil {
            w.WriteHeader(http.StatusNotFound)
        } else {
            if response.StatusCode == 401 {
                w.WriteHeader(http.StatusUnauthorized)
                fmt.Fprintf(w,"%v","The microservice is not currently authorized to communicate with our email partner, Twilio. Please check to make sure the API Key is valid and all authorization is setup per: https://github.com/sendgrid/sendgrid-go")
            }
            fmt.Fprintf(w,"%+v",response)
        }
    } else {
        w.WriteHeader(http.StatusBadRequest)
    }
}

func route_handler() {
    router := mux.NewRouter()
    router.HandleFunc("/",show_endpoints).Methods("GET")
    router.HandleFunc("/api/contact/email/send_email", send_email).Methods("POST")
    log.Fatal(http.ListenAndServe(":8090", router))
}

func main() {

    args := os.Args
    if(len(args) == 2 && (strings.ToLower(args[1]) == "run" )){

        fmt.Println("Beginning email microservice by Ted Janney");

        fmt.Println("Waiting for email input via REST API")

        fmt.Println("Press x to end the service at any time")




        go route_handler()
        var input string
        fmt.Scan(&input)
        if input == "x" {
            fmt.Println("Leaving")
            os.Exit(0)
        }
            
        
    }   else if len(args) == 2 && strings.ToLower(args[1]) == "help" {
        fmt.Println("Please consult the Readme while I create a help menu.")

    } else {
        panic("The acceptable commands are 'run' or 'help'")
    }
}