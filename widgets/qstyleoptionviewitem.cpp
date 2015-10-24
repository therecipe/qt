#include "qstyleoptionviewitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionViewItem>
#include "_cgo_export.h"

class MyQStyleOptionViewItem: public QStyleOptionViewItem {
public:
};

QtObjectPtr QStyleOptionViewItem_NewQStyleOptionViewItem(){
	return new QStyleOptionViewItem();
}

QtObjectPtr QStyleOptionViewItem_NewQStyleOptionViewItem2(QtObjectPtr other){
	return new QStyleOptionViewItem(*static_cast<QStyleOptionViewItem*>(other));
}

