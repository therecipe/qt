#include "qoffscreensurface.h"
#include <QSurfaceFormat>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScreen>
#include <QSurface>
#include <QOffscreenSurface>
#include "_cgo_export.h"

class MyQOffscreenSurface: public QOffscreenSurface {
public:
void Signal_ScreenChanged(QScreen * screen){callbackQOffscreenSurfaceScreenChanged(this->objectName().toUtf8().data(), screen);};
};

QtObjectPtr QOffscreenSurface_NewQOffscreenSurface(QtObjectPtr targetScreen){
	return new QOffscreenSurface(static_cast<QScreen*>(targetScreen));
}

void QOffscreenSurface_Create(QtObjectPtr ptr){
	static_cast<QOffscreenSurface*>(ptr)->create();
}

void QOffscreenSurface_Destroy(QtObjectPtr ptr){
	static_cast<QOffscreenSurface*>(ptr)->destroy();
}

int QOffscreenSurface_IsValid(QtObjectPtr ptr){
	return static_cast<QOffscreenSurface*>(ptr)->isValid();
}

QtObjectPtr QOffscreenSurface_Screen(QtObjectPtr ptr){
	return static_cast<QOffscreenSurface*>(ptr)->screen();
}

void QOffscreenSurface_ConnectScreenChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QOffscreenSurface*>(ptr), static_cast<void (QOffscreenSurface::*)(QScreen *)>(&QOffscreenSurface::screenChanged), static_cast<MyQOffscreenSurface*>(ptr), static_cast<void (MyQOffscreenSurface::*)(QScreen *)>(&MyQOffscreenSurface::Signal_ScreenChanged));;
}

void QOffscreenSurface_DisconnectScreenChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QOffscreenSurface*>(ptr), static_cast<void (QOffscreenSurface::*)(QScreen *)>(&QOffscreenSurface::screenChanged), static_cast<MyQOffscreenSurface*>(ptr), static_cast<void (MyQOffscreenSurface::*)(QScreen *)>(&MyQOffscreenSurface::Signal_ScreenChanged));;
}

void QOffscreenSurface_SetFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QOffscreenSurface*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QOffscreenSurface_SetScreen(QtObjectPtr ptr, QtObjectPtr newScreen){
	static_cast<QOffscreenSurface*>(ptr)->setScreen(static_cast<QScreen*>(newScreen));
}

int QOffscreenSurface_SurfaceType(QtObjectPtr ptr){
	return static_cast<QOffscreenSurface*>(ptr)->surfaceType();
}

void QOffscreenSurface_DestroyQOffscreenSurface(QtObjectPtr ptr){
	static_cast<QOffscreenSurface*>(ptr)->~QOffscreenSurface();
}

