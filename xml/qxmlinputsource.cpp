#include "qxmlinputsource.h"
#include <QModelIndex>
#include <QIODevice>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QXmlInputSource>
#include "_cgo_export.h"

class MyQXmlInputSource: public QXmlInputSource {
public:
};

void* QXmlInputSource_NewQXmlInputSource(){
	return new QXmlInputSource();
}

void* QXmlInputSource_NewQXmlInputSource2(void* dev){
	return new QXmlInputSource(static_cast<QIODevice*>(dev));
}

char* QXmlInputSource_Data(void* ptr){
	return static_cast<QXmlInputSource*>(ptr)->data().toUtf8().data();
}

void QXmlInputSource_FetchData(void* ptr){
	static_cast<QXmlInputSource*>(ptr)->fetchData();
}

void QXmlInputSource_Reset(void* ptr){
	static_cast<QXmlInputSource*>(ptr)->reset();
}

void QXmlInputSource_SetData2(void* ptr, void* dat){
	static_cast<QXmlInputSource*>(ptr)->setData(*static_cast<QByteArray*>(dat));
}

void QXmlInputSource_SetData(void* ptr, char* dat){
	static_cast<QXmlInputSource*>(ptr)->setData(QString(dat));
}

void QXmlInputSource_DestroyQXmlInputSource(void* ptr){
	static_cast<QXmlInputSource*>(ptr)->~QXmlInputSource();
}

