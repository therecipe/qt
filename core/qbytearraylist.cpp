#include "qbytearraylist.h"
#include <QList>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QList>
#include <QByteArrayList>
#include "_cgo_export.h"

class MyQByteArrayList: public QByteArrayList {
public:
};

void* QByteArrayList_Join(void* ptr){
	return new QByteArray(static_cast<QByteArrayList*>(ptr)->join());
}

void* QByteArrayList_Join3(void* ptr, char* separator){
	return new QByteArray(static_cast<QByteArrayList*>(ptr)->join(*separator));
}

void* QByteArrayList_Join2(void* ptr, void* separator){
	return new QByteArray(static_cast<QByteArrayList*>(ptr)->join(*static_cast<QByteArray*>(separator)));
}

