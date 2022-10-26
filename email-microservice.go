
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


    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type endpoint_information struct {
    name string 
    method string 
    description string 
    example string 
}

type incoming_email struct {
    Sender string `json:"sender"`
    Recipient string `json:"recipient"`
    Message string `json:"message"`
    Subject string `json:"subject"`
    SenderName string `json:"sendername"`
    RecipientName string `json:"recipientname"`
}

func show_endpoints(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category: %v\n", vars["category"])
    all := &endpoint_information{
        name: "All",
        method: "GET",
        description: "Returns all endpoints currently enabled",
        example: "/api/contact/email/all",
    }

    print_all, _ := json.Marshal(all)
    fmt.Println(print_all)

    fmt.Fprintf(w, "All routes: %s",print_all)
    
}

func send_email(w http.ResponseWriter, r *http.Request) {
    reqBody, err := ioutil.ReadAll(r.Body)
    if err == nil {
        var incoming incoming_email 
        err := json.Unmarshal(reqBody, &incoming)
        if err != nil{
            panic(err)
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
            w.WriteHeader(http.StatusNotFound  )
        } else {
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
    if(len(args) == 2 && (strings.ToLower(args[1]) == "run" || strings.ToLower(args[1]) == "help")){

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
            
        
    } else {
        panic("The acceptable commands are 'run' or 'help'")
    }
}