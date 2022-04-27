package hcregistry

var Alive = true
var Ready = true

func SetAlive(state bool) {
	Alive = state
}

func SetReady(state bool) {
	Ready = state
}
