```
                  #@#@;`
                '##@ #@###@,
               +##'@###@+:. `
               ###@#@.                      A HoneyBadger API Client
              @####+                          (That _does_ give a F!@$%)
              @###@
              ####      .+@@#;,...,:;:`
              @###    @@; `++,`   `,'@@#:
              ####  +#.                 `'@@'
              `###,+#                        ;@@@#####@@@@@@@@.
               @####+                              `.`         '
               `####:
                ####@                                         ;@@#
                 #####                          :@+.       :@#`+##.
                 ;####@`                     ;@` ;#@@##########@###;
                 .####;@@+`               :@,.@#####################.
                  #####@'`,+@@@@@##;. :@#,:@########################
                  #########@@@@@@@@@#+#@@################@@@@@##@@#
                  @################################@+.
                  :#############################@,
                   @##########################@`  +@
                   '#########@@##@@###########  @##@
                    #######@       .########@ '#####
                    ######@  @####@  #######  @#####`
                    @#####:  ######  ######@  #######
                    @#####`  ######  ;#####@  ,######
                    @#####:  @#####. `######:  ######@
                     @####@   @####@  +@#####  `@#####:
```

# Honeybadger API Client

A small, lightweight honeybadger API mapping for Golang. This was mapped against the [V1 API](https://www.honeybadger.io/documentation/read_api), so YMMV.

## Installing

`go get github.com/robhurring/honeybadger`

## Example Usage

```go
package main

import (
  "fmt"
  "os"

  "github.com/robhurring/honeybadger"
)

func main() {
  // get your token on the https://app.honeybadger.io/users/edit page
  token := os.Getenv("HONEYBADGER_API_TOKEN")
  honeybadger := honeybadger.New(token)
  params := honeybadger.Params{"page": "1"}

  projects, err := honeybadger.Projects(params)
  if err != nil {
    panic(err)
  }

  fmt.Printf("You have %d projects in honeybadger!\n", len(projects.Results))
}

// HONEYBADGER_API_TOKEN=<my-token> go run main.go
```

## To-Do

1. Implement pagination

## API Documentation

For more information about the API and returned data, you can view the api docs here: https://www.honeybadger.io/documentation/read_api

# Contributing

* Add your features, submit a pull-request.
* Check `rake filesizes` to make sure we aren't bloating assets