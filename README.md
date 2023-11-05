# Homework task

#### The task
> Using any language/framework you desire, retrieve news items from the [GetNewsForApp endpoint](https://partner.steamgames.com/doc/webapi/ISteamNews) and create a web page to display them on-screen.  
>
> Here is an example of this endpoint's use:
> http://api.steampowered.com/ISteamNews/GetNewsForApp/v0002/?appid=570&count=4&maxlength=400&format=json

#### Tech stack
* Server/client functionality: Golang with `net/HTTP` and `html(text)/template` packages
* CSS/JS Toolkit: [Bootstrap](https://getbootstrap.com/)
* Unit tests: Go `testing` package
* GitHub Actions workflow for building and running the tests when something gets pushed to `master`

### Build and run the server
To start the application, run  `go run cmd/main.go`, or build using `go build cmd/main.go` then start the binary.
The app can be accessed on http://localhost:8080 by default.

### Tests
To execute all the unit tests manually, run  `go test -v ./...`.

### Sidenotes
*Initially I went with the html/template package but ended up switching to the text/template one due to the Steam API response containing HTML-formatted content, which the html/template package escapes. This way the response can be displayed in its originally formatted way.*

### Possible next steps
The number of unit tests could be extended with the use of `ginkgo` and `gomega` frameworks.
The lists within the returned news-components could be parsed and displayed like actual list items using the tags in square brackets.
