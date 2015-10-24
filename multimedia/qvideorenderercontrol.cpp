#include "qvideorenderercontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractVideoSurface>
#include <QVideoRendererControl>
#include "_cgo_export.h"

class MyQVideoRendererControl: public QVideoRendererControl {
public:
};

void QVideoRendererControl_SetSurface(QtObjectPtr ptr, QtObjectPtr surface){
	static_cast<QVideoRendererControl*>(ptr)->setSurface(static_cast<QAbstractVideoSurface*>(surface));
}

QtObjectPtr QVideoRendererControl_Surface(QtObjectPtr ptr){
	return static_cast<QVideoRendererControl*>(ptr)->surface();
}

void QVideoRendererControl_DestroyQVideoRendererControl(QtObjectPtr ptr){
	static_cast<QVideoRendererControl*>(ptr)->~QVideoRendererControl();
}

