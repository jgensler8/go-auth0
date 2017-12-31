package main

import (
  "log"
  "bytes"
  "io/ioutil"
  auth0 "github.com/jgensler8/go-auth0/generated/client"
)

func main(){
  cfg := auth0.DefaultTransportConfig().WithHost("vastorchard.auth0.com")
  client := auth0.NewHTTPClientWithConfig(nil, cfg)


  buf := bytes.NewBufferString("")
  _, err := client.Operations.GetPEM(nil, buf)
  if err != nil {
    log.Print(err)
  } else {
    cert, err := ioutil.ReadAll(buf)
    if err != nil {
      log.Print(err)
    } else {
      log.Printf("%s", cert)
    }
  }
}
