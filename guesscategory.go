package main

import (
	"errors"
	"strings"
)

// TODO: Use an external file to read the mappings from (possibly JSON)

const (
	// Decides the order of the keyword/category checks
	// (try to order from the more specific/specialized categories to the more general)
	tracker = iota
	texttools
	graphics2d
	scanning
	utility
	settings
	hardwaresettings
	audio
	video
	education
	math
	cs
	compression
	filetools
	model3d
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
		tracker:       {"sequencer"},
		model3d:       {"rendering", "modeling", "modelling", "modeler", "render", "raytracing", "CAD"},
		multimedia:    {"non-linear", "audio", "sound", "graphics", "demo", "music"},
		graphics:      {"draw", "pixelart", "animated"},
		network:       {"network", "p2p", "browser", "remote"},
		email:         {"gmail", "email", "e-mail", "mail"},
		audiovideo:    {"synth", "synthesizer", "ffmpeg", "guitar"},
		office:        {"ebook", "e-book", "spreadsheet", "calculator", "processor", "documents"},
		editor:        {"editor"},
		science:       {"gps", "inspecting", "molecular", "mathematics"},
		vcs:           {"git"},
		arcadegame:    {"combat", "arcade", "racing", "fighting", "fight", "shooter"},
		actiongame:    {"shooter", "fps"},
		adventuregame: {"roguelike", "rpg"},
		logicgame:     {"puzzle"},
		boardgame:     {"board", "chess", "goban", "chessboard", "checkers", "reversi", "go"},
		// "emulator" and "player" aren't always for games, but those cases will be
		// picked up by one of the other categories first, as orderd by the constants above
		game:        {"game", "rts", "mmorpg", "emulator", "player"},
		programming: {"code", "ide", "programming", "develop", "compile", "interpret", "valgrind"},
		system:      {"sensor", "bus", "calibration", "usb", "file"},
	}
	categorymap = map[int]string{
		tracker:          "Application;Multimedia;Audio;Sequencer;Music",
		model3d:          "Application;Graphics;3DGraphics",
		multimedia:       "Application;Multimedia",
		graphics:         "Application;Graphics",
		network:          "Application;Network",
		email:            "Application;Network;Email",
		audiovideo:       "Application;AudioVideo",
		office:           "Application;Office",
		editor:           "Application;Development;TextEditor",
		science:          "Application;Science",
		vcs:              "Application;Development;RevisionControl",
		arcadegame:       "Application;Game;ArcadeGame",
		actiongame:       "Application;Game;ActionGame",
		adventuregame:    "Application;Game;AdventureGame",
		logicgame:        "Application;Game",
		boardgame:        "Application;Game;BoardGame",
		game:             "Application;Game",
		programming:      "Application;Development",
		system:           "Application;System",
		texttools:        "Application;TextTools",
		graphics2d:       "Application;Graphics;2DGraphics",
		scanning:         "Application;Grahpics;Scanning",
		utility:          "Application;Utility",
		settings:         "Application;Settings",
		hardwaresettings: "Application;HardwareSettings;Settings",
		audio:            "Application;AudioVideo;Audio",
		video:            "Application;AudioVideo;Video",
		education:        "Application;Science",
		math:             "Application;Science;Math",
		cs:               "Application;Science;ComputerScience",
		compression:      "Application;Utility;Archiving",
		filetools:        "Application;System;FileTools",
	}
)

// ValidCategoryWords validates each word in 'categoryWords' against a list of accepted categories derived from 'categorymap'.
// It ensures all provided words represent valid application categories, returning an error for any unrecognized category.
func ValidCategoryWords(categoryWords []string) error {
	var validWords []string
	for _, v := range categorymap {
		fields := strings.Split(v, ";")
		for _, field := range fields {
			if strings.TrimSpace(field) == "" {
				continue
			}
			if !hasS(validWords, field) {
				validWords = append(validWords, field)
			}
		}
	}
	for _, word := range categoryWords {
		if !hasS(validWords, word) {
			return errors.New(word + " is an unrecognized category")
		}
	}
	return nil
}

// GuessCategory will try to guess which category an application belongs to,
// given a short package description.
// If no guess is made, it will return "Application".
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
