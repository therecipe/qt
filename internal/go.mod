module github.com/therecipe/qt/internal

require (
	github.com/gopherjs/gopherjs v0.0.0-20190309154008-847fc94819f9
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/sirupsen/logrus v1.4.0
	github.com/stretchr/testify v1.3.0
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/net v0.0.0-20190326090315-15845e8f865b // indirect
	golang.org/x/sys v0.0.0-20190322080309-f49334f85ddc // indirect
	golang.org/x/tools v0.0.0-20190326221639-3ad05305c9b0
)

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c

replace golang.org/x/net => github.com/golang/net v0.0.0-20190326090315-15845e8f865b

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190322080309-f49334f85ddc

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20190326221639-3ad05305c9b0
