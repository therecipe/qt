#include "qstyleoptiontabbarbase.h"
#include <QStyleOption>
#include <QStyleOptionTab>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOptionTabBarBase>
#include "_cgo_export.h"

class MyQStyleOptionTabBarBase: public QStyleOptionTabBarBase {
public:
};

void* QStyleOptionTabBarBase_NewQStyleOptionTabBarBase(){
	return new QStyleOptionTabBarBase();
}

void* QStyleOptionTabBarBase_NewQStyleOptionTabBarBase2(void* other){
	return new QStyleOptionTabBarBase(*static_cast<QStyleOptionTabBarBase*>(other));
}

