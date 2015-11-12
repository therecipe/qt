#include "qabstractgraphicsshapeitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBrush>
#include <QPen>
#include <QGraphicsItem>
#include <QAbstractGraphicsShapeItem>
#include "_cgo_export.h"

class MyQAbstractGraphicsShapeItem: public QAbstractGraphicsShapeItem {
public:
};

void* QAbstractGraphicsShapeItem_Brush(void* ptr){
	return new QBrush(static_cast<QAbstractGraphicsShapeItem*>(ptr)->brush());
}

int QAbstractGraphicsShapeItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QAbstractGraphicsShapeItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QAbstractGraphicsShapeItem_SetBrush(void* ptr, void* brush){
	static_cast<QAbstractGraphicsShapeItem*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QAbstractGraphicsShapeItem_SetPen(void* ptr, void* pen){
	static_cast<QAbstractGraphicsShapeItem*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QAbstractGraphicsShapeItem_DestroyQAbstractGraphicsShapeItem(void* ptr){
	static_cast<QAbstractGraphicsShapeItem*>(ptr)->~QAbstractGraphicsShapeItem();
}

