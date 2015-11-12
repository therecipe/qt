#include "qbackingstore.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QRegion>
#include <QSize>
#include <QWindow>
#include <QBackingStore>
#include "_cgo_export.h"

class MyQBackingStore: public QBackingStore {
public:
};

void* QBackingStore_PaintDevice(void* ptr){
	return static_cast<QBackingStore*>(ptr)->paintDevice();
}

void* QBackingStore_NewQBackingStore(void* window){
	return new QBackingStore(static_cast<QWindow*>(window));
}

void QBackingStore_BeginPaint(void* ptr, void* region){
	static_cast<QBackingStore*>(ptr)->beginPaint(*static_cast<QRegion*>(region));
}

void QBackingStore_EndPaint(void* ptr){
	static_cast<QBackingStore*>(ptr)->endPaint();
}

void QBackingStore_Flush(void* ptr, void* region, void* win, void* offset){
	static_cast<QBackingStore*>(ptr)->flush(*static_cast<QRegion*>(region), static_cast<QWindow*>(win), *static_cast<QPoint*>(offset));
}

int QBackingStore_HasStaticContents(void* ptr){
	return static_cast<QBackingStore*>(ptr)->hasStaticContents();
}

void QBackingStore_Resize(void* ptr, void* size){
	static_cast<QBackingStore*>(ptr)->resize(*static_cast<QSize*>(size));
}

int QBackingStore_Scroll(void* ptr, void* area, int dx, int dy){
	return static_cast<QBackingStore*>(ptr)->scroll(*static_cast<QRegion*>(area), dx, dy);
}

void QBackingStore_SetStaticContents(void* ptr, void* region){
	static_cast<QBackingStore*>(ptr)->setStaticContents(*static_cast<QRegion*>(region));
}

void* QBackingStore_StaticContents(void* ptr){
	return new QRegion(static_cast<QBackingStore*>(ptr)->staticContents());
}

void* QBackingStore_Window(void* ptr){
	return static_cast<QBackingStore*>(ptr)->window();
}

void QBackingStore_DestroyQBackingStore(void* ptr){
	static_cast<QBackingStore*>(ptr)->~QBackingStore();
}

