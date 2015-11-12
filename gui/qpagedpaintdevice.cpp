#include "qpagedpaintdevice.h"
#include <QSize>
#include <QSizeF>
#include <QPageLayout>
#include <QPageSize>
#include <QMargins>
#include <QVariant>
#include <QUrl>
#include <QMarginsF>
#include <QString>
#include <QModelIndex>
#include <QPagedPaintDevice>
#include "_cgo_export.h"

class MyQPagedPaintDevice: public QPagedPaintDevice {
public:
};

int QPagedPaintDevice_NewPage(void* ptr){
	return static_cast<QPagedPaintDevice*>(ptr)->newPage();
}

int QPagedPaintDevice_PageSize(void* ptr){
	return static_cast<QPagedPaintDevice*>(ptr)->pageSize();
}

int QPagedPaintDevice_SetPageLayout(void* ptr, void* newPageLayout){
	return static_cast<QPagedPaintDevice*>(ptr)->setPageLayout(*static_cast<QPageLayout*>(newPageLayout));
}

int QPagedPaintDevice_SetPageMargins(void* ptr, void* margins){
	return static_cast<QPagedPaintDevice*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins));
}

int QPagedPaintDevice_SetPageMargins2(void* ptr, void* margins, int units){
	return static_cast<QPagedPaintDevice*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins), static_cast<QPageLayout::Unit>(units));
}

int QPagedPaintDevice_SetPageOrientation(void* ptr, int orientation){
	return static_cast<QPagedPaintDevice*>(ptr)->setPageOrientation(static_cast<QPageLayout::Orientation>(orientation));
}

int QPagedPaintDevice_SetPageSize(void* ptr, void* pageSize){
	return static_cast<QPagedPaintDevice*>(ptr)->setPageSize(*static_cast<QPageSize*>(pageSize));
}

void QPagedPaintDevice_SetPageSize2(void* ptr, int size){
	static_cast<QPagedPaintDevice*>(ptr)->setPageSize(static_cast<QPagedPaintDevice::PageSize>(size));
}

void QPagedPaintDevice_SetPageSizeMM(void* ptr, void* size){
	static_cast<QPagedPaintDevice*>(ptr)->setPageSizeMM(*static_cast<QSizeF*>(size));
}

void QPagedPaintDevice_DestroyQPagedPaintDevice(void* ptr){
	static_cast<QPagedPaintDevice*>(ptr)->~QPagedPaintDevice();
}

