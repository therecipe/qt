#include "qabstracttextdocumentlayout.h"
#include <QUrl>
#include <QModelIndex>
#include <QPaintDevice>
#include <QPointF>
#include <QObject>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QAbstractTextDocumentLayout>
#include "_cgo_export.h"

class MyQAbstractTextDocumentLayout: public QAbstractTextDocumentLayout {
public:
void Signal_PageCountChanged(int newPages){callbackQAbstractTextDocumentLayoutPageCountChanged(this->objectName().toUtf8().data(), newPages);};
};

char* QAbstractTextDocumentLayout_AnchorAt(QtObjectPtr ptr, QtObjectPtr position){
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->anchorAt(*static_cast<QPointF*>(position)).toUtf8().data();
}

QtObjectPtr QAbstractTextDocumentLayout_Document(QtObjectPtr ptr){
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->document();
}

QtObjectPtr QAbstractTextDocumentLayout_HandlerForObject(QtObjectPtr ptr, int objectType){
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->handlerForObject(objectType);
}

int QAbstractTextDocumentLayout_PageCount(QtObjectPtr ptr){
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->pageCount();
}

void QAbstractTextDocumentLayout_ConnectPageCountChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractTextDocumentLayout*>(ptr), static_cast<void (QAbstractTextDocumentLayout::*)(int)>(&QAbstractTextDocumentLayout::pageCountChanged), static_cast<MyQAbstractTextDocumentLayout*>(ptr), static_cast<void (MyQAbstractTextDocumentLayout::*)(int)>(&MyQAbstractTextDocumentLayout::Signal_PageCountChanged));;
}

void QAbstractTextDocumentLayout_DisconnectPageCountChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractTextDocumentLayout*>(ptr), static_cast<void (QAbstractTextDocumentLayout::*)(int)>(&QAbstractTextDocumentLayout::pageCountChanged), static_cast<MyQAbstractTextDocumentLayout*>(ptr), static_cast<void (MyQAbstractTextDocumentLayout::*)(int)>(&MyQAbstractTextDocumentLayout::Signal_PageCountChanged));;
}

QtObjectPtr QAbstractTextDocumentLayout_PaintDevice(QtObjectPtr ptr){
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->paintDevice();
}

void QAbstractTextDocumentLayout_RegisterHandler(QtObjectPtr ptr, int objectType, QtObjectPtr component){
	static_cast<QAbstractTextDocumentLayout*>(ptr)->registerHandler(objectType, static_cast<QObject*>(component));
}

void QAbstractTextDocumentLayout_SetPaintDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QAbstractTextDocumentLayout*>(ptr)->setPaintDevice(static_cast<QPaintDevice*>(device));
}

void QAbstractTextDocumentLayout_UnregisterHandler(QtObjectPtr ptr, int objectType, QtObjectPtr component){
	static_cast<QAbstractTextDocumentLayout*>(ptr)->unregisterHandler(objectType, static_cast<QObject*>(component));
}

