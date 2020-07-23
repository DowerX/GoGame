package main

type component interface {
	onStart()
	onUpdate()
	onFixedUpdete()
}
