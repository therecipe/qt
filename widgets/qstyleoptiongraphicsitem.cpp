#include "qstyleoptiongraphicsitem.h"
#include <QStyleOption>
#include <QTransform>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOptionGraphicsItem>
#include "_cgo_export.h"

class MyQStyleOptionGraphicsItem: public QStyleOptionGraphicsItem {
public:
};

void* QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem(){
	return new QStyleOptionGraphicsItem();
}

void* QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem2(void* other){
	return new QStyleOptionGraphicsItem(*static_cast<QStyleOptionGraphicsItem*>(other));
}

double QStyleOptionGraphicsItem_QStyleOptionGraphicsItem_LevelOfDetailFromTransform(void* worldTransform){
	return static_cast<double>(QStyleOptionGraphicsItem::levelOfDetailFromTransform(*static_cast<QTransform*>(worldTransform)));
}

