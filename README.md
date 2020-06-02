Introduction
------------

[Qt](https://en.wikipedia.org/wiki/Qt_(software)) is a free and open-source widget toolkit for creating graphical user interfaces as well as cross-platform applications that run on various software and hardware platforms with little or no change in the underlying codebase.

[Go](https://en.wikipedia.org/wiki/Go_(programming_language)), also known as Golang, is a programming language designed at Google.

[therecipe/qt](https://github.com/therecipe/qt) allows you to write Qt applications entirely in Go or JavaScript.

Beside the language bindings provided, `therecipe/qt` also greatly simplifies the deployment of Qt applications to various software and hardware platforms.

At the time of writing, almost all Qt functions and classes are accessible from Go and JavaScript, and you should be able to find everything you need to build fully featured Qt applications.

Impressions
-----------

[Gallery](https://github.com/therecipe/qt/wiki/Gallery) of example applications.

[JavaScript Demo](https://therecipe.github.io/entry) | *[source](https://github.com/therecipe/entry)*

Installation
------------

The following instructions assume that you already installed [Go](https://golang.org/dl/) and [Git](https://git-scm.com/downloads)

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

-	[Installation](https://github.com/therecipe/qt/wiki/Installation)
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

Qt itself is licensed and available under multiple [licenses](https://www.qt.io/licensing).
