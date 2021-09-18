package controllers

import (
	"sort"
	"os"
  "time"
  "strconv"
	"strings"
  "crypto/sha1"
  "encoding/hex"
	beego "github.com/beego/beego/v2/server/web"
)

type CloudinaryController struct {
	beego.Controller
}

func (c *CloudinaryController) Post() {
	params := c.Ctx.Request.Form
	timestamp := time.Now().Unix()
	params["timestamp"] = []string{strconv.FormatInt(timestamp, 10)}

  var res = CldCResponse{
    ApiKey: os.Getenv("CLOUDINARY_API_KEY"),
    Timestamp: timestamp,
  }

  res.Sign(params);

  c.Data["json"] = &res
  c.ServeJSON()
}

type CldCResponse struct {
  ApiKey string `json:"api_key"`
  Timestamp int64 `json:"timestamp"`
  Signature string `json:"signature"`
}

func (res *CldCResponse) Sign(params map[string][]string) {
	keys := make([]string, 0, len(params))
	for k := range params {
    keys = append(keys, k)
  }
	sort.Strings(keys)

	str := ""
	for _, k := range keys {
    str = str + k + "=" + params[k][0] + "&"
  }
	str = strings.TrimSuffix(str, "&")

  str = str + os.Getenv("CLOUDINARY_API_SECRET")

  h := sha1.New()
  h.Write([]byte(str))
  res.Signature = hex.EncodeToString(h.Sum(nil))
}
