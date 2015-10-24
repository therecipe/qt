#include "qtextlist.h"
#include <QList>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QList>
#include <QTextListFormat>
#include <QTextBlock>
#include <QTextList>
#include "_cgo_export.h"

class MyQTextList: public QTextList {
public:
};

int QTextList_ItemNumber(QtObjectPtr ptr, QtObjectPtr block){
	return static_cast<QTextList*>(ptr)->itemNumber(*static_cast<QTextBlock*>(block));
}

char* QTextList_ItemText(QtObjectPtr ptr, QtObjectPtr block){
	return static_cast<QTextList*>(ptr)->itemText(*static_cast<QTextBlock*>(block)).toUtf8().data();
}

void QTextList_Add(QtObjectPtr ptr, QtObjectPtr block){
	static_cast<QTextList*>(ptr)->add(*static_cast<QTextBlock*>(block));
}

int QTextList_Count(QtObjectPtr ptr){
	return static_cast<QTextList*>(ptr)->count();
}

void QTextList_RemoveItem(QtObjectPtr ptr, int i){
	static_cast<QTextList*>(ptr)->removeItem(i);
}

void QTextList_SetFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QTextList*>(ptr)->setFormat(*static_cast<QTextListFormat*>(format));
}

