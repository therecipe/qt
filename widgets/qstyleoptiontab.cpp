#include "qstyleoptiontab.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionTab>
#include "_cgo_export.h"

class MyQStyleOptionTab: public QStyleOptionTab {
public:
};

QtObjectPtr QStyleOptionTab_NewQStyleOptionTab(){
	return new QStyleOptionTab();
}

QtObjectPtr QStyleOptionTab_NewQStyleOptionTab2(QtObjectPtr other){
	return new QStyleOptionTab(*static_cast<QStyleOptionTab*>(other));
}

