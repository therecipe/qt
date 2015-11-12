#include "qtextlist.h"
#include <QList>
#include <QModelIndex>
#include <QTextListFormat>
#include <QTextBlock>
#include <QList>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTextList>
#include "_cgo_export.h"

class MyQTextList: public QTextList {
public:
};

int QTextList_ItemNumber(void* ptr, void* block){
	return static_cast<QTextList*>(ptr)->itemNumber(*static_cast<QTextBlock*>(block));
}

char* QTextList_ItemText(void* ptr, void* block){
	return static_cast<QTextList*>(ptr)->itemText(*static_cast<QTextBlock*>(block)).toUtf8().data();
}

void QTextList_Add(void* ptr, void* block){
	static_cast<QTextList*>(ptr)->add(*static_cast<QTextBlock*>(block));
}

int QTextList_Count(void* ptr){
	return static_cast<QTextList*>(ptr)->count();
}

void QTextList_RemoveItem(void* ptr, int i){
	static_cast<QTextList*>(ptr)->removeItem(i);
}

void QTextList_SetFormat(void* ptr, void* format){
	static_cast<QTextList*>(ptr)->setFormat(*static_cast<QTextListFormat*>(format));
}

