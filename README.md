# go-mailer-sandbox

1. generate a key with `cipher.GenerateKey()` method.
2. save the key in `.env`
3. cipher your email password with `cipher.Encrypt()` method.
4. add the password into `.env` file
5. change the `from`, `to` fields in `main.go`
6. `go run main.go`
