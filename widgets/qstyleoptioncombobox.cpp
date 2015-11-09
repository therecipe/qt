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

void* QStyleOptionComboBox_NewQStyleOptionComboBox(){
	return new QStyleOptionComboBox();
}

void* QStyleOptionComboBox_NewQStyleOptionComboBox2(void* other){
	return new QStyleOptionComboBox(*static_cast<QStyleOptionComboBox*>(other));
}

