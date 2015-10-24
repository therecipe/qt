#include "qkeysequence.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QKeySequence>
#include "_cgo_export.h"

class MyQKeySequence: public QKeySequence {
public:
};

QtObjectPtr QKeySequence_NewQKeySequence(){
	return new QKeySequence();
}

QtObjectPtr QKeySequence_NewQKeySequence5(int key){
	return new QKeySequence(static_cast<QKeySequence::StandardKey>(key));
}

QtObjectPtr QKeySequence_NewQKeySequence4(QtObjectPtr keysequence){
	return new QKeySequence(*static_cast<QKeySequence*>(keysequence));
}

QtObjectPtr QKeySequence_NewQKeySequence2(char* key, int format){
	return new QKeySequence(QString(key), static_cast<QKeySequence::SequenceFormat>(format));
}

QtObjectPtr QKeySequence_NewQKeySequence3(int k1, int k2, int k3, int k4){
	return new QKeySequence(k1, k2, k3, k4);
}

int QKeySequence_Count(QtObjectPtr ptr){
	return static_cast<QKeySequence*>(ptr)->count();
}

int QKeySequence_IsEmpty(QtObjectPtr ptr){
	return static_cast<QKeySequence*>(ptr)->isEmpty();
}

int QKeySequence_Matches(QtObjectPtr ptr, QtObjectPtr seq){
	return static_cast<QKeySequence*>(ptr)->matches(*static_cast<QKeySequence*>(seq));
}

void QKeySequence_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QKeySequence*>(ptr)->swap(*static_cast<QKeySequence*>(other));
}

char* QKeySequence_ToString(QtObjectPtr ptr, int format){
	return static_cast<QKeySequence*>(ptr)->toString(static_cast<QKeySequence::SequenceFormat>(format)).toUtf8().data();
}

void QKeySequence_DestroyQKeySequence(QtObjectPtr ptr){
	static_cast<QKeySequence*>(ptr)->~QKeySequence();
}

