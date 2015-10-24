#include "qstyleoptiontoolbox.h"
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionToolBox>
#include "_cgo_export.h"

class MyQStyleOptionToolBox: public QStyleOptionToolBox {
public:
};

QtObjectPtr QStyleOptionToolBox_NewQStyleOptionToolBox(){
	return new QStyleOptionToolBox();
}

QtObjectPtr QStyleOptionToolBox_NewQStyleOptionToolBox2(QtObjectPtr other){
	return new QStyleOptionToolBox(*static_cast<QStyleOptionToolBox*>(other));
}

