# Thoughts and design decisions
We know that every 15 minutes, each player queries an API to see if a new version is available and then updates itself.
So this part is already automated.

My task was to make sure we have a tool, which would allow to push updates to players any time and on a large scale.
The design I came up with includes the following:

1. A middle-tier HTTP server which would accept the requests from client and forward it to the appropriate player and return back the result of update with the JSON, showing the updated player version.
2. A tool, which would read a content of CSV file and push latest software version (JSON) to macaddresses specified. This tool would use the HTTP server endpoint to route the requests.

Potentially we could allow each player have those endpoints locally but it would require maintaining each service separatly on each player, so I decided to have server in the middle.
I didn't implement REAL routing, since I don't have players on network, but this won't be too hard to implement with the current baseline I developed.
Each request is being verified for header and body validity. ClientID and token should be provided. Ideally we would ise JWS tokens, but for the assignement I use Token structure with token itself and expiration flag. Token is stored in a plain view, which of course is not secure. 

I made this tool cross-platform, so Go is perfect, since we can build for number of different OS. I used Windows OS, but MacOS and Linux shoulnd't be a problem. (It builds with no problem, but didn't have a chance to test runtime). 
Unit tests are designed to test UpdateHandler for different scenarios to make sure 200, 409, 404, 401 error codes are thrown when expected.

# Brief package overview
##### Updater Service contains request handler and verificator portion (updater.go, verificator.go). 
##### Updater Tool (Tool) is independent command line utility, so it has separate main package under /tool/, not related to Updater Service. 
##### Common package: provides common functions used by updater tool and updater service.

### Documentation

Go provides an internal documentation tool: 
```
go doc <package name> 
```
The comments I used were intended to get the most of internal go doc functionality. Some extra effort is required to make top-notch, as somehow it doesn't properly print out public funcion definitions.


# Things to improve on

* Internal documentation
* Implement proper routing
* Implement proper token authentication
* 500 Internal Server error is not returned to client.
* Improve tool to accept single MAC address.





