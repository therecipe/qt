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

void QPaintDevice_DestroyQPaintDevice(void* ptr){
	static_cast<QPaintDevice*>(ptr)->~QPaintDevice();
}

int QPaintDevice_ColorCount(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->colorCount();
}

int QPaintDevice_Depth(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->depth();
}

int QPaintDevice_DevicePixelRatio(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->devicePixelRatio();
}

int QPaintDevice_Height(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->height();
}

int QPaintDevice_HeightMM(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->heightMM();
}

int QPaintDevice_LogicalDpiX(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->logicalDpiX();
}

int QPaintDevice_LogicalDpiY(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->logicalDpiY();
}

void* QPaintDevice_PaintEngine(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->paintEngine();
}

int QPaintDevice_PaintingActive(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->paintingActive();
}

int QPaintDevice_PhysicalDpiX(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->physicalDpiX();
}

int QPaintDevice_PhysicalDpiY(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->physicalDpiY();
}

int QPaintDevice_Width(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->width();
}

int QPaintDevice_WidthMM(void* ptr){
	return static_cast<QPaintDevice*>(ptr)->widthMM();
}

