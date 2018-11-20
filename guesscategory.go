package main

// TODO: Use an external file to read the mappings from (possibly JSON)

const (
	// Decides the order of the keyword/category checks
	// (try to order from the more specific/specialized categories to the more general)
	model3d_kw = iota
	multimedia_kw
	graphics_kw
	network_kw
	email_kw
	audiovideo_kw
	office_kw
	editor_kw
	science_kw
	vcs_kw
	arcadegame_kw
	actiongame_kw
	adventuregame_kw
	logicgame_kw
	boardgame_kw
	game_kw
	programming_kw
	system_kw
	last_kw
)

var (
	keywordmap = map[int][]string{
		model3d_kw:       []string{"rendering", "modeling", "modelling", "modeler", "render", "raytracing"},
		multimedia_kw:    []string{"non-linear", "audio", "sound", "graphics", "demo"},
		graphics_kw:      []string{"draw", "pixelart"},
		network_kw:       []string{"network", "p2p", "browser"},
		email_kw:         []string{"gmail"},
		audiovideo_kw:    []string{"synth", "synthesizer", "ffmpeg"},
		office_kw:        []string{"ebook", "e-book", "spreadsheet", "calculator", "processor", "documents"},
		editor_kw:        []string{"editor"},
		science_kw:       []string{"gps", "inspecting", "molecular", "mathematics"},
		vcs_kw:           []string{"git"},
		arcadegame_kw:    []string{"combat", "arcade", "racing", "fighting", "fight", "shooter"},
		actiongame_kw:    []string{"shooter", "fps"},
		adventuregame_kw: []string{"roguelike", "rpg"},
		logicgame_kw:     []string{"puzzle"},
		boardgame_kw:     []string{"board", "chess", "goban", "chessboard", "checkers"},
		// "emulator" and "player" aren't always for games, but those cases will be
		// picked up by one of the other categories first
		game_kw:        []string{"game", "rts", "mmorpg", "emulator", "player"},
		programming_kw: []string{"code", "ide", "programming", "develop", "compile", "interpret", "valgrind"},
		system_kw:      []string{"sensor", "bus", "calibration", "usb", "file"},
	}
	categorymap = map[int]string{
		model3d_kw:       "Application;Graphics;3DGraphics",
		multimedia_kw:    "Application;Multimedia",
		graphics_kw:      "Application;Graphics",
		network_kw:       "Application;Network",
		email_kw:         "Application;Network;Email",
		audiovideo_kw:    "Application;AudioVideo",
		office_kw:        "Application;Office",
		editor_kw:        "Application;Development;TextEditor",
		science_kw:       "Application;Science",
		vcs_kw:           "Application;Development;RevisionControl",
		arcadegame_kw:    "Application;Game;ArcadeGame",
		actiongame_kw:    "Application;Game;ActionGame",
		adventuregame_kw: "Application;Game;AdventureGame",
		logicgame_kw:     "Application;Game;",
		boardgame_kw:     "Application;Game;BoardGame",
		game_kw:          "Application;Game",
		programming_kw:   "Application;Development",
		system_kw:        "Application;System",
	}
)

// GuessCategory will try to guess which category an application belongs to,
// given a short package description.
// If not guess is made, just "Application" will be returned.
func GuessCategory(pkgdesc string) string {
	var keywordList []string
	for key := 0; key < last_kw; key++ {
		keywordList = keywordmap[key]
		if keywordsInDescription(pkgdesc, keywordList) {
			return categorymap[key]
		}

	}
	return "Application"
}
