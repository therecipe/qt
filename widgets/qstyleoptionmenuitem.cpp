#include "qstyleoptionmenuitem.h"
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOptionMenuItem>
#include "_cgo_export.h"

class MyQStyleOptionMenuItem: public QStyleOptionMenuItem {
public:
};

QtObjectPtr QStyleOptionMenuItem_NewQStyleOptionMenuItem(){
	return new QStyleOptionMenuItem();
}

QtObjectPtr QStyleOptionMenuItem_NewQStyleOptionMenuItem2(QtObjectPtr other){
	return new QStyleOptionMenuItem(*static_cast<QStyleOptionMenuItem*>(other));
}

