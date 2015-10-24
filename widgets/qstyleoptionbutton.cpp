#include "qstyleoptionbutton.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionButton>
#include "_cgo_export.h"

class MyQStyleOptionButton: public QStyleOptionButton {
public:
};

QtObjectPtr QStyleOptionButton_NewQStyleOptionButton(){
	return new QStyleOptionButton();
}

QtObjectPtr QStyleOptionButton_NewQStyleOptionButton2(QtObjectPtr other){
	return new QStyleOptionButton(*static_cast<QStyleOptionButton*>(other));
}

