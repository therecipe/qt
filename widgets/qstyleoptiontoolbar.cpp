#include "qstyleoptiontoolbar.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
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

