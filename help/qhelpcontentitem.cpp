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

QtObjectPtr QHelpContentItem_Child(QtObjectPtr ptr, int row){
	return static_cast<QHelpContentItem*>(ptr)->child(row);
}

int QHelpContentItem_ChildCount(QtObjectPtr ptr){
	return static_cast<QHelpContentItem*>(ptr)->childCount();
}

int QHelpContentItem_ChildPosition(QtObjectPtr ptr, QtObjectPtr child){
	return static_cast<QHelpContentItem*>(ptr)->childPosition(static_cast<QHelpContentItem*>(child));
}

QtObjectPtr QHelpContentItem_Parent(QtObjectPtr ptr){
	return static_cast<QHelpContentItem*>(ptr)->parent();
}

int QHelpContentItem_Row(QtObjectPtr ptr){
	return static_cast<QHelpContentItem*>(ptr)->row();
}

char* QHelpContentItem_Title(QtObjectPtr ptr){
	return static_cast<QHelpContentItem*>(ptr)->title().toUtf8().data();
}

char* QHelpContentItem_Url(QtObjectPtr ptr){
	return static_cast<QHelpContentItem*>(ptr)->url().toString().toUtf8().data();
}

void QHelpContentItem_DestroyQHelpContentItem(QtObjectPtr ptr){
	static_cast<QHelpContentItem*>(ptr)->~QHelpContentItem();
}

