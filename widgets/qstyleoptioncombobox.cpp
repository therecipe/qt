#include "qstyleoptioncombobox.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionComboBox>
#include "_cgo_export.h"

class MyQStyleOptionComboBox: public QStyleOptionComboBox {
public:
};

QtObjectPtr QStyleOptionComboBox_NewQStyleOptionComboBox(){
	return new QStyleOptionComboBox();
}

QtObjectPtr QStyleOptionComboBox_NewQStyleOptionComboBox2(QtObjectPtr other){
	return new QStyleOptionComboBox(*static_cast<QStyleOptionComboBox*>(other));
}

