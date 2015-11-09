#include "qtextitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextItem>
#include "_cgo_export.h"

class MyQTextItem: public QTextItem {
public:
};

double QTextItem_Ascent(void* ptr){
	return static_cast<double>(static_cast<QTextItem*>(ptr)->ascent());
}

double QTextItem_Descent(void* ptr){
	return static_cast<double>(static_cast<QTextItem*>(ptr)->descent());
}

int QTextItem_RenderFlags(void* ptr){
	return static_cast<QTextItem*>(ptr)->renderFlags();
}

char* QTextItem_Text(void* ptr){
	return static_cast<QTextItem*>(ptr)->text().toUtf8().data();
}

double QTextItem_Width(void* ptr){
	return static_cast<double>(static_cast<QTextItem*>(ptr)->width());
}

