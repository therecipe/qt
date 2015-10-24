#include "qstyleoptiontoolbar.h"
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionToolBar>
#include "_cgo_export.h"

class MyQStyleOptionToolBar: public QStyleOptionToolBar {
public:
};

QtObjectPtr QStyleOptionToolBar_NewQStyleOptionToolBar(){
	return new QStyleOptionToolBar();
}

QtObjectPtr QStyleOptionToolBar_NewQStyleOptionToolBar2(QtObjectPtr other){
	return new QStyleOptionToolBar(*static_cast<QStyleOptionToolBar*>(other));
}

