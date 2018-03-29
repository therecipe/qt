Introduction
------------

[Qt](https://en.wikipedia.org/wiki/Qt_(software)) is a cross-platform application framework that is used for developing application software that can be run on various software and hardware platforms with little or no change in the underlying codebase.

[Go](https://en.wikipedia.org/wiki/Go_(programming_language)) (often referred to as golang) is a programming language created at Google.

This package allows you to write Qt applications entirely in Go.

[Gallery](https://github.com/therecipe/qt/wiki/Gallery) of applications making use of this package.

Status
------

**WIP**

Most Qt functions are accessible from Go.

The package should already contain everything you need to build fully featured applications, the webengine/webview packages don't work on Windows though.

Please pin the repo to a commit that is known to work for you, because there have been no releases so far.

Resources
---------

-	[Wiki](https://github.com/therecipe/qt/wiki)
-	[FAQ](https://github.com/therecipe/qt/wiki/FAQ)
-	[Installation](https://github.com/therecipe/qt/wiki/Installation)
-	[Getting Started](https://github.com/therecipe/qt/wiki/Getting-Started)
-	[Qt Documentation](https://doc.qt.io/qt-5/classes.html)
-	[#qt-binding](https://gophers.slack.com/messages/qt-binding/details) Slack channel ([invite](https://invite.slack.golangbridge.org)\)

Deployment Targets
------------------

| Target                   | Arch       | Linkage                 | Docker Deployment | Host OS |
|:------------------------:|:----------:|:-----------------------:|:-----------------:|:-------:|
|         Windows          | (32 / 64)  |   (dynamic / static)    |        Yes        |   Any   |
|          Linux           |     64     | (dynamic / system libs) |        Yes        |   Any   |
|     Android (+Wear)      |    arm     |         dynamic         |        Yes        |   Any   |
| Android-Emulator (+Wear) |     32     |         dynamic         |        Yes        |   Any   |
|   Raspberry Pi (1/2/3)   |    arm     | (dynamic / system libs) |        Yes        |   Any   |
|        SailfishOS        |    arm     |       system libs       |        Yes        |   Any   |
|   SailfishOS-Emulator    |     32     |       system libs       |        Yes        |   Any   |
|       Ubuntu Touch       | (arm / 64) |       system libs       |        Yes        |   Any   |
|          macOS           |     64     |         dynamic         |        No         |  macOS  |
|           iOS            |   arm64    |         static          |        No         |  macOS  |
|      iOS-Simulator       |     64     |         static          |        No         |  macOS  |
|        AsteroidOS        |    arm     |       system libs       |        No         |  Linux  |

License
-------

This package is released under [LGPLv3](https://opensource.org/licenses/LGPL-3.0).

Qt is available under multiple [licenses](https://www.qt.io/licensing).
