package helpers

import (
  "golang.org/x/crypto/bcrypt"
  "strings"
  "os"
  "context"
  "time"
  "fmt"
  "hash/crc32"

  "github.com/cloudinary/cloudinary-go"
  "github.com/cloudinary/cloudinary-go/api/uploader"
  beego "github.com/beego/beego/v2/server/web"
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

func StructsToCrc32(structs []interface{}) uint32 {
  var data []byte
  for i := 0; i < len(structs); i++ {
    data = append(data, []byte(fmt.Sprintf("%v", structs[i]))...)
  }
  return crc32.ChecksumIEEE(data)
}

func StructToCrc32(strct interface{}) uint32 {
  return crc32.ChecksumIEEE([]byte(fmt.Sprintf("%v", strct)))
}

func TplLastModifiedString(tpl string) string {
  file, err := os.Stat("views/" + tpl)
  if(err != nil){
    fmt.Println(err.Error())
    return "0"
  }

  return fmt.Sprintf("%d", file.ModTime().Unix())
}

func HandleEtag(c *beego.Controller, etag string) {
  if(strings.Trim(c.Ctx.Input.Header("If-None-Match"), "\"") != etag){
    c.Ctx.Output.Header("ETag", fmt.Sprintf("\"%s\"", etag))
  }else{
    c.Abort("304")
  }
}