package server

type Server interface {
	Up(string)
}

func Instance() Server {

}


