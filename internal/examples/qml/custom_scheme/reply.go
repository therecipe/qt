package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
)

//the custom QNetworkReply is partially modeled after
//https://code.qt.io/cgit/qt/qtbase.git/tree/src/network/access/qnetworkreplyfileimpl.cpp?h=5.7
//and
//https://blogs.kde.org/2010/08/28/implementing-reusable-custom-qnetworkreply

func NewCustomReply(op network.QNetworkAccessManager__Operation, req *network.QNetworkRequest) *network.QNetworkReply {

	var customReply = network.NewQNetworkReply(nil)
	customReply.SetRequest(req)
	customReply.SetUrl(req.Url())
	customReply.SetOperation(op)
	customReply.SetFinished(true)

	//there will be at least 2 requests
	//one for the qmldir file: https://doc.qt.io/qt-5/qtqml-syntax-directoryimports.html#directory-listing-qmldir-files
	//and another request for each *.qml file that is exported through the qmldir file

	println("requested:", req.Url().ToString(0))

	var content []byte
	if req.Url().FileName(0) == "qmldir" {
		content = []byte("SomeClass myfile.qml")
	} else {
		content = []byte(`import QtQuick 2.0

    Rectangle {
        id: page
        width: 320; height: 480
        color: "lightgreen"

        Text {
            id: helloText
            text: "hello from custom:///"
            y: 30
            anchors.horizontalCenter: page.horizontalCenter
            font.pointSize: 24; font.bold: true
        }
    }`)
	}

	//in case of multiple customReply::ReadData calls for a single reply, an offset is used
	var offset int64
	customReply.ConnectReadData(func(data *string, maxSize int64) int64 {

		if offset >= int64(len(content)) {
			return -1
		}

		var number = func(a, b int64) int64 {
			if a <= b {
				return a
			}
			return b
		}(maxSize, int64(len(content))-offset)

		*data = string(content[offset : offset+number])
		offset += number

		customReply.SetAttribute(network.QNetworkRequest__HttpStatusCodeAttribute, core.NewQVariant7(200))
		customReply.SetAttribute(network.QNetworkRequest__HttpReasonPhraseAttribute, core.NewQVariant14("OK"))

		return number
	})

	customReply.ConnectIsSequential(func() bool {
		return true
	})

	customReply.ConnectBytesAvailable(func() int64 {
		return int64(len(content)) - offset
	})

	customReply.Open(core.QIODevice__ReadOnly | core.QIODevice__Unbuffered)
	customReply.SetHeader(network.QNetworkRequest__ContentLengthHeader, core.NewQVariant9(int64(len(content))))

	customReply.MetaDataChanged()
	customReply.DownloadProgress(int64(len(content)), int64(len(content)))
	customReply.ReadyRead()
	customReply.Finished()

	return customReply
}
