// game/game.go
package game

type Game interface {
	Play(parameters ...interface{})
}
