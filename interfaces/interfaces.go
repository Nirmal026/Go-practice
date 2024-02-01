package main

import "fmt"

type bank interface {
	deposite(int)
	withdrawal(int)
	mainBalance()
}

type hdfc struct {
	balance int
}

func (a *hdfc) deposite(num int) {
	a.balance = a.balance + num
	fmt.Printf("Your balance amount after deposite : %d\n", a.balance)
}

func (a *hdfc) withdrawal(num int) {
	a.balance = a.balance - num
	if a.balance < 500 {
		fmt.Println("You have only minimum balance")
	} else {
		fmt.Printf("Your balance amount after withdrawal : %d\n", a.balance)
	}
}

func (a *hdfc) mainBalance() {
	// a.balance := a.balance
	fmt.Println("Balance amount :", a.balance)
}

func main() {
	var account bank
	account = &hdfc{500}
	var deposited, withdraw, choice int
loop:
	for {
		fmt.Println("Press 1 - to deposite")
		fmt.Println("Press 2 - to withdrawal")
		fmt.Println("Press 3 - to know the balance")
		fmt.Println("Press 4 - to exit")
		fmt.Println("Enter your choice :")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Print("Enter the amount you want to deposite : ")
			fmt.Scanln(&deposited)
			account.deposite(deposited)
		case 2:
			fmt.Print("Enter the amount you want to withdraw :")
			fmt.Scanln(&withdraw)
			account.withdrawal(withdraw)
		case 3:
			account.mainBalance()
		case 4:
			fmt.Println("You are exited")
			break loop
		default:
			fmt.Println("Invalid input")
		}
	}
}
