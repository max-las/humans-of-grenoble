package helpers

import (

  "golang.org/x/crypto/bcrypt"
  "strings"
  "os"
  "context"
  "time"
  "fmt"

  "github.com/cloudinary/cloudinary-go"
  "github.com/cloudinary/cloudinary-go/api/uploader"

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

func RemoveCldAsset(public_id string) (*uploader.DestroyResult, error){
  fmt.Println("Removing asset", public_id)

  ctxBackground := context.Background()
  ctxWithTimeout, _ := context.WithTimeout(ctxBackground, time.Duration(3)*time.Second)

  cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
  return cld.Upload.Destroy(ctxWithTimeout, uploader.DestroyParams{
    PublicID: public_id,
    ResourceType: "image",
    Invalidate: true,
  })
}

func MinInt(a, b int) int {
  if(a < b){
    return a
  }else{
    return b
  }
}