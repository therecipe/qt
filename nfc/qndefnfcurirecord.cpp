#include "qndefnfcurirecord.h"
#include <QUrl>
#include <QModelIndex>
#include <QNdefRecord>
#include <QString>
#include <QVariant>
#include <QNdefNfcUriRecord>
#include "_cgo_export.h"

class MyQNdefNfcUriRecord: public QNdefNfcUriRecord {
public:
};

void* QNdefNfcUriRecord_NewQNdefNfcUriRecord(){
	return new QNdefNfcUriRecord();
}

void* QNdefNfcUriRecord_NewQNdefNfcUriRecord2(void* other){
	return new QNdefNfcUriRecord(*static_cast<QNdefRecord*>(other));
}

void QNdefNfcUriRecord_SetUri(void* ptr, void* uri){
	static_cast<QNdefNfcUriRecord*>(ptr)->setUri(*static_cast<QUrl*>(uri));
}

