#include "qgraphicsrectitem.h"
#include <QString>
#include <QWidget>
#include <QStyle>
#include <QPointF>
#include <QPoint>
#include <QPainter>
#include <QRect>
#include <QRectF>
#include <QVariant>
#include <QGraphicsItem>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOptionGraphicsItem>
#include <QStyleOption>
#include <QGraphicsRectItem>
#include "_cgo_export.h"

class MyQGraphicsRectItem: public QGraphicsRectItem {
public:
};

void QGraphicsRectItem_SetRect(void* ptr, void* rectangle){
	static_cast<QGraphicsRectItem*>(ptr)->setRect(*static_cast<QRectF*>(rectangle));
}

int QGraphicsRectItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsRectItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsRectItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsRectItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsRectItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsRectItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsRectItem_SetRect2(void* ptr, double x, double y, double width, double height){
	static_cast<QGraphicsRectItem*>(ptr)->setRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(width), static_cast<qreal>(height));
}

int QGraphicsRectItem_Type(void* ptr){
	return static_cast<QGraphicsRectItem*>(ptr)->type();
}

void QGraphicsRectItem_DestroyQGraphicsRectItem(void* ptr){
	static_cast<QGraphicsRectItem*>(ptr)->~QGraphicsRectItem();
}

