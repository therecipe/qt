#include "qkeysequence.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QKeySequence>
#include "_cgo_export.h"

class MyQKeySequence: public QKeySequence {
public:
};

void* QKeySequence_NewQKeySequence(){
	return new QKeySequence();
}

void* QKeySequence_NewQKeySequence5(int key){
	return new QKeySequence(static_cast<QKeySequence::StandardKey>(key));
}

void* QKeySequence_NewQKeySequence4(void* keysequence){
	return new QKeySequence(*static_cast<QKeySequence*>(keysequence));
}

void* QKeySequence_NewQKeySequence2(char* key, int format){
	return new QKeySequence(QString(key), static_cast<QKeySequence::SequenceFormat>(format));
}

void* QKeySequence_NewQKeySequence3(int k1, int k2, int k3, int k4){
	return new QKeySequence(k1, k2, k3, k4);
}

int QKeySequence_Count(void* ptr){
	return static_cast<QKeySequence*>(ptr)->count();
}

int QKeySequence_IsEmpty(void* ptr){
	return static_cast<QKeySequence*>(ptr)->isEmpty();
}

int QKeySequence_Matches(void* ptr, void* seq){
	return static_cast<QKeySequence*>(ptr)->matches(*static_cast<QKeySequence*>(seq));
}

void QKeySequence_Swap(void* ptr, void* other){
	static_cast<QKeySequence*>(ptr)->swap(*static_cast<QKeySequence*>(other));
}

char* QKeySequence_ToString(void* ptr, int format){
	return static_cast<QKeySequence*>(ptr)->toString(static_cast<QKeySequence::SequenceFormat>(format)).toUtf8().data();
}

void QKeySequence_DestroyQKeySequence(void* ptr){
	static_cast<QKeySequence*>(ptr)->~QKeySequence();
}

