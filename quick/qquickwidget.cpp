#include "qquickwidget.h"
#include <QString>
#include <QVariant>
#include <QQmlEngine>
#include <QMetaObject>
#include <QWidget>
#include <QObject>
#include <QSurfaceFormat>
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include <QQuickWindow>
#include <QSurface>
#include <QQuickWidget>
#include "_cgo_export.h"

class MyQQuickWidget: public QQuickWidget {
public:
void Signal_SceneGraphError(QQuickWindow::SceneGraphError error, const QString & message){callbackQQuickWidgetSceneGraphError(this->objectName().toUtf8().data(), error, message.toUtf8().data());};
void Signal_StatusChanged(QQuickWidget::Status status){callbackQQuickWidgetStatusChanged(this->objectName().toUtf8().data(), status);};
};

int QQuickWidget_ResizeMode(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->resizeMode();
}

void QQuickWidget_SetResizeMode(void* ptr, int v){
	static_cast<QQuickWidget*>(ptr)->setResizeMode(static_cast<QQuickWidget::ResizeMode>(v));
}

int QQuickWidget_Status(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->status();
}

void* QQuickWidget_NewQQuickWidget2(void* engine, void* parent){
	return new QQuickWidget(static_cast<QQmlEngine*>(engine), static_cast<QWidget*>(parent));
}

void* QQuickWidget_NewQQuickWidget(void* parent){
	return new QQuickWidget(static_cast<QWidget*>(parent));
}

void* QQuickWidget_NewQQuickWidget3(void* source, void* parent){
	return new QQuickWidget(*static_cast<QUrl*>(source), static_cast<QWidget*>(parent));
}

void* QQuickWidget_Engine(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->engine();
}

void* QQuickWidget_QuickWindow(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->quickWindow();
}

void* QQuickWidget_RootContext(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->rootContext();
}

void* QQuickWidget_RootObject(void* ptr){
	return static_cast<QQuickWidget*>(ptr)->rootObject();
}

void QQuickWidget_ConnectSceneGraphError(void* ptr){
	QObject::connect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWidget::sceneGraphError), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWidget::Signal_SceneGraphError));;
}

void QQuickWidget_DisconnectSceneGraphError(void* ptr){
	QObject::disconnect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&QQuickWidget::sceneGraphError), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWindow::SceneGraphError, const QString &)>(&MyQQuickWidget::Signal_SceneGraphError));;
}

void QQuickWidget_SetClearColor(void* ptr, void* color){
	static_cast<QQuickWidget*>(ptr)->setClearColor(*static_cast<QColor*>(color));
}

void QQuickWidget_SetFormat(void* ptr, void* format){
	static_cast<QQuickWidget*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QQuickWidget_SetSource(void* ptr, void* url){
	QMetaObject::invokeMethod(static_cast<QQuickWidget*>(ptr), "setSource", Q_ARG(QUrl, *static_cast<QUrl*>(url)));
}

void QQuickWidget_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWidget::Status)>(&QQuickWidget::statusChanged), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWidget::Status)>(&MyQQuickWidget::Signal_StatusChanged));;
}

void QQuickWidget_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickWidget*>(ptr), static_cast<void (QQuickWidget::*)(QQuickWidget::Status)>(&QQuickWidget::statusChanged), static_cast<MyQQuickWidget*>(ptr), static_cast<void (MyQQuickWidget::*)(QQuickWidget::Status)>(&MyQQuickWidget::Signal_StatusChanged));;
}

void QQuickWidget_DestroyQQuickWidget(void* ptr){
	static_cast<QQuickWidget*>(ptr)->~QQuickWidget();
}

