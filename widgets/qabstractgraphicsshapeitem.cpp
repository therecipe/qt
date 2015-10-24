#include "qabstractgraphicsshapeitem.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsItem>
#include <QBrush>
#include <QPen>
#include <QString>
#include <QAbstractGraphicsShapeItem>
#include "_cgo_export.h"

class MyQAbstractGraphicsShapeItem: public QAbstractGraphicsShapeItem {
public:
};

int QAbstractGraphicsShapeItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QAbstractGraphicsShapeItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QAbstractGraphicsShapeItem_SetBrush(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QAbstractGraphicsShapeItem*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QAbstractGraphicsShapeItem_SetPen(QtObjectPtr ptr, QtObjectPtr pen){
	static_cast<QAbstractGraphicsShapeItem*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QAbstractGraphicsShapeItem_DestroyQAbstractGraphicsShapeItem(QtObjectPtr ptr){
	static_cast<QAbstractGraphicsShapeItem*>(ptr)->~QAbstractGraphicsShapeItem();
}

