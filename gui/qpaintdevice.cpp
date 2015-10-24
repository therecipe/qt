#include "qpaintdevice.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPaintDevice>
#include "_cgo_export.h"

class MyQPaintDevice: public QPaintDevice {
public:
};

void QPaintDevice_DestroyQPaintDevice(QtObjectPtr ptr){
	static_cast<QPaintDevice*>(ptr)->~QPaintDevice();
}

int QPaintDevice_ColorCount(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->colorCount();
}

int QPaintDevice_Depth(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->depth();
}

int QPaintDevice_DevicePixelRatio(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->devicePixelRatio();
}

int QPaintDevice_Height(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->height();
}

int QPaintDevice_HeightMM(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->heightMM();
}

int QPaintDevice_LogicalDpiX(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->logicalDpiX();
}

int QPaintDevice_LogicalDpiY(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->logicalDpiY();
}

QtObjectPtr QPaintDevice_PaintEngine(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->paintEngine();
}

int QPaintDevice_PaintingActive(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->paintingActive();
}

int QPaintDevice_PhysicalDpiX(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->physicalDpiX();
}

int QPaintDevice_PhysicalDpiY(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->physicalDpiY();
}

int QPaintDevice_Width(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->width();
}

int QPaintDevice_WidthMM(QtObjectPtr ptr){
	return static_cast<QPaintDevice*>(ptr)->widthMM();
}

