#include "qstyleoptiontabbarbase.h"
#include <QModelIndex>
#include <QStyle>
#include <QStyleOptionTab>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionTabBarBase>
#include "_cgo_export.h"

class MyQStyleOptionTabBarBase: public QStyleOptionTabBarBase {
public:
};

QtObjectPtr QStyleOptionTabBarBase_NewQStyleOptionTabBarBase(){
	return new QStyleOptionTabBarBase();
}

QtObjectPtr QStyleOptionTabBarBase_NewQStyleOptionTabBarBase2(QtObjectPtr other){
	return new QStyleOptionTabBarBase(*static_cast<QStyleOptionTabBarBase*>(other));
}

