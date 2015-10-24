#include "qgraphicsellipseitem.h"
#include <QUrl>
#include <QVariant>
#include <QModelIndex>
#include <QPoint>
#include <QPainter>
#include <QStyleOption>
#include <QWidget>
#include <QGraphicsItem>
#include <QStyleOptionGraphicsItem>
#include <QString>
#include <QRect>
#include <QRectF>
#include <QPointF>
#include <QStyle>
#include <QGraphicsEllipseItem>
#include "_cgo_export.h"

class MyQGraphicsEllipseItem: public QGraphicsEllipseItem {
public:
};

int QGraphicsEllipseItem_Contains(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QGraphicsEllipseItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsEllipseItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QGraphicsEllipseItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsEllipseItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsEllipseItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsEllipseItem_SetRect(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QGraphicsEllipseItem*>(ptr)->setRect(*static_cast<QRectF*>(rect));
}

void QGraphicsEllipseItem_SetSpanAngle(QtObjectPtr ptr, int angle){
	static_cast<QGraphicsEllipseItem*>(ptr)->setSpanAngle(angle);
}

void QGraphicsEllipseItem_SetStartAngle(QtObjectPtr ptr, int angle){
	static_cast<QGraphicsEllipseItem*>(ptr)->setStartAngle(angle);
}

int QGraphicsEllipseItem_SpanAngle(QtObjectPtr ptr){
	return static_cast<QGraphicsEllipseItem*>(ptr)->spanAngle();
}

int QGraphicsEllipseItem_StartAngle(QtObjectPtr ptr){
	return static_cast<QGraphicsEllipseItem*>(ptr)->startAngle();
}

int QGraphicsEllipseItem_Type(QtObjectPtr ptr){
	return static_cast<QGraphicsEllipseItem*>(ptr)->type();
}

void QGraphicsEllipseItem_DestroyQGraphicsEllipseItem(QtObjectPtr ptr){
	static_cast<QGraphicsEllipseItem*>(ptr)->~QGraphicsEllipseItem();
}

