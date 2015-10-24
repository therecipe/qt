#include "qstyleoptiongraphicsitem.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QStyleOptionGraphicsItem>
#include "_cgo_export.h"

class MyQStyleOptionGraphicsItem: public QStyleOptionGraphicsItem {
public:
};

QtObjectPtr QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem(){
	return new QStyleOptionGraphicsItem();
}

QtObjectPtr QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem2(QtObjectPtr other){
	return new QStyleOptionGraphicsItem(*static_cast<QStyleOptionGraphicsItem*>(other));
}

