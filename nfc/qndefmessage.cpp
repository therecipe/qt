#include "qndefmessage.h"
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QNdefRecord>
#include <QString>
#include <QVariant>
#include <QNdefMessage>
#include "_cgo_export.h"

class MyQNdefMessage: public QNdefMessage {
public:
};

void* QNdefMessage_NewQNdefMessage(){
	return new QNdefMessage();
}

void* QNdefMessage_NewQNdefMessage3(void* message){
	return new QNdefMessage(*static_cast<QNdefMessage*>(message));
}

void* QNdefMessage_NewQNdefMessage2(void* record){
	return new QNdefMessage(*static_cast<QNdefRecord*>(record));
}

void* QNdefMessage_ToByteArray(void* ptr){
	return new QByteArray(static_cast<QNdefMessage*>(ptr)->toByteArray());
}

