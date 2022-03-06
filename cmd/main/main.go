// Copyright 2021 ThingWave AB.  All rights reserved.
package main

import (
  "os"
  "log"
  "io/ioutil"
  "flag"
  "fmt"
  "time"
  "errors"
  "crypto/x509"
  "crypto/tls"
  "net/http"
)

func main() {
  command := flag.String("cmd", "test-sr", "Eclipse Arrowhead command to execute")
  targetUri := flag.String("sr", "https://127.0.0.1:8443/serviceregistry/echo", "Service registry URI")
  ca := flag.String("cafile", "", "Root CA PEM file")
  cert := flag.String("cert", "", "Client certificate PEM file")
  key := flag.String("key", "", "Client certificate key PEM file")
  flag.Parse()

  if (*ca != "" && (*cert == "" || *key == "")) || (*cert != "" && (*ca == "" || *key == "")) || (*key != "" && (*cert == "" || *ca == ""))   {
    fmt.Println("Error: missing arguments!")
    fmt.Println("Use 'ahctl --help' to see usage.")
    return
  }

  if *command != "test-sr" {
      fmt.Printf("Error: unknown command '%s'\n", *command)
      return
  }

  if *ca != "" && *cert != "" && *key != "" {
    ok, _ := fileExists(*ca)
    if !ok {
      fmt.Printf("Error: ca file %s does not exist\n", *ca)
      return
    }
    ok, _ = fileExists(*cert)
    if !ok {
      fmt.Printf("Error: certificate file %s does not exist\n", *cert)
      return
    }
    ok, _ = fileExists(*key)
    if !ok {
      fmt.Printf("Error: key file %s does not exist\n", *key)
      return
    }

    caCert, err := ioutil.ReadFile(*ca)
    if err != nil {
      log.Fatalf("Error opening cert file %s, Error: %s", *ca, err)
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    myCert, err := tls.LoadX509KeyPair(*cert, *key)
    if err != nil {
        log.Fatalf("Error creating x509 keypair from client cert file %s and client key file %s", *cert, *key)
    }

    t := &http.Transport{
      TLSClientConfig: &tls.Config{
        Certificates: []tls.Certificate{myCert},
        RootCAs:      caCertPool,
        InsecureSkipVerify: false,
      },
    }

    client := http.Client{Transport: t, Timeout: 10 * time.Second}
    data, err := getData(client, *targetUri)
    if err == nil {
      fmt.Println(data)
    }

  }
}

func getData(client http.Client, uri string) (string, error) {
  request, err := http.NewRequest(http.MethodGet, uri, nil)
  if err != nil {
    log.Fatalf("request failed : %v", err)
  }

  resp, err := client.Do(request)
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err == nil {
    return string(body), nil
  } else {
    fmt.Printf("final error: %v\n", err);
    return "", err
  }
}

func fileExists(fileName string) (bool, error) {
    _, err := os.Stat(fileName)
    if err == nil {
        return true, nil
    }
    if errors.Is(err, os.ErrNotExist) {
        return false, nil
    }
    return false, err
}

