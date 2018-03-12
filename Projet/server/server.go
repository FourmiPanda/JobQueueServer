package server

import (
	"log"
	"net"
	"os"
	"bufio"
	"strconv"
	"sync"
	"fmt"
)

type Tache struct {
	valeur int
}

var wg sync.WaitGroup
var fromCollector chan Tache  = make(chan Tache,5)
var availableWorkers chan chan Tache = make(chan chan Tache,5)


// La fonction start démarre le serveur en instanciant les différent travailleurs, le collecteur et le repartiteur
func Start()  {
	fmt.Println("[Running] func Start() l.23")
	wg.Add(1)

	go collecteur("1337")

	go repartiteur()

	var workChan1 chan Tache = make(chan Tache,1)
	availableWorkers <- workChan1
	go travailleur(workChan1)
	var workChan2 chan Tache = make(chan Tache,1)
	availableWorkers <- workChan2
	go travailleur(workChan2)
	var workChan3 chan Tache = make(chan Tache,1)
	availableWorkers <- workChan3
	go travailleur(workChan3)

	wg.Wait()

}


// La fonction collecteur attend une connexion et envoie les requetes reçues dans le channel fromCollector
func collecteur(port string)  {
	defer wg.Done()

	log.Println("[LOG] Function collecteur starting")
	listener,err := net.Listen("tcp",":"+port)
	if checkError(err) {
		os.Exit(1)
	}


	// Boucle infinie du serveur qui attend une connexion
	for {
		log.Println("[LOG] Waiting for connection")
		connexion, err := listener.Accept()
		if checkError(err) {
			connexion.Close()
		}

		readerClient := bufio.NewReader(connexion)
		messageClient, err := readerClient.ReadString('\n')
		checkError(err)
		read_line := messageClient
		read_line = read_line[:len(read_line)-1]

		val,err := strconv.Atoi(read_line)
		if !checkError(err)  {
			fromCollector <- Tache{val}
			log.Println("[LOG] Ajout d'un int au channel fromCollector")

		} else {
			log.Println("[LOG] Le message n'a pas été ajouté au channel fromCollector")
		}

	}
}


func travailleur(workChan chan Tache){
	log.Println("[LOG] Création d'un travailleur")
	for{
		msg := <- workChan
		log.Println("[LOG] Un travailleur a reçu une tache")
		for i:=0;i<msg.valeur;i++{
			fmt.Println(i)
			//log.Println("[LOG] Un tour de boucle")
		}
		availableWorkers <- workChan
		log.Println("[LOG] Un travailleur est de nouveau opérationel")
	}

}

func repartiteur(){
	log.Println("[LOG] Lancement du repartiteur")
	for{
		tache := <- fromCollector
		log.Println("[LOG] Une nouvelle tache va etre repartie")
		worker := <- availableWorkers
		worker <- tache
	}
}



func checkError(err error) bool {
	if err != nil {
		log.Println("[LOG] Fatal error ", err.Error())
		return true
	}
	return false
}

