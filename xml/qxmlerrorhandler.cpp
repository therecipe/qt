#include "qxmlerrorhandler.h"
#include <QUrl>
#include <QModelIndex>
#include <QXmlParseException>
#include <QString>
#include <QVariant>
#include <QXmlErrorHandler>
#include "_cgo_export.h"

class MyQXmlErrorHandler: public QXmlErrorHandler {
public:
};

int QXmlErrorHandler_Error(void* ptr, void* exception){
	return static_cast<QXmlErrorHandler*>(ptr)->error(*static_cast<QXmlParseException*>(exception));
}

char* QXmlErrorHandler_ErrorString(void* ptr){
	return static_cast<QXmlErrorHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlErrorHandler_FatalError(void* ptr, void* exception){
	return static_cast<QXmlErrorHandler*>(ptr)->fatalError(*static_cast<QXmlParseException*>(exception));
}

int QXmlErrorHandler_Warning(void* ptr, void* exception){
	return static_cast<QXmlErrorHandler*>(ptr)->warning(*static_cast<QXmlParseException*>(exception));
}

void QXmlErrorHandler_DestroyQXmlErrorHandler(void* ptr){
	static_cast<QXmlErrorHandler*>(ptr)->~QXmlErrorHandler();
}

