#include "qobjectcleanuphandler.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QObjectCleanupHandler>
#include "_cgo_export.h"

class MyQObjectCleanupHandler: public QObjectCleanupHandler {
public:
};

QtObjectPtr QObjectCleanupHandler_NewQObjectCleanupHandler(){
	return new QObjectCleanupHandler();
}

QtObjectPtr QObjectCleanupHandler_Add(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QObjectCleanupHandler*>(ptr)->add(static_cast<QObject*>(object));
}

void QObjectCleanupHandler_Clear(QtObjectPtr ptr){
	static_cast<QObjectCleanupHandler*>(ptr)->clear();
}

int QObjectCleanupHandler_IsEmpty(QtObjectPtr ptr){
	return static_cast<QObjectCleanupHandler*>(ptr)->isEmpty();
}

void QObjectCleanupHandler_DestroyQObjectCleanupHandler(QtObjectPtr ptr){
	static_cast<QObjectCleanupHandler*>(ptr)->~QObjectCleanupHandler();
}

