#include "qstyleoptionspinbox.h"
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
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

