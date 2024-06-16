package helper

func CheckPasswordsMatch(storedPassword, suppliedPassword string) bool {
	return storedPassword == suppliedPassword
}
