#include "qpagedpaintdevice.h"
#include <QString>
#include <QMargins>
#include <QPageSize>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPageLayout>
#include <QMarginsF>
#include <QSizeF>
#include <QSize>
#include <QPagedPaintDevice>
#include "_cgo_export.h"

class MyQPagedPaintDevice: public QPagedPaintDevice {
public:
};

int QPagedPaintDevice_NewPage(QtObjectPtr ptr){
	return static_cast<QPagedPaintDevice*>(ptr)->newPage();
}

int QPagedPaintDevice_PageSize(QtObjectPtr ptr){
	return static_cast<QPagedPaintDevice*>(ptr)->pageSize();
}

int QPagedPaintDevice_SetPageLayout(QtObjectPtr ptr, QtObjectPtr newPageLayout){
	return static_cast<QPagedPaintDevice*>(ptr)->setPageLayout(*static_cast<QPageLayout*>(newPageLayout));
}

int QPagedPaintDevice_SetPageMargins(QtObjectPtr ptr, QtObjectPtr margins){
	return static_cast<QPagedPaintDevice*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins));
}

int QPagedPaintDevice_SetPageMargins2(QtObjectPtr ptr, QtObjectPtr margins, int units){
	return static_cast<QPagedPaintDevice*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins), static_cast<QPageLayout::Unit>(units));
}

int QPagedPaintDevice_SetPageOrientation(QtObjectPtr ptr, int orientation){
	return static_cast<QPagedPaintDevice*>(ptr)->setPageOrientation(static_cast<QPageLayout::Orientation>(orientation));
}

int QPagedPaintDevice_SetPageSize(QtObjectPtr ptr, QtObjectPtr pageSize){
	return static_cast<QPagedPaintDevice*>(ptr)->setPageSize(*static_cast<QPageSize*>(pageSize));
}

void QPagedPaintDevice_SetPageSize2(QtObjectPtr ptr, int size){
	static_cast<QPagedPaintDevice*>(ptr)->setPageSize(static_cast<QPagedPaintDevice::PageSize>(size));
}

void QPagedPaintDevice_SetPageSizeMM(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QPagedPaintDevice*>(ptr)->setPageSizeMM(*static_cast<QSizeF*>(size));
}

void QPagedPaintDevice_DestroyQPagedPaintDevice(QtObjectPtr ptr){
	static_cast<QPagedPaintDevice*>(ptr)->~QPagedPaintDevice();
}

