package client

import (
	"net"
	"log"
	"os"
	"bufio"
	"fmt"
)

func Start_client() {

	//--------------------- Connection au serveur -------------------

	//Le client se connecte au serveur avec une adresse IP et un port défini.
	conn, err := net.Dial("tcp", "172.21.66.101:1337")
	if err != nil {
		log.Println("[LOG] Fatal error :",err.Error())
		os.Exit(1)
	}

	//---------------- Envoie d'un nombre au serveur ----------------

	for {
		//On demande au client d'écrire un nombre dans le terminal
		//qui sera envoyé au serveur
		nombre := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		mot, _ := nombre.ReadString('\n')


		//On envoie le nombre saisit par le client au serveur, ce nombre est contenue dans
		//la variable 'mot'
		writer := bufio.NewWriter(conn)
		_, err = writer.WriteString(mot)
		writer.Flush()
		if err != nil {
			log.Println("[LOG] Fatal error ",err.Error())
			os.Exit(1)
		}
	}

}

