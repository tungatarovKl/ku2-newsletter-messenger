# ku2-newsletter-messenger
Микросервис на go который делает рассылку пользователя telegram
```golang 
	http.HandleFunc("/newsletter", handleNewsletter) // newsletter routing
	http.ListenAndServe(":5000", nil) // Starting server on 5000 port
```
## func handleNewsletter(w http.ResponseWriter, r *http.Request) 
## Function bases on request type form response - Get or Post using ```switch``` statement. 




