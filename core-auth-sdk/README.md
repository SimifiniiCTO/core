# Blackspace Auth-SDK

Keratin AuthN is an authentication service that keeps you in control of the experience without forcing you to be an expert in web security.

This library provides utilities to help integrate with a Go application. You may also need a client for your frontend, such as [https://github.com
/keratin/authn-js](https://github.com/keratin/authn-js).

[![Godoc](https://godoc.org/github.com/keratin/authn-go/authn?status.svg)](https://godoc.org/github.com/keratin/authn-go/authn)
[![Gitter](https://badges.gitter.im/keratin/authn-server.svg)](https://gitter.im/keratin/authn-server?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![Build Status](https://travis-ci.org/keratin/authn-go.svg?branch=master)](https://travis-ci.org/keratin/authn-go)
[![Go Report](https://goreportcard.com/badge/github.com/keratin/authn-go)](https://goreportcard.com/report/github.com/keratin/authn-go)

## Installation

```bash
go get github.com/Lens-Platform/Platform/src/libraries/core/core-auth-sdk
```

## Example

```go
package main

import (
  "fmt"
  sdk "github.com/Lens-Platform/Platform/src/libraries/core/core-auth-sdk"
)

var jwt1 = `<your test jwt here>`
var accountID = `<test ID>`

func main() {
  err := sdk.NewClient(sdk.Config{
    // The AUTHN_URL of your Keratin AuthN server. This will be used to verify tokens created by
    // AuthN, and will also be used for API calls unless PrivateBaseURL is also set.
    Issuer:         "https://issuer.example.com",

    // The domain of your application (no protocol). This domain should be listed in the APP_DOMAINS
    // of your Keratin AuthN server.
    Audience:       "application.example.com",

    // Credentials for AuthN's private endpoints. These will be used to execute admin actions using
    // the Client provided by this library.
    //
    // TIP: make them extra secure in production!
    Username:       "<Authn Username>",
    Password:       "<Authn Password>",

    // RECOMMENDED: Send private API calls to AuthN using private network routing. This can be
    // necessary if your environment has a firewall to limit public endpoints.
    PrivateBaseURL: "http://private.example.com",
  })
  fmt.Println(err)

  // SubjectFrom will return an AuthN account ID that you can use as to identify the user, if and
  // only if the token is valid.
  sub, err := sdk.SubjectFrom(jwt1)
  fmt.Println(sub)
  fmt.Println(err)

  // LockAccount will lock an AuthN account using the same ID that you saw in the user's JWT when
  // they signed up. That account will be unable to log in until it is unlocked.
  //
  // See the godocs for all actions that you can take on an account.
  err = sdk.LockAccount(accountID)
  fmt.Println(err)
}
```
