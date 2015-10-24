#include "qopenglpaintdevice.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QOpenGLPaintDevice>
#include "_cgo_export.h"

class MyQOpenGLPaintDevice: public QOpenGLPaintDevice {
public:
};

QtObjectPtr QOpenGLPaintDevice_NewQOpenGLPaintDevice(){
	return new QOpenGLPaintDevice();
}

QtObjectPtr QOpenGLPaintDevice_NewQOpenGLPaintDevice2(QtObjectPtr size){
	return new QOpenGLPaintDevice(*static_cast<QSize*>(size));
}

QtObjectPtr QOpenGLPaintDevice_NewQOpenGLPaintDevice3(int width, int height){
	return new QOpenGLPaintDevice(width, height);
}

QtObjectPtr QOpenGLPaintDevice_Context(QtObjectPtr ptr){
	return static_cast<QOpenGLPaintDevice*>(ptr)->context();
}

void QOpenGLPaintDevice_EnsureActiveTarget(QtObjectPtr ptr){
	static_cast<QOpenGLPaintDevice*>(ptr)->ensureActiveTarget();
}

QtObjectPtr QOpenGLPaintDevice_PaintEngine(QtObjectPtr ptr){
	return static_cast<QOpenGLPaintDevice*>(ptr)->paintEngine();
}

int QOpenGLPaintDevice_PaintFlipped(QtObjectPtr ptr){
	return static_cast<QOpenGLPaintDevice*>(ptr)->paintFlipped();
}

void QOpenGLPaintDevice_SetPaintFlipped(QtObjectPtr ptr, int flipped){
	static_cast<QOpenGLPaintDevice*>(ptr)->setPaintFlipped(flipped != 0);
}

void QOpenGLPaintDevice_SetSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QOpenGLPaintDevice*>(ptr)->setSize(*static_cast<QSize*>(size));
}

void QOpenGLPaintDevice_DestroyQOpenGLPaintDevice(QtObjectPtr ptr){
	static_cast<QOpenGLPaintDevice*>(ptr)->~QOpenGLPaintDevice();
}

