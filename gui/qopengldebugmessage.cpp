#include "qopengldebugmessage.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLDebugMessage>
#include "_cgo_export.h"

class MyQOpenGLDebugMessage: public QOpenGLDebugMessage {
public:
};

QtObjectPtr QOpenGLDebugMessage_NewQOpenGLDebugMessage(){
	return new QOpenGLDebugMessage();
}

QtObjectPtr QOpenGLDebugMessage_NewQOpenGLDebugMessage2(QtObjectPtr debugMessage){
	return new QOpenGLDebugMessage(*static_cast<QOpenGLDebugMessage*>(debugMessage));
}

char* QOpenGLDebugMessage_Message(QtObjectPtr ptr){
	return static_cast<QOpenGLDebugMessage*>(ptr)->message().toUtf8().data();
}

int QOpenGLDebugMessage_Severity(QtObjectPtr ptr){
	return static_cast<QOpenGLDebugMessage*>(ptr)->severity();
}

int QOpenGLDebugMessage_Source(QtObjectPtr ptr){
	return static_cast<QOpenGLDebugMessage*>(ptr)->source();
}

void QOpenGLDebugMessage_Swap(QtObjectPtr ptr, QtObjectPtr debugMessage){
	static_cast<QOpenGLDebugMessage*>(ptr)->swap(*static_cast<QOpenGLDebugMessage*>(debugMessage));
}

int QOpenGLDebugMessage_Type(QtObjectPtr ptr){
	return static_cast<QOpenGLDebugMessage*>(ptr)->type();
}

void QOpenGLDebugMessage_DestroyQOpenGLDebugMessage(QtObjectPtr ptr){
	static_cast<QOpenGLDebugMessage*>(ptr)->~QOpenGLDebugMessage();
}

