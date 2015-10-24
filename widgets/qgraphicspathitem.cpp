#include "qgraphicspathitem.h"
#include <QStyleOptionGraphicsItem>
#include <QStyle>
#include <QPainterPath>
#include <QVariant>
#include <QModelIndex>
#include <QStyleOption>
#include <QPointF>
#include <QGraphicsItem>
#include <QPoint>
#include <QString>
#include <QUrl>
#include <QWidget>
#include <QPainter>
#include <QGraphicsPathItem>
#include "_cgo_export.h"

class MyQGraphicsPathItem: public QGraphicsPathItem {
public:
};

int QGraphicsPathItem_Contains(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QGraphicsPathItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsPathItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QGraphicsPathItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsPathItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsPathItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsPathItem_SetPath(QtObjectPtr ptr, QtObjectPtr path){
	static_cast<QGraphicsPathItem*>(ptr)->setPath(*static_cast<QPainterPath*>(path));
}

int QGraphicsPathItem_Type(QtObjectPtr ptr){
	return static_cast<QGraphicsPathItem*>(ptr)->type();
}

void QGraphicsPathItem_DestroyQGraphicsPathItem(QtObjectPtr ptr){
	static_cast<QGraphicsPathItem*>(ptr)->~QGraphicsPathItem();
}

