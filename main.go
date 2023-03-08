package main

import "fmt" //stands for format package

func main() {
	//bu := anlamı değişken olur ve ona bir değer ata.
	conferenceName := "GO Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("conference ticket is %T, remainigticket is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//array tanımlıyoruz.
	//empty array tanımı
	//var bookings = [50]string{}
	//bookingsEmpty[0] = "hasan"
	//bookingsEmpty[1] = "ali"
	//dolu array tanımlama
	//var bookings = [50]string{"İsmail", "Hüseyin"}
	var bookings = []string{}

	var userName string
	var lastName string
	var email string
	var userTickets int

	//clidan username almak için kullanacağız...
	//bu şekilde çalıştırırsak bomboş geçer ve uygulama çalışır ama hiçbir değişken atanmaz.
	//bu boş geçmeyi önlemek için & pointer değişkenin önüne eklenir.
	fmt.Println("Please enter your username")
	fmt.Scan(&userName)
	fmt.Println("Please enter your lastname")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email")
	fmt.Scan(&email)
	fmt.Println("How many ticket you want to buy?")
	fmt.Scan(&userTickets)
	//userName = "oğan"
	//userTickets = 4

	//fmt.Println(remainingTickets)
	//fmt.Println(&remainingTickets)

	fmt.Printf("User %v, booked %v tickets \n", userName, userTickets)
	availableTicketsCount := (int(remainingTickets) - userTickets)

	fmt.Printf("So now we have %v tickets available \n", availableTicketsCount)
	//bookings[0] = userName + " " + lastName
	bookings = append(bookings, userName+" "+lastName)

	fmt.Printf("The Whole Slice: %v \n", bookings)
	fmt.Printf("The First item of Slice: %v \n", bookings[0])
	fmt.Printf("Type of: %T \n", bookings)
	fmt.Printf("Slice length: %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", userName, lastName, userTickets, email)

}

/*
Creates a new  module
Module path cah correspond to a repository  you plan to publish your module to(e.g github.com/username/reponame)
>> go mod init booking-app
bu komutu çalıştırınca
go.mod dosyası yaratılıyor
modülü tarif ediyor : with name/module path
and go version used in program

the module path is also the import path
(e.g import github.com/username/reponame)
*/

/*
Go programs are organized into packages
Go's standart library, provides diffrent core packages for us to use
"fmt" is one of these , which you can use by importing it.
a package is a collection of source file
diğer paketlere bakmak için
pkg.go.dev /fmmt

*/

/*
Print formatted data
fmt.printf("some text %s", myvariable)

//gopath'de sorun vardı
//https://github.com/golang/vscode-go/issues/166
in settings.json

{
"go.goroot": "/snap/go/current"
}
https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04#step-1-installing-go
https://stackoverflow.com/questions/15407719/in-gos-http-package-how-do-i-get-the-query-string-on-a-post-request

Part-1
https://www.youtube.com/watch?v=yyUHQIec83I
42:33

Part-2


https://www.youtube.com/watch?v=bLirVi8gw40
https://www.youtube.com/watch?v=8sU7y9rRVYY
https://www.youtube.com/watch?v=evJltOj78CM&list=PLfCVAl3pRAFgmuToYPlkSszL3ygoVQDmE
https://www.youtube.com/watch?v=Ap8clduWwVc
https://www.youtube.com/watch?v=0xBPHNAFbWk&list=PLr48dQTh3FFyWe4twhEmeoO-Yvt6rsDPw

https://www.youtube.com/watch?v=YS4e4q9oBaU
https://www.youtube.com/watch?v=qJqx7dxWW8I&list=PL-Hkw4CrSVq96dPr33xTdBjSgn9wKLHPa

*/
