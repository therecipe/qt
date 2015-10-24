#include "qbackingstore.h"
#include <QRegion>
#include <QSize>
#include <QWindow>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBackingStore>
#include "_cgo_export.h"

class MyQBackingStore: public QBackingStore {
public:
};

QtObjectPtr QBackingStore_PaintDevice(QtObjectPtr ptr){
	return static_cast<QBackingStore*>(ptr)->paintDevice();
}

QtObjectPtr QBackingStore_NewQBackingStore(QtObjectPtr window){
	return new QBackingStore(static_cast<QWindow*>(window));
}

void QBackingStore_BeginPaint(QtObjectPtr ptr, QtObjectPtr region){
	static_cast<QBackingStore*>(ptr)->beginPaint(*static_cast<QRegion*>(region));
}

void QBackingStore_EndPaint(QtObjectPtr ptr){
	static_cast<QBackingStore*>(ptr)->endPaint();
}

void QBackingStore_Flush(QtObjectPtr ptr, QtObjectPtr region, QtObjectPtr win, QtObjectPtr offset){
	static_cast<QBackingStore*>(ptr)->flush(*static_cast<QRegion*>(region), static_cast<QWindow*>(win), *static_cast<QPoint*>(offset));
}

int QBackingStore_HasStaticContents(QtObjectPtr ptr){
	return static_cast<QBackingStore*>(ptr)->hasStaticContents();
}

void QBackingStore_Resize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QBackingStore*>(ptr)->resize(*static_cast<QSize*>(size));
}

int QBackingStore_Scroll(QtObjectPtr ptr, QtObjectPtr area, int dx, int dy){
	return static_cast<QBackingStore*>(ptr)->scroll(*static_cast<QRegion*>(area), dx, dy);
}

void QBackingStore_SetStaticContents(QtObjectPtr ptr, QtObjectPtr region){
	static_cast<QBackingStore*>(ptr)->setStaticContents(*static_cast<QRegion*>(region));
}

QtObjectPtr QBackingStore_Window(QtObjectPtr ptr){
	return static_cast<QBackingStore*>(ptr)->window();
}

void QBackingStore_DestroyQBackingStore(QtObjectPtr ptr){
	static_cast<QBackingStore*>(ptr)->~QBackingStore();
}

