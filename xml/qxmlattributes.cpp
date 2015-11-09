#include "qxmlattributes.h"
#include <QUrl>
#include <QModelIndex>
#include <QLatin1String>
#include <QString>
#include <QVariant>
#include <QXmlAttributes>
#include "_cgo_export.h"

class MyQXmlAttributes: public QXmlAttributes {
public:
};

void* QXmlAttributes_NewQXmlAttributes(){
	return new QXmlAttributes();
}

void QXmlAttributes_DestroyQXmlAttributes(void* ptr){
	static_cast<QXmlAttributes*>(ptr)->~QXmlAttributes();
}

void QXmlAttributes_Append(void* ptr, char* qName, char* uri, char* localPart, char* value){
	static_cast<QXmlAttributes*>(ptr)->append(QString(qName), QString(uri), QString(localPart), QString(value));
}

void QXmlAttributes_Clear(void* ptr){
	static_cast<QXmlAttributes*>(ptr)->clear();
}

int QXmlAttributes_Count(void* ptr){
	return static_cast<QXmlAttributes*>(ptr)->count();
}

int QXmlAttributes_Index2(void* ptr, void* qName){
	return static_cast<QXmlAttributes*>(ptr)->index(*static_cast<QLatin1String*>(qName));
}

int QXmlAttributes_Index(void* ptr, char* qName){
	return static_cast<QXmlAttributes*>(ptr)->index(QString(qName));
}

int QXmlAttributes_Index3(void* ptr, char* uri, char* localPart){
	return static_cast<QXmlAttributes*>(ptr)->index(QString(uri), QString(localPart));
}

int QXmlAttributes_Length(void* ptr){
	return static_cast<QXmlAttributes*>(ptr)->length();
}

char* QXmlAttributes_LocalName(void* ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->localName(index).toUtf8().data();
}

char* QXmlAttributes_QName(void* ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->qName(index).toUtf8().data();
}

char* QXmlAttributes_Type2(void* ptr, char* qName){
	return static_cast<QXmlAttributes*>(ptr)->type(QString(qName)).toUtf8().data();
}

char* QXmlAttributes_Type3(void* ptr, char* uri, char* localName){
	return static_cast<QXmlAttributes*>(ptr)->type(QString(uri), QString(localName)).toUtf8().data();
}

char* QXmlAttributes_Type(void* ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->type(index).toUtf8().data();
}

char* QXmlAttributes_Uri(void* ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->uri(index).toUtf8().data();
}

char* QXmlAttributes_Value3(void* ptr, void* qName){
	return static_cast<QXmlAttributes*>(ptr)->value(*static_cast<QLatin1String*>(qName)).toUtf8().data();
}

char* QXmlAttributes_Value2(void* ptr, char* qName){
	return static_cast<QXmlAttributes*>(ptr)->value(QString(qName)).toUtf8().data();
}

char* QXmlAttributes_Value4(void* ptr, char* uri, char* localName){
	return static_cast<QXmlAttributes*>(ptr)->value(QString(uri), QString(localName)).toUtf8().data();
}

char* QXmlAttributes_Value(void* ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->value(index).toUtf8().data();
}

