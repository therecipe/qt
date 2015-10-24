#include "qinputmethodevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QInputMethod>
#include <QInputMethodEvent>
#include "_cgo_export.h"

class MyQInputMethodEvent: public QInputMethodEvent {
public:
};

QtObjectPtr QInputMethodEvent_NewQInputMethodEvent(){
	return new QInputMethodEvent();
}

QtObjectPtr QInputMethodEvent_NewQInputMethodEvent3(QtObjectPtr other){
	return new QInputMethodEvent(*static_cast<QInputMethodEvent*>(other));
}

char* QInputMethodEvent_CommitString(QtObjectPtr ptr){
	return static_cast<QInputMethodEvent*>(ptr)->commitString().toUtf8().data();
}

char* QInputMethodEvent_PreeditString(QtObjectPtr ptr){
	return static_cast<QInputMethodEvent*>(ptr)->preeditString().toUtf8().data();
}

int QInputMethodEvent_ReplacementLength(QtObjectPtr ptr){
	return static_cast<QInputMethodEvent*>(ptr)->replacementLength();
}

int QInputMethodEvent_ReplacementStart(QtObjectPtr ptr){
	return static_cast<QInputMethodEvent*>(ptr)->replacementStart();
}

void QInputMethodEvent_SetCommitString(QtObjectPtr ptr, char* commitString, int replaceFrom, int replaceLength){
	static_cast<QInputMethodEvent*>(ptr)->setCommitString(QString(commitString), replaceFrom, replaceLength);
}

