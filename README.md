Desktop File Generator
======================

[![Build Status](https://travis-ci.org/xyproto/gendesk.svg?branch=master)](https://travis-ci.org/xyproto/gendesk) [![Report Card](https://img.shields.io/badge/go_report-A+-brightgreen.svg?style=flat)](http://goreportcard.com/report/xyproto/gendesk)

Generates .desktop files and downloads .png icons based on command line arguments.

See `gendesk --help` or the man page for more info.

Pull requests are welcome.

TODO
----
* Move kw/category mappings into a separate configuration file.

Changes from 0.7.0 to 1.0.0
---------------------------
* Add `--icon` flag, ref #7.
* Update to the desktop-entry-spec 1.2 format (remove `Encoding` and specify `Version`), ref #8.
* Several minor changes, as suggested by the `golint` utility.
* Tested with Go 1.11.

Changes from 0.6.5 to 0.7.0
---------------------------
* Updated vendored dependencies.
* Added support for [goreleaser](https://github.com/goreleaser/goreleaser).
* Improved handling of icons, if an icon is missing.
* Minor changes and refactoring.

Changes from 0.6.4 to 0.6.5
---------------------------
* Ignore the `-svn` suffix in package names (same as for `-git`, thanks @mstraube).
* Use `text/template` for generating the `.desktop` file contents.
* Minor changes to the command line output/documentation.
* Some refactoring.
* Tested with Go 1.9.

Changes from 0.6.3 to 0.6.4
---------------------------
* Fix bug where some flags could not be overridden.

Changes from 0.6.2 to 0.6.3
---------------------------
* Will now ignore the `-git` suffix if it is part of a package name.

Changes from 0.6.1 to 0.6.2
---------------------------
* Added the possibility of having a configuration file for specifying a different URL for searching for missing icons.
* Remove the `--iconurl` flag.
* Refactored out some code to an external package.

Changes from 0.6 to 0.6.1
-------------------------
* Support for `StartupNotify=true`/`false`
* Both `--mimetype` and `--mimetypes` are allowed
* Guesses more categories than before


Changes from 0.5.5 to 0.6
-------------------------
* Added an option for generating .desktop files for launching window managers


Changes from 0.5.4 to 0.5.5
---------------------------
* Bug fix when generating .desktop files from PKGBUILD files.


Changes from 0.5.3 to 0.5.4
---------------------------
* Added a `-f` flag for overwriting files (will not overwrite without it).
* Some refactoring


Changes from 0.5.2 to 0.5.3
---------------------------
* Added a `--terminal` flag for specifying if the application should be run in a terminal.
* Some refactoring.


Changes from 0.5.1 to 0.5.2
---------------------------
* Support for additional environment variables.


Changes from 0.5.0 to 0.5.1
---------------------------
* Support for `$pkgname` and `$pkgdesc`.
* Updated the man page.
* Will try to download icons specified with `--iconurl`.


Changes from 0.4.4 to 0.5.0
---------------------------
* Command line options, no need to specify a PKGBUILD.


Changes from 0.4.3 to 0.4.4
---------------------------
* Changed the URL for searching for icons from Fedora to Open Icon Library


Changes from 0.4.2 to 0.4.3
---------------------------
* Fixed minor bug where puzzle games were not placed in the right category
* Added \_categories=()


Changes from 0.4.1 to 0.4.2
---------------------------
* Added category "Graphics;3DGraphics;" for 3D modellers
* Added category "System;" for sensor monitors
* Added category "Game;BoardGame;" for kw "board", "chess", "goban" or "chessboard"
* Added category "Office" for kw "e-book" and "ebook"
* Doesn't use ".png" by default when specifying an icon


Changes from 0.4 to 0.4.1
-------------------------
* Fixed a bug where \_name=() and \_comment=() didn't work as they should


Changes from 0.3 to 0.4
-----------------------
* Added \_genericname=()
* Added \_comment=()
* Added \_mimetype=()
* Added Type=Application
* Added category "Game;LogicGame" for keyword "puzzle"
* Added category "Game;ArcadeGame" for keyword "fighting"
* Fixed weird formatting in --help output
* Added \_custom=() for adding custom fields at the end of the .desktop file
* Glob for existing .svg icons too
* Shorter lines
* Moved functions and settings related to terminal output to a separate file


Changes from 0.2 to 0.3
-----------------------
* New flag: -q for quiet
* New flag: --nocolor for no color
* New flag: -n for not downloading anything (only generate a .desktop file)
* New flag: -q for quiet (no stdout output)
* Added \_name=('Name') to be able to specify a name that isn't only lowercase (like "ZynAddSubFX" or "jEdit")
* kw "synthesizer" is now category AudioVideo
* kw "editor" is now category TextEditor and/or Development;TextEditor
* kw "emulator" is now category "Game"
* kw "game" is now category "Game"
* kw "combat" is now be category "Game;ArcadeGame"
* kw "GPS" or "inspecting" is now category "Application;Science"
* kw "player" is now category "Application;Game;"
* kw "shooter" is now "Application;Game;ActionGame;"
* kw "roguelike" is now "Application;Game;AdventureGame;"
* kw "git" is now category Development;RevisionControl


General information
-------------------

* Version: 1.0.0
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
* License: GPL2
