#include "qgraphicsellipseitem.h"
#include <QUrl>
#include <QStyleOptionGraphicsItem>
#include <QVariant>
#include <QRect>
#include <QPoint>
#include <QWidget>
#include <QRectF>
#include <QString>
#include <QStyle>
#include <QGraphicsItem>
#include <QStyleOption>
#include <QModelIndex>
#include <QPainter>
#include <QPointF>
#include <QGraphicsEllipseItem>
#include "_cgo_export.h"

class MyQGraphicsEllipseItem: public QGraphicsEllipseItem {
public:
};

int QGraphicsEllipseItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsEllipseItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsEllipseItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsEllipseItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsEllipseItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsEllipseItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsEllipseItem_SetRect(void* ptr, void* rect){
	static_cast<QGraphicsEllipseItem*>(ptr)->setRect(*static_cast<QRectF*>(rect));
}

void QGraphicsEllipseItem_SetRect2(void* ptr, double x, double y, double width, double height){
	static_cast<QGraphicsEllipseItem*>(ptr)->setRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(width), static_cast<qreal>(height));
}

void QGraphicsEllipseItem_SetSpanAngle(void* ptr, int angle){
	static_cast<QGraphicsEllipseItem*>(ptr)->setSpanAngle(angle);
}

void QGraphicsEllipseItem_SetStartAngle(void* ptr, int angle){
	static_cast<QGraphicsEllipseItem*>(ptr)->setStartAngle(angle);
}

int QGraphicsEllipseItem_SpanAngle(void* ptr){
	return static_cast<QGraphicsEllipseItem*>(ptr)->spanAngle();
}

int QGraphicsEllipseItem_StartAngle(void* ptr){
	return static_cast<QGraphicsEllipseItem*>(ptr)->startAngle();
}

int QGraphicsEllipseItem_Type(void* ptr){
	return static_cast<QGraphicsEllipseItem*>(ptr)->type();
}

void QGraphicsEllipseItem_DestroyQGraphicsEllipseItem(void* ptr){
	static_cast<QGraphicsEllipseItem*>(ptr)->~QGraphicsEllipseItem();
}

