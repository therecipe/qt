#include "qquickview.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWindow>
#include <QMetaObject>
#include <QQmlEngine>
#include <QString>
#include <QQuickView>
#include "_cgo_export.h"

class MyQQuickView: public QQuickView {
public:
void Signal_StatusChanged(QQuickView::Status status){callbackQQuickViewStatusChanged(this->objectName().toUtf8().data(), status);};
};

int QQuickView_ResizeMode(void* ptr){
	return static_cast<QQuickView*>(ptr)->resizeMode();
}

void QQuickView_SetResizeMode(void* ptr, int v){
	static_cast<QQuickView*>(ptr)->setResizeMode(static_cast<QQuickView::ResizeMode>(v));
}

int QQuickView_Status(void* ptr){
	return static_cast<QQuickView*>(ptr)->status();
}

void* QQuickView_NewQQuickView2(void* engine, void* parent){
	return new QQuickView(static_cast<QQmlEngine*>(engine), static_cast<QWindow*>(parent));
}

void* QQuickView_NewQQuickView(void* parent){
	return new QQuickView(static_cast<QWindow*>(parent));
}

void* QQuickView_NewQQuickView3(void* source, void* parent){
	return new QQuickView(*static_cast<QUrl*>(source), static_cast<QWindow*>(parent));
}

void* QQuickView_Engine(void* ptr){
	return static_cast<QQuickView*>(ptr)->engine();
}

void* QQuickView_RootContext(void* ptr){
	return static_cast<QQuickView*>(ptr)->rootContext();
}

void* QQuickView_RootObject(void* ptr){
	return static_cast<QQuickView*>(ptr)->rootObject();
}

void QQuickView_SetSource(void* ptr, void* url){
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setSource", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQuickView_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QQuickView*>(ptr), static_cast<void (QQuickView::*)(QQuickView::Status)>(&QQuickView::statusChanged), static_cast<MyQQuickView*>(ptr), static_cast<void (MyQQuickView::*)(QQuickView::Status)>(&MyQQuickView::Signal_StatusChanged));;
}

void QQuickView_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickView*>(ptr), static_cast<void (QQuickView::*)(QQuickView::Status)>(&QQuickView::statusChanged), static_cast<MyQQuickView*>(ptr), static_cast<void (MyQQuickView::*)(QQuickView::Status)>(&MyQQuickView::Signal_StatusChanged));;
}

void QQuickView_DestroyQQuickView(void* ptr){
	static_cast<QQuickView*>(ptr)->~QQuickView();
}

