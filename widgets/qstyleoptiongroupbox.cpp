#include "qstyleoptiongroupbox.h"
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QStyleOptionGroupBox>
#include "_cgo_export.h"

class MyQStyleOptionGroupBox: public QStyleOptionGroupBox {
public:
};

void* QStyleOptionGroupBox_NewQStyleOptionGroupBox(){
	return new QStyleOptionGroupBox();
}

void* QStyleOptionGroupBox_NewQStyleOptionGroupBox2(void* other){
	return new QStyleOptionGroupBox(*static_cast<QStyleOptionGroupBox*>(other));
}

