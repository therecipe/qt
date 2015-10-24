#include "qstyleoptiontoolbutton.h"
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionToolButton>
#include "_cgo_export.h"

class MyQStyleOptionToolButton: public QStyleOptionToolButton {
public:
};

QtObjectPtr QStyleOptionToolButton_NewQStyleOptionToolButton(){
	return new QStyleOptionToolButton();
}

QtObjectPtr QStyleOptionToolButton_NewQStyleOptionToolButton2(QtObjectPtr other){
	return new QStyleOptionToolButton(*static_cast<QStyleOptionToolButton*>(other));
}

