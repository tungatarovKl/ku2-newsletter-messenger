# ku2-newsletter-messenger

Микросервис на go который делает рассылку пользователя telegram

```golang
	http.HandleFunc("/newsletter", handleNewsletter) // newsletter routing
	http.ListenAndServe(":5000", nil) // Starting server on 5000 port
```

func handleNewsletter(wl http.ResponseWriter, r \*http.Request)
Function bases on request type form response - Get or Post using `switch` statement.

## Как запустить бота?

1. Ввести токен бота в config/local.toml
2. `go run main.go -config ../../config/local.toml`
