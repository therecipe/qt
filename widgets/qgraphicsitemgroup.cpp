#include "qgraphicsitemgroup.h"
#include <QString>
#include <QVariant>
#include <QStyleOption>
#include <QWidget>
#include <QStyleOptionGraphicsItem>
#include <QStyle>
#include <QPainter>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsItem>
#include <QGraphicsItemGroup>
#include "_cgo_export.h"

class MyQGraphicsItemGroup: public QGraphicsItemGroup {
public:
};

QtObjectPtr QGraphicsItemGroup_NewQGraphicsItemGroup(QtObjectPtr parent){
	return new QGraphicsItemGroup(static_cast<QGraphicsItem*>(parent));
}

void QGraphicsItemGroup_AddToGroup(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QGraphicsItemGroup*>(ptr)->addToGroup(static_cast<QGraphicsItem*>(item));
}

int QGraphicsItemGroup_IsObscuredBy(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QGraphicsItemGroup*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsItemGroup_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr option, QtObjectPtr widget){
	static_cast<QGraphicsItemGroup*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsItemGroup_RemoveFromGroup(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QGraphicsItemGroup*>(ptr)->removeFromGroup(static_cast<QGraphicsItem*>(item));
}

int QGraphicsItemGroup_Type(QtObjectPtr ptr){
	return static_cast<QGraphicsItemGroup*>(ptr)->type();
}

void QGraphicsItemGroup_DestroyQGraphicsItemGroup(QtObjectPtr ptr){
	static_cast<QGraphicsItemGroup*>(ptr)->~QGraphicsItemGroup();
}

