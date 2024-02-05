/*
RCON commands used for Palworld.
The [Source RCON Protocol] from SteamCMD allows the moderation of a server without

	requiring an admin to log onto Palworld. The commands below mimic the commands entered into the in-game chatbox.
	RCON must be enabled for the server. Replace "Public_IP:Port" with the server's public IP address and RCON port
	and "AdminPassword" with the server's Admin Password before executing.

[Source RCON Protocol]: https://developer.valvesoftware.com/wiki/Source_RCON_Protocol
*/
package palwrldcmdsgo

import (
	"log"

	"github.com/gorcon/rcon"
)

func BanPlayer(IPAddress string, password string, steamID string) string {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	response, err := conn.Execute("BanPlayer " + steamID)
	if err != nil {
		log.Fatal(err)
	}

	return response

}

func Broadcast(IPAddress string, password string, message string) {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	response, err := conn.Execute("Broadcast " + message)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf(response)
}

func DoExit(IPAddress string, password string) string {
	/*
		Causes a FORCE SHUTDOWN of the Palworld Server. The server will automatically restart
		if being ran from a Linux Server as a Service (systemd)log.
	*/
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

func Info(IPAddress string, password string) string {
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

func KickPlayer(IPAddress string, password string, steamID string) string {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	response, err := conn.Execute("KickPlayer " + steamID)
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func ShowPlayers(IPAddress string, password string) string {

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

func Save(IPAddress string, password string) string {

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

func Shutdown(IPAddress string, password string, seconds string, message string) string {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	response, err := conn.Execute("Shutdown " + seconds + " " + message)
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func Test(IPAddress string, password string) (bool, error, string) {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	response, err := conn.Execute("info")
	if err != nil {
		log.Fatal(err)

		return false, err, response
	}

	log.Println(response)

	log.Printf("Connected Successfully!")

	return true, err, response
}
