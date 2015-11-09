#include "qpaintdevicewindow.h"
#include <QRect>
#include <QRegion>
#include <QPaintDevice>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPaintDeviceWindow>
#include "_cgo_export.h"

class MyQPaintDeviceWindow: public QPaintDeviceWindow {
public:
};

void QPaintDeviceWindow_Update3(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "update");
}

void QPaintDeviceWindow_Update(void* ptr, void* rect){
	static_cast<QPaintDeviceWindow*>(ptr)->update(*static_cast<QRect*>(rect));
}

void QPaintDeviceWindow_Update2(void* ptr, void* region){
	static_cast<QPaintDeviceWindow*>(ptr)->update(*static_cast<QRegion*>(region));
}

