package main

// TODO Use a large hash map or a json data file to store the mappings

var (
	model3d_kw    = []string{"rendering", "modeling", "modeler", "render", "raytracing"}
	multimedia_kw = []string{"non-linear", "audio", "sound", "graphics", "draw", "demo"}
	network_kw    = []string{"network", "p2p", "browser"}
	audiovideo_kw = []string{"synth", "synthesizer"}
	office_kw     = []string{"ebook", "e-book"}
	editor_kw     = []string{"editor"}
	science_kw    = []string{"gps", "inspecting"}
	vcs_kw        = []string{"git"}
	// Emulator and player aren't always for games, but those cases should be
	// picked up by one of the other categories first
	game_kw          = []string{"game", "rts", "mmorpg", "emulator", "player"}
	arcadegame_kw    = []string{"combat", "arcade", "racing", "fighting", "fight"}
	actiongame_kw    = []string{"shooter", "fps"}
	adventuregame_kw = []string{"roguelike", "rpg"}
	logicgame_kw     = []string{"puzzle"}
	boardgame_kw     = []string{"board", "chess", "goban", "chessboard"}
	programming_kw   = []string{"code", "ide", "programming", "develop", "compile"}
	system_kw        = []string{"sensor"}
)

// Approximately identify various categories
func GuessCategory(pkgdesc string) string {
	// TODO Use a loop and hash map instead, return without assigning
	categories := "Application"
	if keywordsInDescription(pkgdesc, model3d_kw) {
		categories = "Application;Graphics;3DGraphics"
	} else if keywordsInDescription(pkgdesc, multimedia_kw) {
		categories = "Application;Multimedia"
	} else if keywordsInDescription(pkgdesc, network_kw) {
		categories = "Application;Network"
	} else if keywordsInDescription(pkgdesc, audiovideo_kw) {
		categories = "Application;AudioVideo"
	} else if keywordsInDescription(pkgdesc, office_kw) {
		categories = "Application;Office"
	} else if keywordsInDescription(pkgdesc, editor_kw) {
		categories = "Application;Development;TextEditor"
	} else if keywordsInDescription(pkgdesc, science_kw) {
		categories = "Application;Science"
	} else if keywordsInDescription(pkgdesc, vcs_kw) {
		categories = "Application;Development;RevisionControl"
	} else if keywordsInDescription(pkgdesc, arcadegame_kw) {
		categories = "Application;Game;ArcadeGame"
	} else if keywordsInDescription(pkgdesc, actiongame_kw) {
		categories = "Application;Game;ActionGame"
	} else if keywordsInDescription(pkgdesc, adventuregame_kw) {
		categories = "Application;Game;AdventureGame"
	} else if keywordsInDescription(pkgdesc, logicgame_kw) {
		categories = "Application;Game;"
	} else if keywordsInDescription(pkgdesc, boardgame_kw) {
		categories = "Application;Game;BoardGame"
	} else if keywordsInDescription(pkgdesc, game_kw) {
		categories = "Application;Game"
	} else if keywordsInDescription(pkgdesc, programming_kw) {
		categories = "Application;Development"
	} else if keywordsInDescription(pkgdesc, system_kw) {
		categories = "Application;System"
	}
	return categories
}
