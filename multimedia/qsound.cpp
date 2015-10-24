#include "qsound.h"
#include <QObject>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSound>
#include "_cgo_export.h"

class MyQSound: public QSound {
public:
};

void QSound_SetLoops(QtObjectPtr ptr, int number){
	static_cast<QSound*>(ptr)->setLoops(number);
}

QtObjectPtr QSound_NewQSound(char* filename, QtObjectPtr parent){
	return new QSound(QString(filename), static_cast<QObject*>(parent));
}

char* QSound_FileName(QtObjectPtr ptr){
	return static_cast<QSound*>(ptr)->fileName().toUtf8().data();
}

int QSound_IsFinished(QtObjectPtr ptr){
	return static_cast<QSound*>(ptr)->isFinished();
}

int QSound_Loops(QtObjectPtr ptr){
	return static_cast<QSound*>(ptr)->loops();
}

int QSound_LoopsRemaining(QtObjectPtr ptr){
	return static_cast<QSound*>(ptr)->loopsRemaining();
}

void QSound_Play2(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSound*>(ptr), "play");
}

void QSound_QSound_Play(char* filename){
	QSound::play(QString(filename));
}

void QSound_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSound*>(ptr), "stop");
}

void QSound_DestroyQSound(QtObjectPtr ptr){
	static_cast<QSound*>(ptr)->~QSound();
}

