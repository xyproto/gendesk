package main

// TODO: Use an external file to read the mappings from (possibly JSON)

const (
	// Decides the order of the keyword/category checks
	// (try to order from the more specific/specialized categories to the more general)
	model3d = iota
	multimedia
	graphics
	network
	email
	audiovideo
	office
	editor
	science
	vcs
	arcadegame
	actiongame
	adventuregame
	logicgame
	boardgame
	game
	programming
	system
	last
)

var (
	keywordmap = map[int][]string{
		model3d:       {"rendering", "modeling", "modelling", "modeler", "render", "raytracing"},
		multimedia:    {"non-linear", "audio", "sound", "graphics", "demo"},
		graphics:      {"draw", "pixelart", "animated"},
		network:       {"network", "p2p", "browser", "remote"},
		email:         {"gmail"},
		audiovideo:    {"synth", "synthesizer", "ffmpeg"},
		office:        {"ebook", "e-book", "spreadsheet", "calculator", "processor", "documents"},
		editor:        {"editor"},
		science:       {"gps", "inspecting", "molecular", "mathematics"},
		vcs:           {"git"},
		arcadegame:    {"combat", "arcade", "racing", "fighting", "fight", "shooter"},
		actiongame:    {"shooter", "fps"},
		adventuregame: {"roguelike", "rpg"},
		logicgame:     {"puzzle"},
		boardgame:     {"board", "chess", "goban", "chessboard", "checkers"},
		// "emulator" and "player" aren't always for games, but those cases will be
		// picked up by one of the other categories first, as orderd by the constants above
		game:        {"game", "rts", "mmorpg", "emulator", "player"},
		programming: {"code", "ide", "programming", "develop", "compile", "interpret", "valgrind"},
		system:      {"sensor", "bus", "calibration", "usb", "file"},
	}
	categorymap = map[int]string{
		model3d:       "Application;Graphics;3DGraphics",
		multimedia:    "Application;Multimedia",
		graphics:      "Application;Graphics",
		network:       "Application;Network",
		email:         "Application;Network;Email",
		audiovideo:    "Application;AudioVideo",
		office:        "Application;Office",
		editor:        "Application;Development;TextEditor",
		science:       "Application;Science",
		vcs:           "Application;Development;RevisionControl",
		arcadegame:    "Application;Game;ArcadeGame",
		actiongame:    "Application;Game;ActionGame",
		adventuregame: "Application;Game;AdventureGame",
		logicgame:     "Application;Game;",
		boardgame:     "Application;Game;BoardGame",
		game:          "Application;Game",
		programming:   "Application;Development",
		system:        "Application;System",
	}
)

// GuessCategory will try to guess which category an application belongs to,
// given a short package description.
// If not guess is made, just "Application" will be returned.
func GuessCategory(pkgdesc string) string {
	var keywordList []string
	for key := 0; key < last; key++ {
		keywordList = keywordmap[key]
		if keywordsInDescription(pkgdesc, keywordList) {
			return categorymap[key]
		}

	}
	return "Application"
}
