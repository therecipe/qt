#include "qabstractvideosurface.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QVideoFrame>
#include <QVideoSurfaceFormat>
#include <QAbstractVideoSurface>
#include "_cgo_export.h"

class MyQAbstractVideoSurface: public QAbstractVideoSurface {
public:
void Signal_ActiveChanged(bool active){callbackQAbstractVideoSurfaceActiveChanged(this->objectName().toUtf8().data(), active);};
void Signal_SupportedFormatsChanged(){callbackQAbstractVideoSurfaceSupportedFormatsChanged(this->objectName().toUtf8().data());};
};

void QAbstractVideoSurface_ConnectActiveChanged(void* ptr){
	QObject::connect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)(bool)>(&QAbstractVideoSurface::activeChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)(bool)>(&MyQAbstractVideoSurface::Signal_ActiveChanged));;
}

void QAbstractVideoSurface_DisconnectActiveChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)(bool)>(&QAbstractVideoSurface::activeChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)(bool)>(&MyQAbstractVideoSurface::Signal_ActiveChanged));;
}

int QAbstractVideoSurface_Error(void* ptr){
	return static_cast<QAbstractVideoSurface*>(ptr)->error();
}

int QAbstractVideoSurface_IsActive(void* ptr){
	return static_cast<QAbstractVideoSurface*>(ptr)->isActive();
}

int QAbstractVideoSurface_IsFormatSupported(void* ptr, void* format){
	return static_cast<QAbstractVideoSurface*>(ptr)->isFormatSupported(*static_cast<QVideoSurfaceFormat*>(format));
}

int QAbstractVideoSurface_Present(void* ptr, void* frame){
	return static_cast<QAbstractVideoSurface*>(ptr)->present(*static_cast<QVideoFrame*>(frame));
}

int QAbstractVideoSurface_Start(void* ptr, void* format){
	return static_cast<QAbstractVideoSurface*>(ptr)->start(*static_cast<QVideoSurfaceFormat*>(format));
}

void QAbstractVideoSurface_Stop(void* ptr){
	static_cast<QAbstractVideoSurface*>(ptr)->stop();
}

void QAbstractVideoSurface_ConnectSupportedFormatsChanged(void* ptr){
	QObject::connect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)()>(&QAbstractVideoSurface::supportedFormatsChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)()>(&MyQAbstractVideoSurface::Signal_SupportedFormatsChanged));;
}

void QAbstractVideoSurface_DisconnectSupportedFormatsChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractVideoSurface*>(ptr), static_cast<void (QAbstractVideoSurface::*)()>(&QAbstractVideoSurface::supportedFormatsChanged), static_cast<MyQAbstractVideoSurface*>(ptr), static_cast<void (MyQAbstractVideoSurface::*)()>(&MyQAbstractVideoSurface::Signal_SupportedFormatsChanged));;
}

void QAbstractVideoSurface_DestroyQAbstractVideoSurface(void* ptr){
	static_cast<QAbstractVideoSurface*>(ptr)->~QAbstractVideoSurface();
}

