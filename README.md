# Homework task
### Written in Golang

#### The task
Using any language/framework you desire, retrieve news items from the [GetNewsForApp endpoint](https://partner.steamgames.com/doc/webapi/ISteamNews) and create a web page to display them on-screen.  

Here is an example of this endpoint's use:
http://api.steampowered.com/ISteamNews/GetNewsForApp/v0002/?appid=570&count=4&maxlength=400&format=json

### Sidenotes
Initially I went with the html/template package, but ended up switching to the text/template one due to the Steam API response containing HTML-formatted content, which the html/template package escapes. This way the response can be displayed in its originally formatted way.