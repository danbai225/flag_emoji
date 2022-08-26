package flag_emoji

import "strings"

var codeMap = map[string]string{
	"a": "🇦",
	"b": "🇧",
	"c": "🇨",
	"d": "🇩",
	"e": "🇪",
	"f": "🇫",
	"g": "🇬",
	"h": "🇭",
	"i": "🇮",
	"j": "🇯",
	"k": "🇰",
	"l": "🇱",
	"m": "🇲",
	"n": "🇳",
	"o": "🇴",
	"p": "🇵",
	"q": "🇶",
	"r": "🇷",
	"s": "🇸",
	"t": "🇹",
	"u": "🇺",
	"v": "🇻",
	"w": "🇼",
	"x": "🇽",
	"y": "🇾",
	"z": "🇿",
}

func Gen(c1, c2 string) string {
	return codeMap[strings.ToLower(c1)] + codeMap[strings.ToLower(c2)]
}
