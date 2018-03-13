# JobQueueServer

*Basic client -> server architecture in Go*

## 1. Description

Ce programme représente une architecture client/serveur basique. Dans l'exemple
fourni par le programme, un client peut envoyer un nombre vers le serveur. Quand
    le serveur reçoit un nombre il est collecté dans un channel, le nombre est ensuite
    envoyé à un travailleur disponible. Le travailleur a comme tache d'executer une
    boucle le nombre de fois qu'il lui a été demandé. Une fois la tache fini le travailleur
    est de nouveau disponible.

## 2. Fonctionnement

    Pour lancer le programme, il faut executer le fichier "main/runServer.go" pour lancer le
    serveur ou bien "main/runClient.go" pour lancer 1 ou plusieur client.
    Une fois le client lancé un invité de commande est lancé en attente d'un nombre à envoyé
    au serveur.

## 3. Aide

    Si vous avez besoin de plus ample explication par rapport au code veuillez nous contacter
    sur l'adresse e-mail suivante jobqueueserver@help.fake

