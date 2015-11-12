#include "qcollator.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStringRef>
#include <QLocale>
#include <QChar>
#include <QString>
#include <QCollator>
#include "_cgo_export.h"

class MyQCollator: public QCollator {
public:
};

int QCollator_CaseSensitivity(void* ptr){
	return static_cast<QCollator*>(ptr)->caseSensitivity();
}

int QCollator_IgnorePunctuation(void* ptr){
	return static_cast<QCollator*>(ptr)->ignorePunctuation();
}

int QCollator_NumericMode(void* ptr){
	return static_cast<QCollator*>(ptr)->numericMode();
}

void QCollator_SetCaseSensitivity(void* ptr, int sensitivity){
	static_cast<QCollator*>(ptr)->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(sensitivity));
}

void QCollator_SetIgnorePunctuation(void* ptr, int on){
	static_cast<QCollator*>(ptr)->setIgnorePunctuation(on != 0);
}

void QCollator_SetNumericMode(void* ptr, int on){
	static_cast<QCollator*>(ptr)->setNumericMode(on != 0);
}

void* QCollator_NewQCollator3(void* other){
	return new QCollator(*static_cast<QCollator*>(other));
}

void* QCollator_NewQCollator2(void* other){
	return new QCollator(*static_cast<QCollator*>(other));
}

void* QCollator_NewQCollator(void* locale){
	return new QCollator(*static_cast<QLocale*>(locale));
}

void QCollator_SetLocale(void* ptr, void* locale){
	static_cast<QCollator*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QCollator_Swap(void* ptr, void* other){
	static_cast<QCollator*>(ptr)->swap(*static_cast<QCollator*>(other));
}

void QCollator_DestroyQCollator(void* ptr){
	static_cast<QCollator*>(ptr)->~QCollator();
}

int QCollator_Compare3(void* ptr, void* s1, int len1, void* s2, int len2){
	return static_cast<QCollator*>(ptr)->compare(static_cast<QChar*>(s1), len1, static_cast<QChar*>(s2), len2);
}

int QCollator_Compare(void* ptr, char* s1, char* s2){
	return static_cast<QCollator*>(ptr)->compare(QString(s1), QString(s2));
}

int QCollator_Compare2(void* ptr, void* s1, void* s2){
	return static_cast<QCollator*>(ptr)->compare(*static_cast<QStringRef*>(s1), *static_cast<QStringRef*>(s2));
}

