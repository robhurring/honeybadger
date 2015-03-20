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

## Example Usage

```go
// get your token on the https://app.honeybadger.io/users/edit page
token := os.Getenv("HONEYBADGER_API_TOKEN")
honeybadger := honeybadger.New(token)

projects, err := honeybadger.Projects()
if err != nil {
  panic(err)
}

fmt.Printf("You have %d projects in honeybadger!\n", len(projects.Results))
```

## API Documentation

For more information about the API and returned data, you can view the api docs here: https://www.honeybadger.io/documentation/read_api

# Contributing

* Add your features, submit a pull-request.
* Check `rake filesizes` to make sure we aren't bloating assets