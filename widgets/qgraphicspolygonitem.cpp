#include "qgraphicspolygonitem.h"
#include <QPolygonF>
#include <QStyleOptionGraphicsItem>
#include <QPoint>
#include <QVariant>
#include <QModelIndex>
#include <QPointF>
#include <QPolygon>
#include <QString>
#include <QStyle>
#include <QStyleOption>
#include <QPainter>
#include <QUrl>
#include <QWidget>
#include <QGraphicsItem>
#include <QGraphicsPolygonItem>
#include "_cgo_export.h"

class MyQGraphicsPolygonItem: public QGraphicsPolygonItem {
public:
};

int QGraphicsPolygonItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsPolygonItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsPolygonItem_FillRule(void* ptr){
	return static_cast<QGraphicsPolygonItem*>(ptr)->fillRule();
}

int QGraphicsPolygonItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsPolygonItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsPolygonItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsPolygonItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsPolygonItem_SetFillRule(void* ptr, int rule){
	static_cast<QGraphicsPolygonItem*>(ptr)->setFillRule(static_cast<Qt::FillRule>(rule));
}

void QGraphicsPolygonItem_SetPolygon(void* ptr, void* polygon){
	static_cast<QGraphicsPolygonItem*>(ptr)->setPolygon(*static_cast<QPolygonF*>(polygon));
}

int QGraphicsPolygonItem_Type(void* ptr){
	return static_cast<QGraphicsPolygonItem*>(ptr)->type();
}

void QGraphicsPolygonItem_DestroyQGraphicsPolygonItem(void* ptr){
	static_cast<QGraphicsPolygonItem*>(ptr)->~QGraphicsPolygonItem();
}

