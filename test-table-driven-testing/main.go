package main

func pingPong(n int) string {
	switch {
	case n%2 == 0 && n%3 == 0:
		return "ping-pong"
	case n%2 == 0:
		return "ping"
	case n%3 == 0:
		return "pong"
	default:
		return ""
	}
}
