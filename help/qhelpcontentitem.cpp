#include "qhelpcontentitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHelpContentItem>
#include "_cgo_export.h"

class MyQHelpContentItem: public QHelpContentItem {
public:
};

void* QHelpContentItem_Child(void* ptr, int row){
	return static_cast<QHelpContentItem*>(ptr)->child(row);
}

int QHelpContentItem_ChildCount(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->childCount();
}

int QHelpContentItem_ChildPosition(void* ptr, void* child){
	return static_cast<QHelpContentItem*>(ptr)->childPosition(static_cast<QHelpContentItem*>(child));
}

void* QHelpContentItem_Parent(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->parent();
}

int QHelpContentItem_Row(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->row();
}

char* QHelpContentItem_Title(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->title().toUtf8().data();
}

void QHelpContentItem_DestroyQHelpContentItem(void* ptr){
	static_cast<QHelpContentItem*>(ptr)->~QHelpContentItem();
}

