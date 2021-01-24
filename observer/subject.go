package main

type subject interface {
	register(o observer)
	deregister(o observer)
	notifyALL()
}
