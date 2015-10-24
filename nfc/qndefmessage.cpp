#include "qndefmessage.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNdefRecord>
#include <QNdefMessage>
#include "_cgo_export.h"

class MyQNdefMessage: public QNdefMessage {
public:
};

QtObjectPtr QNdefMessage_NewQNdefMessage(){
	return new QNdefMessage();
}

QtObjectPtr QNdefMessage_NewQNdefMessage3(QtObjectPtr message){
	return new QNdefMessage(*static_cast<QNdefMessage*>(message));
}

QtObjectPtr QNdefMessage_NewQNdefMessage2(QtObjectPtr record){
	return new QNdefMessage(*static_cast<QNdefRecord*>(record));
}

