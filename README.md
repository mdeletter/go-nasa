# Go-nasa

Enjoy all the great features of Go (e.g. goroutines, channels, http packages) in this awesome (small) workshop.

# Before you start

To launch your spaceship execute the following command inside a new terminal window.

```
$ go run spaceship/spaceship.go
```

Double-check if the spaceship is correctly launched by browsing to:

[http://localhost:8000/](http://localhost:8000/)

## Disclaimer
This is not your typical copy-paste workshop/tutorial.
Instead use the official Go language documentation, Go-by-example or Google to finish this mission.

# Your mission

Last night, for some unknown reason, one of NASA's spaceships was launched.
All communication channels with the spaceship are dead, besides 1 HTTP uplink.
It's your task to investigate what happened using Go.
 
Enjoy and don't forget: *To infinity…and beyond!*

## 1. Reach out to spaceship

Open ```groundcontrol.go```, search for reachout() and start coding to connect to the spaceship.

Check out: [GO HTTP client](https://golang.org/src/net/http/client.go), Defer, multiple return parameters.

To check your groundcontrol code, execute this:
```
$ go run groundcontrol.go
```
Open [Ground control](http://localhost:8001/)

## 2. Add more capacity
First checkout branch ```step2```

Use a couple of workers to throw more capacity to the problem. Also tune the total of workers to prevent TCP connection errors.
Check out: go routines, wait groups, method pointer receivers.

If you implement this correctly and you wait 10 minutes, you see what happened inside the spaceship.  
 
### Extra: With the speed of light
You probably also noted that some connections are really slow and holding us down. 
Try to implement timeouts with Go channels to speed up our process.

Check out: select channel, timeout channel.

### To infinity
Launch 4 more spaceships and gathering all the information inside groundcontrol.

Check out: maps, channels, goroutines

# Thanks

Hopefully you enjoyed this short mission and experienced some great benefits of Go.
That's all folks.
