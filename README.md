Desktop File Generator
======================

Generates .desktop files and downloads .png icons for a given PKGBUILD.

See gendesk --help for more info.

Patches and pull requests are welcome.


TODO
----
* Contact upstream about missing icons and desktop files
* Add even more categories and keywords
* Check if gendesk can be installed with go install
* Possibly write a makepkg patch to lower the threshold for using gendesk
* Possibly put kw/category mappings in a single map instead
* Possibly add support for StartupNotify=true/false and Terminal=true/false


Changes from 0.4 to 0.4.1 (released)
------------------------------------
* Fixed a bug where \_name=() and \_comment=() didn't work as they should


Changes from 0.3 to 0.4 (released)
----------------------------------
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


Changes from 0.2 to 0.3 (released)
----------------------------------
* New flag: -q for quiet
* New flag: --nocolor for no color
* New flag: -n for not downloading anything (only generate a .desktop file)
* New flag: -q for quiet (no stdout output)
* Added \_name=('Name') to be able to specifiy a name that isn't only lowercase (like "ZynAddSubFX" or "jEdit")
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


