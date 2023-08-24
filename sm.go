package main

import (
    "bufio"
    "fmt"
    "os"
    "net/smtp"
)

func main() {
    // we are authorizing as
    from := "yourusername@danwin1210.de"

    // user we are sending email to
    if len(os.Args) != 2 {
        fmt.Println("Usage: sm recipient@domain.com < message_with_headers.txt")
        os.Exit(1)
    }
    to := os.Args[1]

    // server we are authorized to send email through
    host := "danwin1210.de"

    auth := smtp.PlainAuth("", from, "yourpassword", host)

    scanner := bufio.NewScanner(os.Stdin)
    message := ""
    for scanner.Scan() {
        message += scanner.Text() + "\n"
    }

    if err := smtp.SendMail(host+":587", auth, from, []string{to}, []byte(message)); err != nil {
        fmt.Println("Error SendMail: ", err)
        os.Exit(1)
    }
    fmt.Println("Email Sent!")
}
