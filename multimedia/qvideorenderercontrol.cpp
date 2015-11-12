#include "qvideorenderercontrol.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractVideoSurface>
#include <QString>
#include <QVideoRendererControl>
#include "_cgo_export.h"

class MyQVideoRendererControl: public QVideoRendererControl {
public:
};

void QVideoRendererControl_SetSurface(void* ptr, void* surface){
	static_cast<QVideoRendererControl*>(ptr)->setSurface(static_cast<QAbstractVideoSurface*>(surface));
}

void* QVideoRendererControl_Surface(void* ptr){
	return static_cast<QVideoRendererControl*>(ptr)->surface();
}

void QVideoRendererControl_DestroyQVideoRendererControl(void* ptr){
	static_cast<QVideoRendererControl*>(ptr)->~QVideoRendererControl();
}

