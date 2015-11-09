#include "qstyleoptiontoolbutton.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QStyleOptionToolButton>
#include "_cgo_export.h"

class MyQStyleOptionToolButton: public QStyleOptionToolButton {
public:
};

void* QStyleOptionToolButton_NewQStyleOptionToolButton(){
	return new QStyleOptionToolButton();
}

void* QStyleOptionToolButton_NewQStyleOptionToolButton2(void* other){
	return new QStyleOptionToolButton(*static_cast<QStyleOptionToolButton*>(other));
}

