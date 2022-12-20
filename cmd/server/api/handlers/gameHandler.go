package handler

type GameHandler struct {
	RootUri string
}

const (
	gameRootURI = "/board"
)

func NewGameHandler() *GameHandler {
	return &GameHandler{
		RootUri: gameRootURI,
	}
}
