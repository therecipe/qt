#include "qgraphicspathitem.h"
#include <QString>
#include <QUrl>
#include <QPoint>
#include <QStyle>
#include <QPainterPath>
#include <QVariant>
#include <QModelIndex>
#include <QPainter>
#include <QWidget>
#include <QStyleOptionGraphicsItem>
#include <QGraphicsItem>
#include <QStyleOption>
#include <QPointF>
#include <QGraphicsPathItem>
#include "_cgo_export.h"

class MyQGraphicsPathItem: public QGraphicsPathItem {
public:
};

int QGraphicsPathItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsPathItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsPathItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsPathItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsPathItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsPathItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsPathItem_SetPath(void* ptr, void* path){
	static_cast<QGraphicsPathItem*>(ptr)->setPath(*static_cast<QPainterPath*>(path));
}

int QGraphicsPathItem_Type(void* ptr){
	return static_cast<QGraphicsPathItem*>(ptr)->type();
}

void QGraphicsPathItem_DestroyQGraphicsPathItem(void* ptr){
	static_cast<QGraphicsPathItem*>(ptr)->~QGraphicsPathItem();
}

