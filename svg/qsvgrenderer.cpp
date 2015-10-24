#include "qsvgrenderer.h"
#include <QMetaObject>
#include <QObject>
#include <QXmlStreamReader>
#include <QRect>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QString>
#include <QPainter>
#include <QRectF>
#include <QSvgRenderer>
#include "_cgo_export.h"

class MyQSvgRenderer: public QSvgRenderer {
public:
void Signal_RepaintNeeded(){callbackQSvgRendererRepaintNeeded(this->objectName().toUtf8().data());};
};

int QSvgRenderer_FramesPerSecond(QtObjectPtr ptr){
	return static_cast<QSvgRenderer*>(ptr)->framesPerSecond();
}

void QSvgRenderer_SetFramesPerSecond(QtObjectPtr ptr, int num){
	static_cast<QSvgRenderer*>(ptr)->setFramesPerSecond(num);
}

void QSvgRenderer_SetViewBox(QtObjectPtr ptr, QtObjectPtr viewbox){
	static_cast<QSvgRenderer*>(ptr)->setViewBox(*static_cast<QRect*>(viewbox));
}

void QSvgRenderer_SetViewBox2(QtObjectPtr ptr, QtObjectPtr viewbox){
	static_cast<QSvgRenderer*>(ptr)->setViewBox(*static_cast<QRectF*>(viewbox));
}

QtObjectPtr QSvgRenderer_NewQSvgRenderer(QtObjectPtr parent){
	return new QSvgRenderer(static_cast<QObject*>(parent));
}

QtObjectPtr QSvgRenderer_NewQSvgRenderer4(QtObjectPtr contents, QtObjectPtr parent){
	return new QSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QObject*>(parent));
}

QtObjectPtr QSvgRenderer_NewQSvgRenderer3(QtObjectPtr contents, QtObjectPtr parent){
	return new QSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QObject*>(parent));
}

QtObjectPtr QSvgRenderer_NewQSvgRenderer2(char* filename, QtObjectPtr parent){
	return new QSvgRenderer(QString(filename), static_cast<QObject*>(parent));
}

int QSvgRenderer_Animated(QtObjectPtr ptr){
	return static_cast<QSvgRenderer*>(ptr)->animated();
}

int QSvgRenderer_ElementExists(QtObjectPtr ptr, char* id){
	return static_cast<QSvgRenderer*>(ptr)->elementExists(QString(id));
}

int QSvgRenderer_IsValid(QtObjectPtr ptr){
	return static_cast<QSvgRenderer*>(ptr)->isValid();
}

int QSvgRenderer_Load3(QtObjectPtr ptr, QtObjectPtr contents){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QXmlStreamReader*, static_cast<QXmlStreamReader*>(contents)));
}

int QSvgRenderer_Load2(QtObjectPtr ptr, QtObjectPtr contents){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QByteArray, *static_cast<QByteArray*>(contents)));
}

int QSvgRenderer_Load(QtObjectPtr ptr, char* filename){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QString, QString(filename)));
}

void QSvgRenderer_Render(QtObjectPtr ptr, QtObjectPtr painter){
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)));
}

void QSvgRenderer_Render2(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr bounds){
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)), Q_ARG(QRectF, *static_cast<QRectF*>(bounds)));
}

void QSvgRenderer_Render3(QtObjectPtr ptr, QtObjectPtr painter, char* elementId, QtObjectPtr bounds){
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)), Q_ARG(QString, QString(elementId)), Q_ARG(QRectF, *static_cast<QRectF*>(bounds)));
}

void QSvgRenderer_ConnectRepaintNeeded(QtObjectPtr ptr){
	QObject::connect(static_cast<QSvgRenderer*>(ptr), static_cast<void (QSvgRenderer::*)()>(&QSvgRenderer::repaintNeeded), static_cast<MyQSvgRenderer*>(ptr), static_cast<void (MyQSvgRenderer::*)()>(&MyQSvgRenderer::Signal_RepaintNeeded));;
}

void QSvgRenderer_DisconnectRepaintNeeded(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSvgRenderer*>(ptr), static_cast<void (QSvgRenderer::*)()>(&QSvgRenderer::repaintNeeded), static_cast<MyQSvgRenderer*>(ptr), static_cast<void (MyQSvgRenderer::*)()>(&MyQSvgRenderer::Signal_RepaintNeeded));;
}

void QSvgRenderer_DestroyQSvgRenderer(QtObjectPtr ptr){
	static_cast<QSvgRenderer*>(ptr)->~QSvgRenderer();
}

