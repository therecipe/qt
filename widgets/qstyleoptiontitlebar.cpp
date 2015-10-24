#include "qstyleoptiontitlebar.h"
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionTitleBar>
#include "_cgo_export.h"

class MyQStyleOptionTitleBar: public QStyleOptionTitleBar {
public:
};

QtObjectPtr QStyleOptionTitleBar_NewQStyleOptionTitleBar(){
	return new QStyleOptionTitleBar();
}

QtObjectPtr QStyleOptionTitleBar_NewQStyleOptionTitleBar2(QtObjectPtr other){
	return new QStyleOptionTitleBar(*static_cast<QStyleOptionTitleBar*>(other));
}

