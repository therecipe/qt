#include "qsvgrenderer.h"
#include <QVariant>
#include <QModelIndex>
#include <QMetaObject>
#include <QRectF>
#include <QRect>
#include <QXmlStreamReader>
#include <QString>
#include <QUrl>
#include <QObject>
#include <QPainter>
#include <QByteArray>
#include <QSvgRenderer>
#include "_cgo_export.h"

class MyQSvgRenderer: public QSvgRenderer {
public:
void Signal_RepaintNeeded(){callbackQSvgRendererRepaintNeeded(this->objectName().toUtf8().data());};
};

int QSvgRenderer_FramesPerSecond(void* ptr){
	return static_cast<QSvgRenderer*>(ptr)->framesPerSecond();
}

void QSvgRenderer_SetFramesPerSecond(void* ptr, int num){
	static_cast<QSvgRenderer*>(ptr)->setFramesPerSecond(num);
}

void QSvgRenderer_SetViewBox(void* ptr, void* viewbox){
	static_cast<QSvgRenderer*>(ptr)->setViewBox(*static_cast<QRect*>(viewbox));
}

void QSvgRenderer_SetViewBox2(void* ptr, void* viewbox){
	static_cast<QSvgRenderer*>(ptr)->setViewBox(*static_cast<QRectF*>(viewbox));
}

void* QSvgRenderer_NewQSvgRenderer(void* parent){
	return new QSvgRenderer(static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer4(void* contents, void* parent){
	return new QSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer3(void* contents, void* parent){
	return new QSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer2(char* filename, void* parent){
	return new QSvgRenderer(QString(filename), static_cast<QObject*>(parent));
}

int QSvgRenderer_Animated(void* ptr){
	return static_cast<QSvgRenderer*>(ptr)->animated();
}

int QSvgRenderer_ElementExists(void* ptr, char* id){
	return static_cast<QSvgRenderer*>(ptr)->elementExists(QString(id));
}

int QSvgRenderer_IsValid(void* ptr){
	return static_cast<QSvgRenderer*>(ptr)->isValid();
}

int QSvgRenderer_Load3(void* ptr, void* contents){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QXmlStreamReader*, static_cast<QXmlStreamReader*>(contents)));
}

int QSvgRenderer_Load2(void* ptr, void* contents){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QByteArray, *static_cast<QByteArray*>(contents)));
}

int QSvgRenderer_Load(void* ptr, char* filename){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QString, QString(filename)));
}

void QSvgRenderer_Render(void* ptr, void* painter){
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)));
}

void QSvgRenderer_Render2(void* ptr, void* painter, void* bounds){
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)), Q_ARG(QRectF, *static_cast<QRectF*>(bounds)));
}

void QSvgRenderer_Render3(void* ptr, void* painter, char* elementId, void* bounds){
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)), Q_ARG(QString, QString(elementId)), Q_ARG(QRectF, *static_cast<QRectF*>(bounds)));
}

void QSvgRenderer_ConnectRepaintNeeded(void* ptr){
	QObject::connect(static_cast<QSvgRenderer*>(ptr), static_cast<void (QSvgRenderer::*)()>(&QSvgRenderer::repaintNeeded), static_cast<MyQSvgRenderer*>(ptr), static_cast<void (MyQSvgRenderer::*)()>(&MyQSvgRenderer::Signal_RepaintNeeded));;
}

void QSvgRenderer_DisconnectRepaintNeeded(void* ptr){
	QObject::disconnect(static_cast<QSvgRenderer*>(ptr), static_cast<void (QSvgRenderer::*)()>(&QSvgRenderer::repaintNeeded), static_cast<MyQSvgRenderer*>(ptr), static_cast<void (MyQSvgRenderer::*)()>(&MyQSvgRenderer::Signal_RepaintNeeded));;
}

void QSvgRenderer_DestroyQSvgRenderer(void* ptr){
	static_cast<QSvgRenderer*>(ptr)->~QSvgRenderer();
}

