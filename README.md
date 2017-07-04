# HtmlRefresh

This is a sample project used for demoing deployments

It is written in [.NET Core](http://dot.net/core) using as few dependencies as possible :-)

# Options

The following options can be specified via environment variables:


|Name|Description|Default|
|-|-|-|
|`HTML_REFRESH_MESSAGE`|The message to display. Environment variables will be expanded (N.B. Use`%COMPUTERNAME%` on Windows)|`$HOSTNAME|
|`HTML_REFRESH_BGCOLOUR`|Colour to use for the page background|`azure`|
|`HTML_REFRESH_FGCOLOUR`|Colour to use for the page foreground (i.e. font colour)|`black`|
|`HTML_REFRESH_FONTSIZE`|Size to use for the font (will need tweaking based on the message)|`16vw`|
|`HTML_REFRESH_RATE`|The refresh rate to use for the page|`1`|


Additionally, you can configure `ASPNETCORE_URLS` to control the ports that the web server binds to, e.g. `http://localhost:5001'`(see [Url Prefixes](https://docs.microsoft.com/en-us/aspnet/core/fundamentals/servers/kestrel#url-prefixes) for more examples)
