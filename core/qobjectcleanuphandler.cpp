#include "qobjectcleanuphandler.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObjectCleanupHandler>
#include "_cgo_export.h"

class MyQObjectCleanupHandler: public QObjectCleanupHandler {
public:
};

void* QObjectCleanupHandler_NewQObjectCleanupHandler(){
	return new QObjectCleanupHandler();
}

void* QObjectCleanupHandler_Add(void* ptr, void* object){
	return static_cast<QObjectCleanupHandler*>(ptr)->add(static_cast<QObject*>(object));
}

void QObjectCleanupHandler_Clear(void* ptr){
	static_cast<QObjectCleanupHandler*>(ptr)->clear();
}

int QObjectCleanupHandler_IsEmpty(void* ptr){
	return static_cast<QObjectCleanupHandler*>(ptr)->isEmpty();
}

void QObjectCleanupHandler_DestroyQObjectCleanupHandler(void* ptr){
	static_cast<QObjectCleanupHandler*>(ptr)->~QObjectCleanupHandler();
}

