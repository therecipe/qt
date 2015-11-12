#include "qstyleoptiontoolbox.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionToolBox>
#include "_cgo_export.h"

class MyQStyleOptionToolBox: public QStyleOptionToolBox {
public:
};

void* QStyleOptionToolBox_NewQStyleOptionToolBox(){
	return new QStyleOptionToolBox();
}

void* QStyleOptionToolBox_NewQStyleOptionToolBox2(void* other){
	return new QStyleOptionToolBox(*static_cast<QStyleOptionToolBox*>(other));
}

