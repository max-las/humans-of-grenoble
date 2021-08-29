package helpers

import (

  "golang.org/x/crypto/bcrypt"
  "strings"
  "os"
  "fmt"

)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func FirstWords(s string, n int) string {
  words := strings.Split(s, " ")
  res := words[0]
  for i := 1; i < len(words); i++ {
    if(i == n){
      break;
    }

    res = res + " " + words[i]
  }

  return res + "…"
}

func RemoveStaticByUrl(url string) bool{
  filePath := strings.TrimPrefix(url, "/")
  err := os.Remove(filePath)
  if(err != nil){
    fmt.Println("Failed to remove file " + filePath, err.Error())
    return false
  }
  return true
}
