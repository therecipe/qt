#include "qoffscreensurface.h"
#include <QModelIndex>
#include <QSurfaceFormat>
#include <QScreen>
#include <QObject>
#include <QSurface>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QOffscreenSurface>
#include "_cgo_export.h"

class MyQOffscreenSurface: public QOffscreenSurface {
public:
void Signal_ScreenChanged(QScreen * screen){callbackQOffscreenSurfaceScreenChanged(this->objectName().toUtf8().data(), screen);};
};

void* QOffscreenSurface_NewQOffscreenSurface(void* targetScreen){
	return new QOffscreenSurface(static_cast<QScreen*>(targetScreen));
}

void QOffscreenSurface_Create(void* ptr){
	static_cast<QOffscreenSurface*>(ptr)->create();
}

void QOffscreenSurface_Destroy(void* ptr){
	static_cast<QOffscreenSurface*>(ptr)->destroy();
}

int QOffscreenSurface_IsValid(void* ptr){
	return static_cast<QOffscreenSurface*>(ptr)->isValid();
}

void* QOffscreenSurface_Screen(void* ptr){
	return static_cast<QOffscreenSurface*>(ptr)->screen();
}

void QOffscreenSurface_ConnectScreenChanged(void* ptr){
	QObject::connect(static_cast<QOffscreenSurface*>(ptr), static_cast<void (QOffscreenSurface::*)(QScreen *)>(&QOffscreenSurface::screenChanged), static_cast<MyQOffscreenSurface*>(ptr), static_cast<void (MyQOffscreenSurface::*)(QScreen *)>(&MyQOffscreenSurface::Signal_ScreenChanged));;
}

void QOffscreenSurface_DisconnectScreenChanged(void* ptr){
	QObject::disconnect(static_cast<QOffscreenSurface*>(ptr), static_cast<void (QOffscreenSurface::*)(QScreen *)>(&QOffscreenSurface::screenChanged), static_cast<MyQOffscreenSurface*>(ptr), static_cast<void (MyQOffscreenSurface::*)(QScreen *)>(&MyQOffscreenSurface::Signal_ScreenChanged));;
}

void QOffscreenSurface_SetFormat(void* ptr, void* format){
	static_cast<QOffscreenSurface*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QOffscreenSurface_SetScreen(void* ptr, void* newScreen){
	static_cast<QOffscreenSurface*>(ptr)->setScreen(static_cast<QScreen*>(newScreen));
}

int QOffscreenSurface_SurfaceType(void* ptr){
	return static_cast<QOffscreenSurface*>(ptr)->surfaceType();
}

void QOffscreenSurface_DestroyQOffscreenSurface(void* ptr){
	static_cast<QOffscreenSurface*>(ptr)->~QOffscreenSurface();
}

