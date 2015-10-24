#include "qgraphicspixmapitem.h"
#include <QString>
#include <QStyleOptionGraphicsItem>
#include <QPainter>
#include <QPoint>
#include <QPixmap>
#include <QGraphicsItem>
#include <QWidget>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QPointF>
#include <QGraphicsPixmapItem>
#include "_cgo_export.h"

class MyQGraphicsPixmapItem: public QGraphicsPixmapItem {
public:
};

QtObjectPtr QGraphicsPixmapItem_NewQGraphicsPixmapItem(QtObjectPtr parent){
	return new QGraphicsPixmapItem(static_cast<QGraphicsItem*>(parent));
}

QtObjectPtr QGraphicsPixmapItem_NewQGraphicsPixmapItem2(QtObjectPtr pixmap, QtObjectPtr parent){
	return new QGraphicsPixmapItem(*static_cast<QPixmap*>(pixmap), static_cast<QGraphicsItem*>(parent));
}

int QGraphicsPixmapItem_Contains(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QGraphicsPixmapItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsPixmapItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QGraphicsPixmapItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsPixmapItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsPixmapItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsPixmapItem_SetOffset(QtObjectPtr ptr, QtObjectPtr offset){
	static_cast<QGraphicsPixmapItem*>(ptr)->setOffset(*static_cast<QPointF*>(offset));
}

void QGraphicsPixmapItem_SetPixmap(QtObjectPtr ptr, QtObjectPtr pixmap){
	static_cast<QGraphicsPixmapItem*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap));
}

void QGraphicsPixmapItem_SetShapeMode(QtObjectPtr ptr, int mode){
	static_cast<QGraphicsPixmapItem*>(ptr)->setShapeMode(static_cast<QGraphicsPixmapItem::ShapeMode>(mode));
}

void QGraphicsPixmapItem_SetTransformationMode(QtObjectPtr ptr, int mode){
	static_cast<QGraphicsPixmapItem*>(ptr)->setTransformationMode(static_cast<Qt::TransformationMode>(mode));
}

int QGraphicsPixmapItem_ShapeMode(QtObjectPtr ptr){
	return static_cast<QGraphicsPixmapItem*>(ptr)->shapeMode();
}

int QGraphicsPixmapItem_TransformationMode(QtObjectPtr ptr){
	return static_cast<QGraphicsPixmapItem*>(ptr)->transformationMode();
}

int QGraphicsPixmapItem_Type(QtObjectPtr ptr){
	return static_cast<QGraphicsPixmapItem*>(ptr)->type();
}

void QGraphicsPixmapItem_DestroyQGraphicsPixmapItem(QtObjectPtr ptr){
	static_cast<QGraphicsPixmapItem*>(ptr)->~QGraphicsPixmapItem();
}

