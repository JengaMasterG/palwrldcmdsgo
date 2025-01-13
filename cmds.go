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

func Dial(IPAddress string, password string) (*rcon.Conn, error) {
	conn, err := rcon.Dial(IPAddress, password)
	return conn, err
}

func BanPlayer(IPAddress string, password string, steamID string) (string, error) {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Print("[WARN]: Error Occured")
		return "", err
	}
	defer conn.Close()

	response, err := conn.Execute("BanPlayer " + steamID)

	return response, err
}

func Broadcast(IPAddress string, password string, message string) error {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Print("[WARN]: Error Occured")
		return err
	}
	defer conn.Close()

	response, err := conn.Execute("Broadcast " + message)

	log.Print(response)

	return err
}

func DoExit(IPAddress string, password string) (string, error) {
	/*
		Causes a FORCE SHUTDOWN of the Palworld Server. The server will automatically restart
		if being ran from a Linux Server as a Service (systemd)log.
	*/
	log.Print("[WARN]:====SERVER FORCE SHUTDOWN STARTED====")

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Print("[WARN]: Error Occured")
		return "", err
	}
	defer conn.Close()

	response, err := conn.Execute("DoExit")

	log.Print("[WARN]: " + response)

	log.Printf("[WARN]:====SERVER FORCE SHUTDOWN COMPLETED====")

	return response, err
}

func Info(IPAddress string, password string) (string, error) {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Print("[WARN]: Error Occured")
		return "", err
	}
	defer conn.Close()

	response, err := conn.Execute("info")

	return response, err
}

func KickPlayer(IPAddress string, password string, steamID string) (string, error) {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Print("[WARN]: Error Occured")
		return "", err
	}
	defer conn.Close()

	response, err := conn.Execute("KickPlayer " + steamID)

	return response, err
}

func ShowPlayers(IPAddress string, password string) (string, error) {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Print("[WARN]: Error Occured")
		return "", err
	}
	defer conn.Close()

	response, err := conn.Execute("ShowPlayers")

	return response, err
}

func Save(IPAddress string, password string) (string, error) {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Print("[WARN]: Error Occured")
		return "", err
	}
	defer conn.Close()

	response, err := conn.Execute("Save")

	return response, err
}

func Shutdown(IPAddress string, password string, seconds string, message string) (string, error) {

	conn, err := rcon.Dial(IPAddress, password)
	if err != nil {
		log.Print("[WARN]: Error Occured")
		return "", err
	}
	defer conn.Close()

	response, err := conn.Execute("Shutdown " + seconds + " " + message)

	return response, err
}

func Test(IPAddress string, password string) (string, error) {

	response, err := Info(IPAddress, password)

	return response, err
}
