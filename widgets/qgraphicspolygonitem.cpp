#include "qgraphicspolygonitem.h"
#include <QString>
#include <QPolygonF>
#include <QStyleOptionGraphicsItem>
#include <QPainter>
#include <QUrl>
#include <QGraphicsItem>
#include <QPolygon>
#include <QStyle>
#include <QPoint>
#include <QVariant>
#include <QModelIndex>
#include <QWidget>
#include <QStyleOption>
#include <QPointF>
#include <QGraphicsPolygonItem>
#include "_cgo_export.h"

class MyQGraphicsPolygonItem: public QGraphicsPolygonItem {
public:
};

int QGraphicsPolygonItem_Contains(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QGraphicsPolygonItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsPolygonItem_FillRule(QtObjectPtr ptr){
	return static_cast<QGraphicsPolygonItem*>(ptr)->fillRule();
}

int QGraphicsPolygonItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QGraphicsPolygonItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsPolygonItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsPolygonItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsPolygonItem_SetFillRule(QtObjectPtr ptr, int rule){
	static_cast<QGraphicsPolygonItem*>(ptr)->setFillRule(static_cast<Qt::FillRule>(rule));
}

void QGraphicsPolygonItem_SetPolygon(QtObjectPtr ptr, QtObjectPtr polygon){
	static_cast<QGraphicsPolygonItem*>(ptr)->setPolygon(*static_cast<QPolygonF*>(polygon));
}

int QGraphicsPolygonItem_Type(QtObjectPtr ptr){
	return static_cast<QGraphicsPolygonItem*>(ptr)->type();
}

void QGraphicsPolygonItem_DestroyQGraphicsPolygonItem(QtObjectPtr ptr){
	static_cast<QGraphicsPolygonItem*>(ptr)->~QGraphicsPolygonItem();
}

