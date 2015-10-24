#include "qgraphicslineitem.h"
#include <QUrl>
#include <QPainter>
#include <QPen>
#include <QLine>
#include <QStyleOptionGraphicsItem>
#include <QString>
#include <QVariant>
#include <QStyle>
#include <QPoint>
#include <QModelIndex>
#include <QGraphicsItem>
#include <QLineF>
#include <QStyleOption>
#include <QPointF>
#include <QWidget>
#include <QGraphicsLineItem>
#include "_cgo_export.h"

class MyQGraphicsLineItem: public QGraphicsLineItem {
public:
};

QtObjectPtr QGraphicsLineItem_NewQGraphicsLineItem(QtObjectPtr parent){
	return new QGraphicsLineItem(static_cast<QGraphicsItem*>(parent));
}

QtObjectPtr QGraphicsLineItem_NewQGraphicsLineItem2(QtObjectPtr line, QtObjectPtr parent){
	return new QGraphicsLineItem(*static_cast<QLineF*>(line), static_cast<QGraphicsItem*>(parent));
}

int QGraphicsLineItem_Contains(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QGraphicsLineItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsLineItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QGraphicsLineItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsLineItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsLineItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsLineItem_SetLine(QtObjectPtr ptr, QtObjectPtr line){
	static_cast<QGraphicsLineItem*>(ptr)->setLine(*static_cast<QLineF*>(line));
}

void QGraphicsLineItem_SetPen(QtObjectPtr ptr, QtObjectPtr pen){
	static_cast<QGraphicsLineItem*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

int QGraphicsLineItem_Type(QtObjectPtr ptr){
	return static_cast<QGraphicsLineItem*>(ptr)->type();
}

void QGraphicsLineItem_DestroyQGraphicsLineItem(QtObjectPtr ptr){
	static_cast<QGraphicsLineItem*>(ptr)->~QGraphicsLineItem();
}

