#include "qgraphicspixmapitem.h"
#include <QVariant>
#include <QStyle>
#include <QStyleOptionGraphicsItem>
#include <QStyleOption>
#include <QPixmap>
#include <QPoint>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QPainter>
#include <QWidget>
#include <QGraphicsItem>
#include <QPointF>
#include <QGraphicsPixmapItem>
#include "_cgo_export.h"

class MyQGraphicsPixmapItem: public QGraphicsPixmapItem {
public:
};

void* QGraphicsPixmapItem_NewQGraphicsPixmapItem(void* parent){
	return new QGraphicsPixmapItem(static_cast<QGraphicsItem*>(parent));
}

void* QGraphicsPixmapItem_NewQGraphicsPixmapItem2(void* pixmap, void* parent){
	return new QGraphicsPixmapItem(*static_cast<QPixmap*>(pixmap), static_cast<QGraphicsItem*>(parent));
}

int QGraphicsPixmapItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsPixmapItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsPixmapItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsPixmapItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsPixmapItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsPixmapItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsPixmapItem_SetOffset(void* ptr, void* offset){
	static_cast<QGraphicsPixmapItem*>(ptr)->setOffset(*static_cast<QPointF*>(offset));
}

void QGraphicsPixmapItem_SetOffset2(void* ptr, double x, double y){
	static_cast<QGraphicsPixmapItem*>(ptr)->setOffset(static_cast<qreal>(x), static_cast<qreal>(y));
}

void QGraphicsPixmapItem_SetPixmap(void* ptr, void* pixmap){
	static_cast<QGraphicsPixmapItem*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap));
}

void QGraphicsPixmapItem_SetShapeMode(void* ptr, int mode){
	static_cast<QGraphicsPixmapItem*>(ptr)->setShapeMode(static_cast<QGraphicsPixmapItem::ShapeMode>(mode));
}

void QGraphicsPixmapItem_SetTransformationMode(void* ptr, int mode){
	static_cast<QGraphicsPixmapItem*>(ptr)->setTransformationMode(static_cast<Qt::TransformationMode>(mode));
}

int QGraphicsPixmapItem_ShapeMode(void* ptr){
	return static_cast<QGraphicsPixmapItem*>(ptr)->shapeMode();
}

int QGraphicsPixmapItem_TransformationMode(void* ptr){
	return static_cast<QGraphicsPixmapItem*>(ptr)->transformationMode();
}

int QGraphicsPixmapItem_Type(void* ptr){
	return static_cast<QGraphicsPixmapItem*>(ptr)->type();
}

void QGraphicsPixmapItem_DestroyQGraphicsPixmapItem(void* ptr){
	static_cast<QGraphicsPixmapItem*>(ptr)->~QGraphicsPixmapItem();
}

