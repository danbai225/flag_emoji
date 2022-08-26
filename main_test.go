package flag_emoji

import "testing"

func TestGen(t *testing.T) {
	println(Gen("c", "n"))
	println(Gen("a", "d"))
	println(Gen("a", "q"))
	println(Gen("c", "a"))
}
