#include "qstyleoptiontoolbar.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QStyleOptionToolBar>
#include "_cgo_export.h"

class MyQStyleOptionToolBar: public QStyleOptionToolBar {
public:
};

void* QStyleOptionToolBar_NewQStyleOptionToolBar(){
	return new QStyleOptionToolBar();
}

void* QStyleOptionToolBar_NewQStyleOptionToolBar2(void* other){
	return new QStyleOptionToolBar(*static_cast<QStyleOptionToolBar*>(other));
}

