package server

import (
  "fmt"
  "net"
  "bufio"
  "log"
  "sync"
  "strconv"
  "os"
)

type Tache struct {
  valeur int
}

var wg sync.WaitGroup
var fromCollector chan Tache  = make(chan Tache,5)


// La fonction start démarre le serveur en instanciant les différent travailleurs, le collecteur
func Start()  {
  fmt.Println("[Running] func Start() l.23")
  wg.Add(1)


  go collecteur("1337")


  wg.Wait()

}


// La fonction collecteur attend une connexion et envoie les requetes reçues dans le channel fromCollector
// TO - DO : Coupure de connexion après envoie d' UN message
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
      continue
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


  func checkError(err error) bool {
    if err != nil {
      log.Println("[LOG] Fatal error ", err.Error())
      return true
    }
    return false
  }
