#include "qsurface.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QSurface>
#include "_cgo_export.h"

class MyQSurface: public QSurface {
public:
};

int QSurface_SupportsOpenGL(void* ptr){
	return static_cast<QSurface*>(ptr)->supportsOpenGL();
}

int QSurface_SurfaceClass(void* ptr){
	return static_cast<QSurface*>(ptr)->surfaceClass();
}

int QSurface_SurfaceType(void* ptr){
	return static_cast<QSurface*>(ptr)->surfaceType();
}

void QSurface_DestroyQSurface(void* ptr){
	static_cast<QSurface*>(ptr)->~QSurface();
}

