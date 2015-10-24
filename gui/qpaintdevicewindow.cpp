#include "qpaintdevicewindow.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QRegion>
#include <QPaintDevice>
#include <QMetaObject>
#include <QPaintDeviceWindow>
#include "_cgo_export.h"

class MyQPaintDeviceWindow: public QPaintDeviceWindow {
public:
};

void QPaintDeviceWindow_Update3(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "update");
}

void QPaintDeviceWindow_Update(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QPaintDeviceWindow*>(ptr)->update(*static_cast<QRect*>(rect));
}

void QPaintDeviceWindow_Update2(QtObjectPtr ptr, QtObjectPtr region){
	static_cast<QPaintDeviceWindow*>(ptr)->update(*static_cast<QRegion*>(region));
}

