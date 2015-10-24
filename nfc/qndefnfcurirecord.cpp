#include "qndefnfcurirecord.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNdefRecord>
#include <QNdefNfcUriRecord>
#include "_cgo_export.h"

class MyQNdefNfcUriRecord: public QNdefNfcUriRecord {
public:
};

QtObjectPtr QNdefNfcUriRecord_NewQNdefNfcUriRecord(){
	return new QNdefNfcUriRecord();
}

QtObjectPtr QNdefNfcUriRecord_NewQNdefNfcUriRecord2(QtObjectPtr other){
	return new QNdefNfcUriRecord(*static_cast<QNdefRecord*>(other));
}

void QNdefNfcUriRecord_SetUri(QtObjectPtr ptr, char* uri){
	static_cast<QNdefNfcUriRecord*>(ptr)->setUri(QUrl(QString(uri)));
}

char* QNdefNfcUriRecord_Uri(QtObjectPtr ptr){
	return static_cast<QNdefNfcUriRecord*>(ptr)->uri().toString().toUtf8().data();
}

