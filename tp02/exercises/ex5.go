package exercises

type Agent interface {
	Start()
}

type PingAgent struct {
	ID string
	c  chan string
}

type PongAgent struct {
	ID string
	c  chan string
}

func Ex5() {

}
