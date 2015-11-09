#include "qstyleoptionspinbox.h"
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QStyleOptionSpinBox>
#include "_cgo_export.h"

class MyQStyleOptionSpinBox: public QStyleOptionSpinBox {
public:
};

void* QStyleOptionSpinBox_NewQStyleOptionSpinBox(){
	return new QStyleOptionSpinBox();
}

void* QStyleOptionSpinBox_NewQStyleOptionSpinBox2(void* other){
	return new QStyleOptionSpinBox(*static_cast<QStyleOptionSpinBox*>(other));
}

