Desktop File Generator
======================

Generates .desktop and downloads .png icons for a given PKGBUILD

See gendesk --help for more info

Patches are welcome

TODO
----

* Add \_comment to be able to specify a shorter comment (for gpg-crypter, see svn log)
* Contact upstream about missing icons and desktop files
* keywords should have a priority-number, as more specific categories are better (AdventureGame vs just "Application")
* Add a way to add MimeTypes
* Possibly add support for StartupNotify, Type=Application and 

* kw "synthesizer" should be category AudioVideo
* kw "editor" should be category TextEditor and/or Development;TextEditor
* kw "emulator" should be category "Game"
* kw "game" should be category "Game"
* kw "combat" could be category "Game;ArcadeGame"
* kw "GPS" or "inspecting" could maybe be category "Application;Science"
* kw "single player", "single-player", "multi player" shooter "multi-player" should be category "Application;Game;"
* kw "shooter" should be "Application;Game;ActionGame;"
* kw "roguelike" should be "Application;Game;AdventureGame;"
* kw "git" should be category Development;RevisionControl


Changes from 0.2 to 0.3
-----------------------
* New flag: -q for quiet
* New flag: --nocolor for no color
* New flag: -n for not downloading anything (only generate a .desktop file)
* New flag: -q for quiet (no stdout output)
* Use \_name=('Name') to be able to specifiy a name that isn't only lowercase (like "ZynAddSubFX" or "jEdit")


