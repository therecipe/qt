#include "qstyleoptiongroupbox.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QStyleOptionGroupBox>
#include "_cgo_export.h"

class MyQStyleOptionGroupBox: public QStyleOptionGroupBox {
public:
};

QtObjectPtr QStyleOptionGroupBox_NewQStyleOptionGroupBox(){
	return new QStyleOptionGroupBox();
}

QtObjectPtr QStyleOptionGroupBox_NewQStyleOptionGroupBox2(QtObjectPtr other){
	return new QStyleOptionGroupBox(*static_cast<QStyleOptionGroupBox*>(other));
}

