package main

import(
    "fmt"
    "os"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    var password string

    file, err := os.OpenFile("../password.txt", os.O_WRONLY, 0600)

    if os.IsNotExist(err) {

        fmt.Print("Type the login password (DONT use spaces): ")
        fmt.Scanln(&password)

        hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

        os.WriteFile("../password.txt", hashed, 0600)

    } else {

        defer file.Close()

        fmt.Print("Type current password: ")
        fmt.Scanln(&password)

        hashed, _ := os.ReadFile("../password.txt")

        if err := bcrypt.CompareHashAndPassword(hashed, []byte(password)); err != nil {
            fmt.Println("Password Incorrect")
            return
        }

        fmt.Print("Type new password: ")
        fmt.Scanln(&password)

        newPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

        file.Write(newPass)

    }
}
