#include "qstyleoptionspinbox.h"
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QStyleOptionSpinBox>
#include "_cgo_export.h"

class MyQStyleOptionSpinBox: public QStyleOptionSpinBox {
public:
};

QtObjectPtr QStyleOptionSpinBox_NewQStyleOptionSpinBox(){
	return new QStyleOptionSpinBox();
}

QtObjectPtr QStyleOptionSpinBox_NewQStyleOptionSpinBox2(QtObjectPtr other){
	return new QStyleOptionSpinBox(*static_cast<QStyleOptionSpinBox*>(other));
}

