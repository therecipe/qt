#include "qstyleoptiontab.h"
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOptionTab>
#include "_cgo_export.h"

class MyQStyleOptionTab: public QStyleOptionTab {
public:
};

void* QStyleOptionTab_NewQStyleOptionTab(){
	return new QStyleOptionTab();
}

void* QStyleOptionTab_NewQStyleOptionTab2(void* other){
	return new QStyleOptionTab(*static_cast<QStyleOptionTab*>(other));
}

