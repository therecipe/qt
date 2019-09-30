package cpp

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
)

type CachingNetworkAccessManager struct {
	network.QNetworkAccessManager

	_ func() `constructor:"init"`

	mUrlIgnoreList []string
}

func (cnam *CachingNetworkAccessManager) init() {
	cnam.mUrlIgnoreList = []string{
		"https://www.qtworldsummit.com/api/schedule/all/",
		"https://www.qtworldsummit.com/api/speakers/all/",
		"https://www.qtworldsummit.com/api/version/show/",
		"https://felgo.com/qml-sources/qtws2017/QtWSVersionCheck-test.qml",
		"https://felgo.com/qml-sources/qtws2017/QtWSVersionCheck.qml",
		"https://felgo.com/qml-sources/qmldir",

		"https://www.qtworldsummit.com/api/schedule/all/",
		"https://www.qtworldsummit.com/api/speakers/all/",
		"https://www.qtworldsummit.com/api/version/show/",
		"https://felgo.com/qml-sources/qtws2017/QtWSVersionCheck-test.qml",
		"https://felgo.com/qml-sources/qtws2017/QtWSVersionCheck.qml",
		"https://felgo.com/qml-sources/qmldir",
	}

	cnam.ConnectCreateRequest(cnam.createRequest)
}

func (cnam *CachingNetworkAccessManager) createRequest(
	op network.QNetworkAccessManager__Operation,
	req *network.QNetworkRequest,
	outgoingData *core.QIODevice) *network.QNetworkReply {

	meta := cnam.Cache().MetaData(req.Url())
	if meta.IsValid() && !cnam.shouldIgnoreUrl(req.Url().Url(0)) {
		//cache contains URL -> return cache reply
		//TODO need to check for expiration date?
		return NewCacheReply(cnam).With(cnam.Cache().Data(req.Url()), req, op, meta).QNetworkReply_PTR()
	}
	return cnam.CreateRequestDefault(op, req, outgoingData)
}

func (cnam *CachingNetworkAccessManager) shouldIgnoreUrl(url string) bool {
	for _, s := range cnam.mUrlIgnoreList {
		if s == url {
			return true
		}
	}
	return false
}

func (cnam *CachingNetworkAccessManager) clearIgnoredUrlsFromCache() {
	for _, s := range cnam.mUrlIgnoreList {
		cnam.Cache().Remove(core.NewQUrl3(s, 0))
	}
}
