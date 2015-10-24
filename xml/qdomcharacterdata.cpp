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

QtObjectPtr QDomCharacterData_NewQDomCharacterData(){
	return new QDomCharacterData();
}

QtObjectPtr QDomCharacterData_NewQDomCharacterData2(QtObjectPtr x){
	return new QDomCharacterData(*static_cast<QDomCharacterData*>(x));
}

void QDomCharacterData_AppendData(QtObjectPtr ptr, char* arg){
	static_cast<QDomCharacterData*>(ptr)->appendData(QString(arg));
}

char* QDomCharacterData_Data(QtObjectPtr ptr){
	return static_cast<QDomCharacterData*>(ptr)->data().toUtf8().data();
}

int QDomCharacterData_Length(QtObjectPtr ptr){
	return static_cast<QDomCharacterData*>(ptr)->length();
}

int QDomCharacterData_NodeType(QtObjectPtr ptr){
	return static_cast<QDomCharacterData*>(ptr)->nodeType();
}

void QDomCharacterData_SetData(QtObjectPtr ptr, char* v){
	static_cast<QDomCharacterData*>(ptr)->setData(QString(v));
}

