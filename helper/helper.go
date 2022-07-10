package helper
import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidNames := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= uint(remainingTickets)
	return isValidEmail, isValidNames, isValidTickets
}