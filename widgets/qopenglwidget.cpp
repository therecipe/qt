#include "qopenglwidget.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QSurface>
#include <QSurfaceFormat>
#include <QWidget>
#include <QString>
#include <QOpenGLWidget>
#include "_cgo_export.h"

class MyQOpenGLWidget: public QOpenGLWidget {
public:
void Signal_AboutToCompose(){callbackQOpenGLWidgetAboutToCompose(this->objectName().toUtf8().data());};
void Signal_AboutToResize(){callbackQOpenGLWidgetAboutToResize(this->objectName().toUtf8().data());};
void Signal_FrameSwapped(){callbackQOpenGLWidgetFrameSwapped(this->objectName().toUtf8().data());};
void Signal_Resized(){callbackQOpenGLWidgetResized(this->objectName().toUtf8().data());};
};

QtObjectPtr QOpenGLWidget_NewQOpenGLWidget(QtObjectPtr parent, int f){
	return new QOpenGLWidget(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QOpenGLWidget_ConnectAboutToCompose(QtObjectPtr ptr){
	QObject::connect(static_cast<QOpenGLWidget*>(ptr), static_cast<void (QOpenGLWidget::*)()>(&QOpenGLWidget::aboutToCompose), static_cast<MyQOpenGLWidget*>(ptr), static_cast<void (MyQOpenGLWidget::*)()>(&MyQOpenGLWidget::Signal_AboutToCompose));;
}

void QOpenGLWidget_DisconnectAboutToCompose(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QOpenGLWidget*>(ptr), static_cast<void (QOpenGLWidget::*)()>(&QOpenGLWidget::aboutToCompose), static_cast<MyQOpenGLWidget*>(ptr), static_cast<void (MyQOpenGLWidget::*)()>(&MyQOpenGLWidget::Signal_AboutToCompose));;
}

void QOpenGLWidget_ConnectAboutToResize(QtObjectPtr ptr){
	QObject::connect(static_cast<QOpenGLWidget*>(ptr), static_cast<void (QOpenGLWidget::*)()>(&QOpenGLWidget::aboutToResize), static_cast<MyQOpenGLWidget*>(ptr), static_cast<void (MyQOpenGLWidget::*)()>(&MyQOpenGLWidget::Signal_AboutToResize));;
}

void QOpenGLWidget_DisconnectAboutToResize(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QOpenGLWidget*>(ptr), static_cast<void (QOpenGLWidget::*)()>(&QOpenGLWidget::aboutToResize), static_cast<MyQOpenGLWidget*>(ptr), static_cast<void (MyQOpenGLWidget::*)()>(&MyQOpenGLWidget::Signal_AboutToResize));;
}

QtObjectPtr QOpenGLWidget_Context(QtObjectPtr ptr){
	return static_cast<QOpenGLWidget*>(ptr)->context();
}

void QOpenGLWidget_DoneCurrent(QtObjectPtr ptr){
	static_cast<QOpenGLWidget*>(ptr)->doneCurrent();
}

int QOpenGLWidget_IsValid(QtObjectPtr ptr){
	return static_cast<QOpenGLWidget*>(ptr)->isValid();
}

void QOpenGLWidget_MakeCurrent(QtObjectPtr ptr){
	static_cast<QOpenGLWidget*>(ptr)->makeCurrent();
}

void QOpenGLWidget_ConnectFrameSwapped(QtObjectPtr ptr){
	QObject::connect(static_cast<QOpenGLWidget*>(ptr), static_cast<void (QOpenGLWidget::*)()>(&QOpenGLWidget::frameSwapped), static_cast<MyQOpenGLWidget*>(ptr), static_cast<void (MyQOpenGLWidget::*)()>(&MyQOpenGLWidget::Signal_FrameSwapped));;
}

void QOpenGLWidget_DisconnectFrameSwapped(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QOpenGLWidget*>(ptr), static_cast<void (QOpenGLWidget::*)()>(&QOpenGLWidget::frameSwapped), static_cast<MyQOpenGLWidget*>(ptr), static_cast<void (MyQOpenGLWidget::*)()>(&MyQOpenGLWidget::Signal_FrameSwapped));;
}

void QOpenGLWidget_ConnectResized(QtObjectPtr ptr){
	QObject::connect(static_cast<QOpenGLWidget*>(ptr), static_cast<void (QOpenGLWidget::*)()>(&QOpenGLWidget::resized), static_cast<MyQOpenGLWidget*>(ptr), static_cast<void (MyQOpenGLWidget::*)()>(&MyQOpenGLWidget::Signal_Resized));;
}

void QOpenGLWidget_DisconnectResized(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QOpenGLWidget*>(ptr), static_cast<void (QOpenGLWidget::*)()>(&QOpenGLWidget::resized), static_cast<MyQOpenGLWidget*>(ptr), static_cast<void (MyQOpenGLWidget::*)()>(&MyQOpenGLWidget::Signal_Resized));;
}

void QOpenGLWidget_SetFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QOpenGLWidget*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QOpenGLWidget_SetUpdateBehavior(QtObjectPtr ptr, int updateBehavior){
	static_cast<QOpenGLWidget*>(ptr)->setUpdateBehavior(static_cast<QOpenGLWidget::UpdateBehavior>(updateBehavior));
}

int QOpenGLWidget_UpdateBehavior(QtObjectPtr ptr){
	return static_cast<QOpenGLWidget*>(ptr)->updateBehavior();
}

void QOpenGLWidget_DestroyQOpenGLWidget(QtObjectPtr ptr){
	static_cast<QOpenGLWidget*>(ptr)->~QOpenGLWidget();
}

