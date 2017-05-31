# darts-go

## BUILD

rPi: env GOOS=linux GOARCH=arm go build server.go
 
## RUN
 
 Start the server command:
 
 ```bash
./server
```

## Usage

I use the domain name darts3. You may replace it with the ip address of the server:

The system consists of two different user interfaces. The game UI and the admin UI.

The game UI does not handle input from the user, only shows the data received from the
server. The game UI is at the address [http://darts3:8080/game/scoreboard](http://darts3:8080/game/scoreboard)

If there is no game started, then the game UI shows a barcode, that points to the admin UI:

![game ui started](https://github.com/IPeter/darts-go/blob/master/doc/gameUI01.png)

The barcode shows the url of the admin UI:



The first step is to enter the name of the players ([http://darts3:8080/admin/setPlayer](http://darts3:8080/admin/setPlayer)):
![game ui started](https://github.com/IPeter/darts-go/blob/master/doc/adminUI01.png)

The start button starts the game. The admin ui will show every throw. The user can fix the
recognized score in case of a wrong recognition. The game UI will show the current status
of the game.


![game ui started](https://github.com/IPeter/darts-go/blob/master/doc/screenshot01.png)

![game ui started](https://github.com/IPeter/darts-go/blob/master/doc/adminUI02.png)

## Connecting other recognition software

This is the path for entering a recognized throw:
 
 http://darts3:8080/cam/command?num=2&modifier=1

where 

* num is a number from: [1 .. 20, 25]
* modifier is in : [-1, 0, 1, 2, 3]. 
  * 1: simple
  * 2: double
  * 3: triple
  * 0: out of bounds
  * -1: darts are removed before the third throw. This will make the game UI switch player, but not implemented yet.



