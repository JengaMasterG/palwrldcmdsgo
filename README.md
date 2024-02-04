# Dedicated Source RCON API Wrapper for Palworld
The [Source RCON Protocol from SteamCMD](developer.valvesoftware.com/wiki/Source_RCON_Protocol) allows the moderation of a server without requiring an admin to log onto Palworld. The commands below mimic the commands entered into the in-game chatbox. RCON must be enabled for the server. 

## Getting Started

### Install

### Usage
In main.go:

```go
package main

import (
    "palwordManager/cmds"
)

var IPAddress, password = "Public_IP:Port", "AdminPassword"

func main(){
    cmds.Test(IPAddress, password)
}
```

Before Executing:
- Replace `Public_IP:Port` with the server's public IP address and RCON port.
- Replace `AdminPassword` with the server's Admin Password.

Running the main.go with `go run .` will run the `Test` function to see if the server is running.