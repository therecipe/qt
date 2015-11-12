#include "qgraphicsitemgroup.h"
#include <QVariant>
#include <QUrl>
#include <QPainter>
#include <QStyleOption>
#include <QString>
#include <QStyle>
#include <QWidget>
#include <QStyleOptionGraphicsItem>
#include <QGraphicsItem>
#include <QModelIndex>
#include <QGraphicsItemGroup>
#include "_cgo_export.h"

class MyQGraphicsItemGroup: public QGraphicsItemGroup {
public:
};

void* QGraphicsItemGroup_NewQGraphicsItemGroup(void* parent){
	return new QGraphicsItemGroup(static_cast<QGraphicsItem*>(parent));
}

void QGraphicsItemGroup_AddToGroup(void* ptr, void* item){
	static_cast<QGraphicsItemGroup*>(ptr)->addToGroup(static_cast<QGraphicsItem*>(item));
}

int QGraphicsItemGroup_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsItemGroup*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsItemGroup_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsItemGroup*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsItemGroup_RemoveFromGroup(void* ptr, void* item){
	static_cast<QGraphicsItemGroup*>(ptr)->removeFromGroup(static_cast<QGraphicsItem*>(item));
}

int QGraphicsItemGroup_Type(void* ptr){
	return static_cast<QGraphicsItemGroup*>(ptr)->type();
}

void QGraphicsItemGroup_DestroyQGraphicsItemGroup(void* ptr){
	static_cast<QGraphicsItemGroup*>(ptr)->~QGraphicsItemGroup();
}

