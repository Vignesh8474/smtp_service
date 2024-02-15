package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	// Load dev.env file
	err := godotenv.Load("dev.env")

	if err != nil {
		log.Fatal(err)
	}

	//Get the value from the env file
	hostName := os.Getenv("hostName")
	from := os.Getenv("from")
	pass_key := os.Getenv("pass_Key")

	//Port return type is string
	port := os.Getenv("port")

	//Convert string into int for port
	smtPort, err := strconv.Atoi(port)

	if err != nil {
		log.Fatal(err)
	}
	//Calling the function
	sendGoMail(hostName, smtPort, from, pass_key)
}

func sendGoMail(hostName string, smtPort int, from, pass_Key string) {

	// Gomail using create a new message
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", "raghulraja@flipopay.com", "rajaram@flipopay.com", "ali.afroze@flipopay.com", "dharani@flipopay.com", "jeyashree.s@flipopay.com")
	m.SetAddressHeader("Cc", "1916081@saec.ac.in", "Vignesh")
	m.SetHeader("Subject", "Welcome Greet")

	//content of the body mail
	body :=
		`
	<html>
		 <style>
		* {
			margin: 0;
			padding: 0;
			box-sziing: border-box;
		}
		</style>

	  <body>
          <h4>Hi,</h4>
		  <h2>Thanks for joining Flipopay.</h2> </br>
		 <p style={{ font-size:20px"}}> We developed our payment gateway service to facilitate seamless sending and receiving of payments, ensuring a hassle-free experience for our users. 
		 
		  Thank you for choosing our payment gateway service. We're committed to providing you with top-notch service.</p> </br>
		  <img src="https://i.ibb.co/HTmh9ck/flipopay-logo.jpg" alt="flipopay-logo" width="15%" > </br>
		  <h3>[Flipopay]</h3>
	  </body>
	</html>
	`

	m.SetBody("text/html", body)

	// Create a NewDailer
	d := gomail.NewDialer(hostName, smtPort, from, pass_Key)

	// Print the value of d
	fmt.Printf("d: %v\n", d)

	// Connect the dialand send the message
	err := d.DialAndSend(m)

	if err != nil {
		// Stop the execution
		panic(err)
	}

}
