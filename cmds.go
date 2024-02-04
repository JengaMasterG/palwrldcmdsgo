/*
RCON commands used for Palworld.
The [Source RCON Protocol] from SteamCMD allows the moderation of a server without
 requiring an admin to log onto Palworld. The commands below mimic the commands entered into the in-game chatbox. 
 RCON must be enabled for the server. Replace "Public_IP:Port" with the server's public IP address and RCON port 
 and "AdminPassword" with the server's Admin Password before executing.

[Source RCON Protocol]: https://developer.valvesoftware.com/wiki/Source_RCON_Protocol
*/
package cmds

import (
	"log"

	"github.com/gorcon/rcon"
)

var IPAddress, password = "IPAddress", "Password"
func BanPlayer(){

}

func Broadcast(){

}

func DoExit()(string){
	log.Printf("[WARN]:====SERVER FORCE SHUTDOWN STARTED====")

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	response, err := conn.Execute("DoExit")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[WARN]: " + response)

	log.Printf("[WARN]:====SERVER FORCE SHUTDOWN COMPLETED====")

	return response
}

func Info()(string){
	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	response, err := conn.Execute("info")
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func KickPlayer(){

}

func ShowPlayers()(string){

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	response, err := conn.Execute("ShowPlayers")
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func Save()(string){

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	response, err := conn.Execute("Save")
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func Shutdown(){

}

func Test() {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	response, err := conn.Execute("info")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)

	log.Printf("Connected Successfully!")
}
