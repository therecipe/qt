#include "qgraphicssimpletextitem.h"
#include <QStyleOption>
#include <QWidget>
#include <QStyleOptionGraphicsItem>
#include <QPoint>
#include <QString>
#include <QUrl>
#include <QFont>
#include <QPainter>
#include <QPointF>
#include <QVariant>
#include <QModelIndex>
#include <QStyle>
#include <QGraphicsItem>
#include <QGraphicsSimpleTextItem>
#include "_cgo_export.h"

class MyQGraphicsSimpleTextItem: public QGraphicsSimpleTextItem {
public:
};

int QGraphicsSimpleTextItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsSimpleTextItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsSimpleTextItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsSimpleTextItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsSimpleTextItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsSimpleTextItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsSimpleTextItem_SetFont(void* ptr, void* font){
	static_cast<QGraphicsSimpleTextItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsSimpleTextItem_SetText(void* ptr, char* text){
	static_cast<QGraphicsSimpleTextItem*>(ptr)->setText(QString(text));
}

char* QGraphicsSimpleTextItem_Text(void* ptr){
	return static_cast<QGraphicsSimpleTextItem*>(ptr)->text().toUtf8().data();
}

int QGraphicsSimpleTextItem_Type(void* ptr){
	return static_cast<QGraphicsSimpleTextItem*>(ptr)->type();
}

void QGraphicsSimpleTextItem_DestroyQGraphicsSimpleTextItem(void* ptr){
	static_cast<QGraphicsSimpleTextItem*>(ptr)->~QGraphicsSimpleTextItem();
}

