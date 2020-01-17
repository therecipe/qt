package cpp

import (
	"unsafe"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
)

type CacheReply struct {
	network.QNetworkReply

	_ func() `constructor:"init"`

	mCacheDev *core.QIODevice
}

func (cr *CacheReply) init() {
	cr.ConnectAbort(cr.abort)
	cr.ConnectBytesAvailable(cr.bytesAvailable)
	cr.ConnectIsSequential(cr.isSequential)
	cr.ConnectSize(cr.size)
	cr.ConnectReadData(cr.readData)
	cr.ConnectDestroyCacheReply(cr.destroyCacheReply)
}

func (cr *CacheReply) With(
	cacheDev *core.QIODevice,
	req *network.QNetworkRequest,
	op network.QNetworkAccessManager__Operation,
	meta *network.QNetworkCacheMetaData) *CacheReply {

	cr.mCacheDev = cacheDev

	cr.SetRequest(req)
	cr.SetUrl(req.Url())
	cr.SetOperation(op)
	cr.SetAttribute(network.QNetworkRequest__HttpStatusCodeAttribute, core.NewQVariant1(200))
	cr.SetFinished(true)
	cr.OpenDefault(core.QIODevice__ReadOnly)

	size := cacheDev.Size()
	for _, header := range meta.RawHeaders() {
		cr.SetRawHeader(header.First(), header.Second())
	}
	cr.MetaDataChanged()
	cr.DownloadProgress(size, size)
	cr.ReadyRead()
	cr.Finished()

	return cr
}

func (cr *CacheReply) abort() {
	cr.Close()
}

func (cr *CacheReply) bytesAvailable() int64 {
	return cr.mCacheDev.BytesAvailable()
}

func (cr *CacheReply) isSequential() bool {
	return cr.mCacheDev.IsSequential()
}

func (cr *CacheReply) size() int64 {
	return cr.mCacheDev.Size()
}

func (cr *CacheReply) readData(data *string, maxlen int64) int64 {
	return cr.mCacheDev.Read(*(*[]byte)(unsafe.Pointer(data)), maxlen)
}

func (cr *CacheReply) destroyCacheReply() {
	cr.Close()
	if cr.mCacheDev != nil {
		cr.mCacheDev.Close()
		cr.mCacheDev.DestroyQIODevice()
	}
}
