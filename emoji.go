package flag_emoji

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
	return codeMap[c1] + codeMap[c2]
}
