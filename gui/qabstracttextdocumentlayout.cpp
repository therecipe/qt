#include "qabstracttextdocumentlayout.h"
#include <QPoint>
#include <QPointF>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPaintDevice>
#include <QAbstractTextDocumentLayout>
#include "_cgo_export.h"

class MyQAbstractTextDocumentLayout: public QAbstractTextDocumentLayout {
public:
void Signal_PageCountChanged(int newPages){callbackQAbstractTextDocumentLayoutPageCountChanged(this->objectName().toUtf8().data(), newPages);};
};

char* QAbstractTextDocumentLayout_AnchorAt(void* ptr, void* position){
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->anchorAt(*static_cast<QPointF*>(position)).toUtf8().data();
}

void* QAbstractTextDocumentLayout_Document(void* ptr){
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->document();
}

void* QAbstractTextDocumentLayout_HandlerForObject(void* ptr, int objectType){
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->handlerForObject(objectType);
}

int QAbstractTextDocumentLayout_PageCount(void* ptr){
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->pageCount();
}

void QAbstractTextDocumentLayout_ConnectPageCountChanged(void* ptr){
	QObject::connect(static_cast<QAbstractTextDocumentLayout*>(ptr), static_cast<void (QAbstractTextDocumentLayout::*)(int)>(&QAbstractTextDocumentLayout::pageCountChanged), static_cast<MyQAbstractTextDocumentLayout*>(ptr), static_cast<void (MyQAbstractTextDocumentLayout::*)(int)>(&MyQAbstractTextDocumentLayout::Signal_PageCountChanged));;
}

void QAbstractTextDocumentLayout_DisconnectPageCountChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractTextDocumentLayout*>(ptr), static_cast<void (QAbstractTextDocumentLayout::*)(int)>(&QAbstractTextDocumentLayout::pageCountChanged), static_cast<MyQAbstractTextDocumentLayout*>(ptr), static_cast<void (MyQAbstractTextDocumentLayout::*)(int)>(&MyQAbstractTextDocumentLayout::Signal_PageCountChanged));;
}

void* QAbstractTextDocumentLayout_PaintDevice(void* ptr){
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->paintDevice();
}

void QAbstractTextDocumentLayout_RegisterHandler(void* ptr, int objectType, void* component){
	static_cast<QAbstractTextDocumentLayout*>(ptr)->registerHandler(objectType, static_cast<QObject*>(component));
}

void QAbstractTextDocumentLayout_SetPaintDevice(void* ptr, void* device){
	static_cast<QAbstractTextDocumentLayout*>(ptr)->setPaintDevice(static_cast<QPaintDevice*>(device));
}

void QAbstractTextDocumentLayout_UnregisterHandler(void* ptr, int objectType, void* component){
	static_cast<QAbstractTextDocumentLayout*>(ptr)->unregisterHandler(objectType, static_cast<QObject*>(component));
}

