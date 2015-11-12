#include "qstyleoptiontoolbutton.h"
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
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

