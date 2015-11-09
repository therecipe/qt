#include "qstyleoptionviewitem.h"
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionViewItem>
#include "_cgo_export.h"

class MyQStyleOptionViewItem: public QStyleOptionViewItem {
public:
};

void* QStyleOptionViewItem_NewQStyleOptionViewItem(){
	return new QStyleOptionViewItem();
}

void* QStyleOptionViewItem_NewQStyleOptionViewItem2(void* other){
	return new QStyleOptionViewItem(*static_cast<QStyleOptionViewItem*>(other));
}

