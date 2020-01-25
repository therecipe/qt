Introduction
------------

[Qt](https://en.wikipedia.org/wiki/Qt_(software)) is a cross-platform application framework that is used for developing application software that can be run on various software and hardware platforms with little or no change in the underlying codebase.

[Go](https://en.wikipedia.org/wiki/Go_(programming_language)) (often referred to as golang) is a programming language created at Google.

This package allows you to write Qt applications entirely in Go and makes deploying them later very easy.

[Gallery](https://github.com/therecipe/qt/wiki/Gallery) of example applications making use of this package.

[WebAssembly Demo (new)](https://therecipe.github.io/entry) | *[repo](https://github.com/therecipe/entry)*

[WebAssembly Demo (old)](https://therecipe.github.io/widgets_playground) | *[repo](https://github.com/therecipe/widgets_playground)*

Status
------

Almost all Qt functions and classes are accessible from Go and you should be able to find everything you need to build fully featured applications.

Installation
------------

##### Windows [(more info)](https://github.com/therecipe/qt/wiki/Installation-on-Windows)

```powershell
set GO111MODULE=off
go get -v github.com/therecipe/qt/cmd/... && for /f %v in ('go env GOPATH') do %v\bin\qtsetup test && %v\bin\qtsetup -test=false
```

##### macOS [(more info)](https://github.com/therecipe/qt/wiki/Installation-on-macOS)

```bash
export GO111MODULE=off; xcode-select --install; go get -v github.com/therecipe/qt/cmd/... && $(go env GOPATH)/bin/qtsetup test && $(go env GOPATH)/bin/qtsetup -test=false
```

##### Linux [(more info)](https://github.com/therecipe/qt/wiki/Installation-on-Linux)

```bash
export GO111MODULE=off; go get -v github.com/therecipe/qt/cmd/... && $(go env GOPATH)/bin/qtsetup test && $(go env GOPATH)/bin/qtsetup -test=false
```

Resources
---------

-	[General Installation](https://github.com/therecipe/qt/wiki/Installation)
-	[Getting Started](https://github.com/therecipe/qt/wiki/Getting-Started)
-	[Wiki](https://github.com/therecipe/qt/wiki)
-	[Qt Documentation](https://doc.qt.io/qt-5/classes.html)
-	[FAQ](https://github.com/therecipe/qt/wiki/FAQ)
-	[#qt-binding](https://gophers.slack.com/messages/qt-binding/details) Slack channel ([invite](https://invite.slack.golangbridge.org)\)

Deployment Targets
------------------

| Target                   | Arch             | Linkage                   | Docker Deployment | Host OS |
|:------------------------:|:----------------:|:-------------------------:|:-----------------:|:-------:|
|         Windows          |     32 / 64      |     dynamic / static      |        Yes        |   Any   |
|          macOS           |        64        |          dynamic          |        Yes        |   Any   |
|          Linux           | arm / arm64 / 64 | dynamic / static / system |        Yes        |   Any   |
|     Android (+Wear)      |   arm / arm64    |          dynamic          |        Yes        |   Any   |
| Android-Emulator (+Wear) |        32        |          dynamic          |        Yes        |   Any   |
|        SailfishOS        |       arm        |          system           |        Yes        |   Any   |
|   SailfishOS-Emulator    |        32        |          system           |        Yes        |   Any   |
|   Raspberry Pi (1/2/3)   |       arm        |     dynamic / system      |        Yes        |   Any   |
|       Ubuntu Touch       |     arm / 64     |          system           |        Yes        |   Any   |
|        JavaScript        |        32        |          static           |        Yes        |   Any   |
|       WebAssembly        |        32        |          static           |        Yes        |   Any   |
|           iOS            |      arm64       |          static           |        No         |  macOS  |
|      iOS-Simulator       |        64        |          static           |        No         |  macOS  |
|        AsteroidOS        |       arm        |          system           |        No         |  Linux  |
|         FreeBSD          |     32 / 64      |          system           |        No         | FreeBSD |

License
-------

This package is released under [LGPLv3](https://opensource.org/licenses/LGPL-3.0)

Qt is available under multiple [licenses](https://www.qt.io/licensing)
