#include "qstyleoptionmenuitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QStyleOptionMenuItem>
#include "_cgo_export.h"

class MyQStyleOptionMenuItem: public QStyleOptionMenuItem {
public:
};

void* QStyleOptionMenuItem_NewQStyleOptionMenuItem(){
	return new QStyleOptionMenuItem();
}

void* QStyleOptionMenuItem_NewQStyleOptionMenuItem2(void* other){
	return new QStyleOptionMenuItem(*static_cast<QStyleOptionMenuItem*>(other));
}

