#include "qcollator.h"
#include <QModelIndex>
#include <QStringRef>
#include <QChar>
#include <QLocale>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QCollator>
#include "_cgo_export.h"

class MyQCollator: public QCollator {
public:
};

int QCollator_CaseSensitivity(QtObjectPtr ptr){
	return static_cast<QCollator*>(ptr)->caseSensitivity();
}

int QCollator_IgnorePunctuation(QtObjectPtr ptr){
	return static_cast<QCollator*>(ptr)->ignorePunctuation();
}

int QCollator_NumericMode(QtObjectPtr ptr){
	return static_cast<QCollator*>(ptr)->numericMode();
}

void QCollator_SetCaseSensitivity(QtObjectPtr ptr, int sensitivity){
	static_cast<QCollator*>(ptr)->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(sensitivity));
}

void QCollator_SetIgnorePunctuation(QtObjectPtr ptr, int on){
	static_cast<QCollator*>(ptr)->setIgnorePunctuation(on != 0);
}

void QCollator_SetNumericMode(QtObjectPtr ptr, int on){
	static_cast<QCollator*>(ptr)->setNumericMode(on != 0);
}

QtObjectPtr QCollator_NewQCollator3(QtObjectPtr other){
	return new QCollator(*static_cast<QCollator*>(other));
}

QtObjectPtr QCollator_NewQCollator2(QtObjectPtr other){
	return new QCollator(*static_cast<QCollator*>(other));
}

QtObjectPtr QCollator_NewQCollator(QtObjectPtr locale){
	return new QCollator(*static_cast<QLocale*>(locale));
}

void QCollator_SetLocale(QtObjectPtr ptr, QtObjectPtr locale){
	static_cast<QCollator*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QCollator_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QCollator*>(ptr)->swap(*static_cast<QCollator*>(other));
}

void QCollator_DestroyQCollator(QtObjectPtr ptr){
	static_cast<QCollator*>(ptr)->~QCollator();
}

int QCollator_Compare3(QtObjectPtr ptr, QtObjectPtr s1, int len1, QtObjectPtr s2, int len2){
	return static_cast<QCollator*>(ptr)->compare(static_cast<QChar*>(s1), len1, static_cast<QChar*>(s2), len2);
}

int QCollator_Compare(QtObjectPtr ptr, char* s1, char* s2){
	return static_cast<QCollator*>(ptr)->compare(QString(s1), QString(s2));
}

int QCollator_Compare2(QtObjectPtr ptr, QtObjectPtr s1, QtObjectPtr s2){
	return static_cast<QCollator*>(ptr)->compare(*static_cast<QStringRef*>(s1), *static_cast<QStringRef*>(s2));
}

