#include "qxmlattributes.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLatin1String>
#include <QString>
#include <QXmlAttributes>
#include "_cgo_export.h"

class MyQXmlAttributes: public QXmlAttributes {
public:
};

QtObjectPtr QXmlAttributes_NewQXmlAttributes(){
	return new QXmlAttributes();
}

void QXmlAttributes_DestroyQXmlAttributes(QtObjectPtr ptr){
	static_cast<QXmlAttributes*>(ptr)->~QXmlAttributes();
}

void QXmlAttributes_Append(QtObjectPtr ptr, char* qName, char* uri, char* localPart, char* value){
	static_cast<QXmlAttributes*>(ptr)->append(QString(qName), QString(uri), QString(localPart), QString(value));
}

void QXmlAttributes_Clear(QtObjectPtr ptr){
	static_cast<QXmlAttributes*>(ptr)->clear();
}

int QXmlAttributes_Count(QtObjectPtr ptr){
	return static_cast<QXmlAttributes*>(ptr)->count();
}

int QXmlAttributes_Index2(QtObjectPtr ptr, QtObjectPtr qName){
	return static_cast<QXmlAttributes*>(ptr)->index(*static_cast<QLatin1String*>(qName));
}

int QXmlAttributes_Index(QtObjectPtr ptr, char* qName){
	return static_cast<QXmlAttributes*>(ptr)->index(QString(qName));
}

int QXmlAttributes_Index3(QtObjectPtr ptr, char* uri, char* localPart){
	return static_cast<QXmlAttributes*>(ptr)->index(QString(uri), QString(localPart));
}

int QXmlAttributes_Length(QtObjectPtr ptr){
	return static_cast<QXmlAttributes*>(ptr)->length();
}

char* QXmlAttributes_LocalName(QtObjectPtr ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->localName(index).toUtf8().data();
}

char* QXmlAttributes_QName(QtObjectPtr ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->qName(index).toUtf8().data();
}

char* QXmlAttributes_Type2(QtObjectPtr ptr, char* qName){
	return static_cast<QXmlAttributes*>(ptr)->type(QString(qName)).toUtf8().data();
}

char* QXmlAttributes_Type3(QtObjectPtr ptr, char* uri, char* localName){
	return static_cast<QXmlAttributes*>(ptr)->type(QString(uri), QString(localName)).toUtf8().data();
}

char* QXmlAttributes_Type(QtObjectPtr ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->type(index).toUtf8().data();
}

char* QXmlAttributes_Uri(QtObjectPtr ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->uri(index).toUtf8().data();
}

char* QXmlAttributes_Value3(QtObjectPtr ptr, QtObjectPtr qName){
	return static_cast<QXmlAttributes*>(ptr)->value(*static_cast<QLatin1String*>(qName)).toUtf8().data();
}

char* QXmlAttributes_Value2(QtObjectPtr ptr, char* qName){
	return static_cast<QXmlAttributes*>(ptr)->value(QString(qName)).toUtf8().data();
}

char* QXmlAttributes_Value4(QtObjectPtr ptr, char* uri, char* localName){
	return static_cast<QXmlAttributes*>(ptr)->value(QString(uri), QString(localName)).toUtf8().data();
}

char* QXmlAttributes_Value(QtObjectPtr ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->value(index).toUtf8().data();
}

