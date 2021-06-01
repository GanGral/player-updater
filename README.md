# player-updater

Player Updater allows automation of the update of multiple music players. 

This repository contains:
#### Updater Service
Middle-tier Service to process incoming requests to update the players. Supports Updater Tool requests, as well as any other means of sending HTTP requests discussed later.
#### Updater Tool
Command line tool to automate the update of multiple music players.

## Installation

Clone the repository.

Compile the updater_service from /player-updater/ directory:

Windows OS:
```bash
$ GOOS=windows GOARCH=amd64 go build -o updater_service.exe
```

Compile the updater tool from /player-updater/tool/ directory:

Windows OS:
```bash
$ GOOS=windows GOARCH=amd64 go build -o tool.exe
```
## Usage
Start Updater Service. This would open a local listener on port 8457 to accept PUT requests to the following endpoint, where macaddress is actual player MAC-address
```
PUT /profiles/clientId:{macaddress}
```
### Updater Tool
Command line tool, which accepts the following parameters and forwards PUT requests to updater service for each player, specified in csv file (macaddress list).
```
  -csvpath string
        Specify the path to the player's Mac Addresses csv file (default "../players.csv")
  -port int
        specifies the updater_service port (default 8457)
  -profile string
        Specify the location of json file with up-to-date player profile (default "../tool/currentVersion.json")
 ```
 ### Postman/Fiddler or any other applications capable of sending HTTP requests
 ```
PUT /profiles/clientId:{macaddress}
```

Required header fields:
```
Content-Type: application/json
x-client-id: required
x-authentication-token: required
```
Make sure to provide a valid body with the target profile verison of a player, for example:
```
{
  "profile": {    
    "applications": [
      {
        "applicationId": "music_app"
        "version": "v1.4.10"
      },
      {
        "applicationId": "diagnostic_app",
        "version": "v1.2.6"
      },
      {
        "applicationId": "settings_app",
        "version": "v1.1.5"
      }
    ]
  }
}
```
### Curl request
```
curl -X PUT -d '{"profile":{"applications":[{"applicationId":"music_app","version":"v1.4.10"},{"applicationId":"diagnostic_app","version":"v1.2.6"},{"applicationId":"settings_app","version":"v1.1.5"}]}}' -H "Content-Type: application/json" -H "X-Client-Id: dummy" -H "X-Authentication-token: dummy" http://localhost:8457/profiles/clientId:a1:bb:cc:dd:ee:ff --verbose
```

### Unit Tests
To run unit test from /player-updater/ run
```
go test -v
```
