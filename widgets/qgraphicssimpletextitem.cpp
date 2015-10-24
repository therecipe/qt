#include "qgraphicssimpletextitem.h"
#include <QString>
#include <QStyleOption>
#include <QFont>
#include <QWidget>
#include <QGraphicsItem>
#include <QStyleOptionGraphicsItem>
#include <QStyle>
#include <QPoint>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPointF>
#include <QPainter>
#include <QGraphicsSimpleTextItem>
#include "_cgo_export.h"

class MyQGraphicsSimpleTextItem: public QGraphicsSimpleTextItem {
public:
};

int QGraphicsSimpleTextItem_Contains(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QGraphicsSimpleTextItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsSimpleTextItem_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QGraphicsSimpleTextItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsSimpleTextItem_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsSimpleTextItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsSimpleTextItem_SetFont(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QGraphicsSimpleTextItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsSimpleTextItem_SetText(QtObjectPtr ptr, char* text){
	static_cast<QGraphicsSimpleTextItem*>(ptr)->setText(QString(text));
}

char* QGraphicsSimpleTextItem_Text(QtObjectPtr ptr){
	return static_cast<QGraphicsSimpleTextItem*>(ptr)->text().toUtf8().data();
}

int QGraphicsSimpleTextItem_Type(QtObjectPtr ptr){
	return static_cast<QGraphicsSimpleTextItem*>(ptr)->type();
}

void QGraphicsSimpleTextItem_DestroyQGraphicsSimpleTextItem(QtObjectPtr ptr){
	static_cast<QGraphicsSimpleTextItem*>(ptr)->~QGraphicsSimpleTextItem();
}

