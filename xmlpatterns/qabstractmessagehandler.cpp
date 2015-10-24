#include "qabstractmessagehandler.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractMessageHandler>
#include "_cgo_export.h"

class MyQAbstractMessageHandler: public QAbstractMessageHandler {
public:
};

void QAbstractMessageHandler_DestroyQAbstractMessageHandler(QtObjectPtr ptr){
	static_cast<QAbstractMessageHandler*>(ptr)->~QAbstractMessageHandler();
}

