#include "qquickview.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWindow>
#include <QQmlEngine>
#include <QMetaObject>
#include <QObject>
#include <QQuickView>
#include "_cgo_export.h"

class MyQQuickView: public QQuickView {
public:
void Signal_StatusChanged(QQuickView::Status status){callbackQQuickViewStatusChanged(this->objectName().toUtf8().data(), status);};
};

int QQuickView_ResizeMode(QtObjectPtr ptr){
	return static_cast<QQuickView*>(ptr)->resizeMode();
}

void QQuickView_SetResizeMode(QtObjectPtr ptr, int v){
	static_cast<QQuickView*>(ptr)->setResizeMode(static_cast<QQuickView::ResizeMode>(v));
}

int QQuickView_Status(QtObjectPtr ptr){
	return static_cast<QQuickView*>(ptr)->status();
}

QtObjectPtr QQuickView_NewQQuickView2(QtObjectPtr engine, QtObjectPtr parent){
	return new QQuickView(static_cast<QQmlEngine*>(engine), static_cast<QWindow*>(parent));
}

QtObjectPtr QQuickView_NewQQuickView(QtObjectPtr parent){
	return new QQuickView(static_cast<QWindow*>(parent));
}

QtObjectPtr QQuickView_NewQQuickView3(char* source, QtObjectPtr parent){
	return new QQuickView(QUrl(QString(source)), static_cast<QWindow*>(parent));
}

QtObjectPtr QQuickView_Engine(QtObjectPtr ptr){
	return static_cast<QQuickView*>(ptr)->engine();
}

QtObjectPtr QQuickView_RootContext(QtObjectPtr ptr){
	return static_cast<QQuickView*>(ptr)->rootContext();
}

QtObjectPtr QQuickView_RootObject(QtObjectPtr ptr){
	return static_cast<QQuickView*>(ptr)->rootObject();
}

void QQuickView_SetSource(QtObjectPtr ptr, char* url){
	QMetaObject::invokeMethod(static_cast<QQuickView*>(ptr), "setSource", Q_ARG(QUrl, QUrl(QString(url))));
}

char* QQuickView_Source(QtObjectPtr ptr){
	return static_cast<QQuickView*>(ptr)->source().toString().toUtf8().data();
}

void QQuickView_ConnectStatusChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickView*>(ptr), static_cast<void (QQuickView::*)(QQuickView::Status)>(&QQuickView::statusChanged), static_cast<MyQQuickView*>(ptr), static_cast<void (MyQQuickView::*)(QQuickView::Status)>(&MyQQuickView::Signal_StatusChanged));;
}

void QQuickView_DisconnectStatusChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickView*>(ptr), static_cast<void (QQuickView::*)(QQuickView::Status)>(&QQuickView::statusChanged), static_cast<MyQQuickView*>(ptr), static_cast<void (MyQQuickView::*)(QQuickView::Status)>(&MyQQuickView::Signal_StatusChanged));;
}

void QQuickView_DestroyQQuickView(QtObjectPtr ptr){
	static_cast<QQuickView*>(ptr)->~QQuickView();
}

