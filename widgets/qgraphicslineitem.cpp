#include "qgraphicslineitem.h"
#include <QVariant>
#include <QPen>
#include <QPointF>
#include <QString>
#include <QLine>
#include <QStyleOptionGraphicsItem>
#include <QPoint>
#include <QStyle>
#include <QStyleOption>
#include <QLineF>
#include <QGraphicsItem>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QPainter>
#include <QGraphicsLineItem>
#include "_cgo_export.h"

class MyQGraphicsLineItem: public QGraphicsLineItem {
public:
};

void* QGraphicsLineItem_NewQGraphicsLineItem(void* parent){
	return new QGraphicsLineItem(static_cast<QGraphicsItem*>(parent));
}

void* QGraphicsLineItem_NewQGraphicsLineItem2(void* line, void* parent){
	return new QGraphicsLineItem(*static_cast<QLineF*>(line), static_cast<QGraphicsItem*>(parent));
}

void* QGraphicsLineItem_NewQGraphicsLineItem3(double x1, double y1, double x2, double y2, void* parent){
	return new QGraphicsLineItem(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2), static_cast<QGraphicsItem*>(parent));
}

int QGraphicsLineItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsLineItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsLineItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsLineItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsLineItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsLineItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsLineItem_SetLine(void* ptr, void* line){
	static_cast<QGraphicsLineItem*>(ptr)->setLine(*static_cast<QLineF*>(line));
}

void QGraphicsLineItem_SetLine2(void* ptr, double x1, double y1, double x2, double y2){
	static_cast<QGraphicsLineItem*>(ptr)->setLine(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2));
}

void QGraphicsLineItem_SetPen(void* ptr, void* pen){
	static_cast<QGraphicsLineItem*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

int QGraphicsLineItem_Type(void* ptr){
	return static_cast<QGraphicsLineItem*>(ptr)->type();
}

void QGraphicsLineItem_DestroyQGraphicsLineItem(void* ptr){
	static_cast<QGraphicsLineItem*>(ptr)->~QGraphicsLineItem();
}

