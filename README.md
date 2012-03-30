Desktop File Generator
======================

Generates .desktop and downloads .png icons for a given PKGBUILD

See gendesk --help for more info

Patches are welcome


TODO
----
* Contact upstream about missing icons and desktop files
* Keywords could have a priority-number, as more specific categories are better (AdventureGame vs just "Application")
* Possibly add support for StartupNotify=true/false and Terminal=true/false
* Add even more categories and keywords
* Add a way to add a custom lines at the end, perhaps with \_custom=()
* Put kw/category mappings in a single map
* Move funcions and settings related to terminal output to a separate file (struct + methods)
* Check if gendesk can be installed with go install
* Glob for .svg icons too
* Test, test, test, then release 0.4
* Just maybe write a makepkg patch to lower the threshold for using gendesk


Changes from 0.3 to 0.4 (not released yet)
----------------------------------
* Added \_genericname=()
* Added \_comment=()
* Added \_mimetype=()
* Added Type=Application
* Added category "Game;LogicGame" for keyword "puzzle"
* Added category "Game;ArcadeGame" for keyword "fighting"
* Fixed weird formatting in --help output


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

