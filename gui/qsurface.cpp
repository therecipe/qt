#include "qsurface.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QSurface>
#include "_cgo_export.h"

class MyQSurface: public QSurface {
public:
};

int QSurface_SupportsOpenGL(QtObjectPtr ptr){
	return static_cast<QSurface*>(ptr)->supportsOpenGL();
}

int QSurface_SurfaceClass(QtObjectPtr ptr){
	return static_cast<QSurface*>(ptr)->surfaceClass();
}

int QSurface_SurfaceType(QtObjectPtr ptr){
	return static_cast<QSurface*>(ptr)->surfaceType();
}

void QSurface_DestroyQSurface(QtObjectPtr ptr){
	static_cast<QSurface*>(ptr)->~QSurface();
}

