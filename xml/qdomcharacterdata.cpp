#include "qdomcharacterdata.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomCharacterData>
#include "_cgo_export.h"

class MyQDomCharacterData: public QDomCharacterData {
public:
};

void* QDomCharacterData_NewQDomCharacterData(){
	return new QDomCharacterData();
}

void* QDomCharacterData_NewQDomCharacterData2(void* x){
	return new QDomCharacterData(*static_cast<QDomCharacterData*>(x));
}

void QDomCharacterData_AppendData(void* ptr, char* arg){
	static_cast<QDomCharacterData*>(ptr)->appendData(QString(arg));
}

char* QDomCharacterData_Data(void* ptr){
	return static_cast<QDomCharacterData*>(ptr)->data().toUtf8().data();
}

int QDomCharacterData_Length(void* ptr){
	return static_cast<QDomCharacterData*>(ptr)->length();
}

int QDomCharacterData_NodeType(void* ptr){
	return static_cast<QDomCharacterData*>(ptr)->nodeType();
}

void QDomCharacterData_SetData(void* ptr, char* v){
	static_cast<QDomCharacterData*>(ptr)->setData(QString(v));
}

