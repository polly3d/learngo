package mp

type player interface {
	Play(source string)
}

// Play ...
func Play(source, mtype string) {
	var player player
	
	player.Play(source)
}
