== Desktop File Generator ==

Generates .desktop and downloads .png icons for a given PKGBUILD

See gendesk --help for more info

Patches are welcome

== TODO ==

* Add category AudioVideo for keyword synthesizer
* Add \_name to be able to specifiy a name (for: zynaddsubfx -> ZynAddSubFX, jedit -> jEdit)
  In the mean time, one can do the following trick, right after the gendesk line:
    setconf "$pkgname.desktop" Name ZynAddSubFX
    or
    setconf "$pkgname.desktop" Name jEdit
* Add \_comment to be able to specify a shorter comment (for gpg-crypter, see svn log)
* Contact upstream about missing icons and desktop files
* New flag: -q for quiet
  Workaround is: > /dev/null
* New flag: -n for no color
* New flag: -d for generate .desktop file only
  Workaround is: touch \_.png; gendesk; rm \_.png to avoid downloading icon
* Add category TextEditor and/or Development;TextEditor for text editors
* kw "emulator" should be category "Game"
* kw "game" should be category "Game"
* kw "combat" could be category "Game;ArcadeGame"
* kw "GPS" or "inspecting" could maybe be category "Application;Science"
* kw "single player", "single-player", "multi player" shooter "multi-player" should be category "Application;Game;"
* kw "shooter" should be "Application;Game;ActionGame;"
* kw "roguelike" should be "Application;Game;AdventureGame;"
* keywords should have a priority-number, as more specific categories are better (AdventureGame vs just "Application")
* Add a way to add MimeTypes
* kw "git" should be category Development;RevisionControl
* Possibly add support for StartupNotify, Type=Application and 
