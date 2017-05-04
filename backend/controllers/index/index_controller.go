package index

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"hexerent/backend/config"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
)

// Mail struct
type Mail struct {
	senderId string
	toIds    []string
	subject  string
	body     string
}

// SmtpServer struct
type SmtpServer struct {
	host string
	port string
}

// ServerName returns a host and port
func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

// BuildMessage builds the message to be sent
func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

// SendMail sends emails
func SendMail(r *http.Request) {

	mail := Mail{}
	mail.senderId = "jagreen1010@gmail.com"
	mail.toIds = []string{"jagreen1010@gmail.com", r.FormValue("emailReceiver")}
	mail.subject = "Hexerent Subscription"
	mail.body = "Hi \n\nThanks for subscribing to Hexerent."

	messageBody := mail.BuildMessage()

	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}

	log.Println(smtpServer.host)
	//build an auth
	auth := smtp.PlainAuth("", mail.senderId, "technology12", smtpServer.host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderId); err != nil {
		log.Panic(err)
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")

}

// Index renders the index.html file
func Index(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost || r.Method == "POST" {

		SendMail(r)

		config.Tpl.ExecuteTemplate(w, "index.html", nil)
	} else if r.Method == http.MethodGet || r.Method == "GET" {

		// Page visits

		pageVisits, err := r.Cookie("cookie-page-visit-counter")

		if err == http.ErrNoCookie {
			pageVisits = &http.Cookie{
				Name:  "cookie-page-visit-counter",
				Value: "0",
			}
		}

		count, err := strconv.Atoi(pageVisits.Value)

		if err != nil {
			log.Fatalln(err)
		}

		count++
		pageVisits.Value = strconv.Itoa(count)

		http.SetCookie(w, pageVisits)

		// Encoding strings

		s := "Love is but a song to sing Fear's the way we die You can make the mountains ring Or make the angels cry Though the bird is on the wing And you may not know why Come on people now Smile on your brother Everybody get together Try to love one another Right now"

		s64 := base64.StdEncoding.EncodeToString([]byte(s))

		fmt.Println(s64)

		bs, err := base64.StdEncoding.DecodeString(s64)
		if err != nil {
			log.Fatalln("I'm giving her all she's got Captain!", err)
		}
		fmt.Println(string(bs))

		config.Tpl.ExecuteTemplate(w, "index.html", nil)
	}

}

// About renders the about.html file
func About(w http.ResponseWriter, r *http.Request) {
	config.Tpl.ExecuteTemplate(w, "about.html", nil)
}
