#include "qmacpasteboardmime.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMimeData>
#include <QMacPasteboardMime>
#include "_cgo_export.h"

class MyQMacPasteboardMime: public QMacPasteboardMime {
public:
};

char* QMacPasteboardMime_ConvertorName(void* ptr){
	return static_cast<QMacPasteboardMime*>(ptr)->convertorName().toUtf8().data();
}

int QMacPasteboardMime_Count(void* ptr, void* mimeData){
	return static_cast<QMacPasteboardMime*>(ptr)->count(static_cast<QMimeData*>(mimeData));
}

char* QMacPasteboardMime_FlavorFor(void* ptr, char* mime){
	return static_cast<QMacPasteboardMime*>(ptr)->flavorFor(QString(mime)).toUtf8().data();
}

char* QMacPasteboardMime_MimeFor(void* ptr, char* flav){
	return static_cast<QMacPasteboardMime*>(ptr)->mimeFor(QString(flav)).toUtf8().data();
}

void QMacPasteboardMime_DestroyQMacPasteboardMime(void* ptr){
	static_cast<QMacPasteboardMime*>(ptr)->~QMacPasteboardMime();
}

