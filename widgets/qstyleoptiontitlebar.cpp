#include "qstyleoptiontitlebar.h"
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOptionTitleBar>
#include "_cgo_export.h"

class MyQStyleOptionTitleBar: public QStyleOptionTitleBar {
public:
};

void* QStyleOptionTitleBar_NewQStyleOptionTitleBar(){
	return new QStyleOptionTitleBar();
}

void* QStyleOptionTitleBar_NewQStyleOptionTitleBar2(void* other){
	return new QStyleOptionTitleBar(*static_cast<QStyleOptionTitleBar*>(other));
}

