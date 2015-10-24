#include "qxmlinputsource.h"
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QXmlInputSource>
#include "_cgo_export.h"

class MyQXmlInputSource: public QXmlInputSource {
public:
};

QtObjectPtr QXmlInputSource_NewQXmlInputSource(){
	return new QXmlInputSource();
}

QtObjectPtr QXmlInputSource_NewQXmlInputSource2(QtObjectPtr dev){
	return new QXmlInputSource(static_cast<QIODevice*>(dev));
}

char* QXmlInputSource_Data(QtObjectPtr ptr){
	return static_cast<QXmlInputSource*>(ptr)->data().toUtf8().data();
}

void QXmlInputSource_FetchData(QtObjectPtr ptr){
	static_cast<QXmlInputSource*>(ptr)->fetchData();
}

void QXmlInputSource_Reset(QtObjectPtr ptr){
	static_cast<QXmlInputSource*>(ptr)->reset();
}

void QXmlInputSource_SetData2(QtObjectPtr ptr, QtObjectPtr dat){
	static_cast<QXmlInputSource*>(ptr)->setData(*static_cast<QByteArray*>(dat));
}

void QXmlInputSource_SetData(QtObjectPtr ptr, char* dat){
	static_cast<QXmlInputSource*>(ptr)->setData(QString(dat));
}

void QXmlInputSource_DestroyQXmlInputSource(QtObjectPtr ptr){
	static_cast<QXmlInputSource*>(ptr)->~QXmlInputSource();
}

