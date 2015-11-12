#include "qinputmethodevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QInputMethod>
#include <QString>
#include <QVariant>
#include <QInputMethodEvent>
#include "_cgo_export.h"

class MyQInputMethodEvent: public QInputMethodEvent {
public:
};

void* QInputMethodEvent_NewQInputMethodEvent(){
	return new QInputMethodEvent();
}

void* QInputMethodEvent_NewQInputMethodEvent3(void* other){
	return new QInputMethodEvent(*static_cast<QInputMethodEvent*>(other));
}

char* QInputMethodEvent_CommitString(void* ptr){
	return static_cast<QInputMethodEvent*>(ptr)->commitString().toUtf8().data();
}

char* QInputMethodEvent_PreeditString(void* ptr){
	return static_cast<QInputMethodEvent*>(ptr)->preeditString().toUtf8().data();
}

int QInputMethodEvent_ReplacementLength(void* ptr){
	return static_cast<QInputMethodEvent*>(ptr)->replacementLength();
}

int QInputMethodEvent_ReplacementStart(void* ptr){
	return static_cast<QInputMethodEvent*>(ptr)->replacementStart();
}

void QInputMethodEvent_SetCommitString(void* ptr, char* commitString, int replaceFrom, int replaceLength){
	static_cast<QInputMethodEvent*>(ptr)->setCommitString(QString(commitString), replaceFrom, replaceLength);
}

