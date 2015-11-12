#include "qsound.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QSound>
#include "_cgo_export.h"

class MyQSound: public QSound {
public:
};

void QSound_SetLoops(void* ptr, int number){
	static_cast<QSound*>(ptr)->setLoops(number);
}

void* QSound_NewQSound(char* filename, void* parent){
	return new QSound(QString(filename), static_cast<QObject*>(parent));
}

char* QSound_FileName(void* ptr){
	return static_cast<QSound*>(ptr)->fileName().toUtf8().data();
}

int QSound_IsFinished(void* ptr){
	return static_cast<QSound*>(ptr)->isFinished();
}

int QSound_Loops(void* ptr){
	return static_cast<QSound*>(ptr)->loops();
}

int QSound_LoopsRemaining(void* ptr){
	return static_cast<QSound*>(ptr)->loopsRemaining();
}

void QSound_Play2(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSound*>(ptr), "play");
}

void QSound_QSound_Play(char* filename){
	QSound::play(QString(filename));
}

void QSound_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSound*>(ptr), "stop");
}

void QSound_DestroyQSound(void* ptr){
	static_cast<QSound*>(ptr)->~QSound();
}

