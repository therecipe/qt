#include "qgraphicsrectitem.h"
#include <QRect>
#include <QWidget>
#include <QPainter>
#include <QPoint>
#include <QPointF>
#include <QModelIndex>
#include <QStyleOptionGraphicsItem>
#include <QGraphicsItem>
#include <QStyle>
#include <QVariant>
#include <QUrl>
#include <QStyleOption>
#include <QString>
#include <QRectF>
#include <QGraphicsRectItem>
#include "_cgo_export.h"

class MyQGraphicsRectItem: public QGraphicsRectItem {
public:
};

void QGraphicsRectItem_SetRect(QtObjectPtr ptr, QtObjectPtr rectangle){
	static_cast<QGraphicsRectItem*>(ptr)->setRect(*static_cast<QRectF*>(rectangle));
}

int QGraphicsRectItem_Contains(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QGraphicsRectItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsRectItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QGraphicsRectItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsRectItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsRectItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

int QGraphicsRectItem_Type(QtObjectPtr ptr){
	return static_cast<QGraphicsRectItem*>(ptr)->type();
}

void QGraphicsRectItem_DestroyQGraphicsRectItem(QtObjectPtr ptr){
	static_cast<QGraphicsRectItem*>(ptr)->~QGraphicsRectItem();
}

