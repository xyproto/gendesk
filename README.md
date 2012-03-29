Desktop File Generator
======================

Generates .desktop and downloads .png icons for a given PKGBUILD

See gendesk --help for more info

Patches are welcome


TODO
----
* Contact upstream about missing icons and desktop files
* Add \_comment to be able to specify a shorter comment (for gpg-crypter, see svn log)
* Add a way to add MimeTypes
* Keywords could have a priority-number, as more specific categories are better (AdventureGame vs just "Application")
* Possibly add support for StartupNotify=true/false and others
* Add more categories and keywords
* Add category "Game;LogicGame" for keyword "puzzle"
* Consider adding \_genericname too
* Consider adding "Type=Application"


Changes from 0.3 to 0.4 (not yet released)
------------------------------------------
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
