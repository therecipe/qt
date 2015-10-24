#include "qxmlerrorhandler.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlParseException>
#include <QXmlErrorHandler>
#include "_cgo_export.h"

class MyQXmlErrorHandler: public QXmlErrorHandler {
public:
};

int QXmlErrorHandler_Error(QtObjectPtr ptr, QtObjectPtr exception){
	return static_cast<QXmlErrorHandler*>(ptr)->error(*static_cast<QXmlParseException*>(exception));
}

char* QXmlErrorHandler_ErrorString(QtObjectPtr ptr){
	return static_cast<QXmlErrorHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlErrorHandler_FatalError(QtObjectPtr ptr, QtObjectPtr exception){
	return static_cast<QXmlErrorHandler*>(ptr)->fatalError(*static_cast<QXmlParseException*>(exception));
}

int QXmlErrorHandler_Warning(QtObjectPtr ptr, QtObjectPtr exception){
	return static_cast<QXmlErrorHandler*>(ptr)->warning(*static_cast<QXmlParseException*>(exception));
}

void QXmlErrorHandler_DestroyQXmlErrorHandler(QtObjectPtr ptr){
	static_cast<QXmlErrorHandler*>(ptr)->~QXmlErrorHandler();
}

