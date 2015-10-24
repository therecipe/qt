#include "qregularexpression.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRegularExpression>
#include "_cgo_export.h"

class MyQRegularExpression: public QRegularExpression {
public:
};

QtObjectPtr QRegularExpression_NewQRegularExpression(){
	return new QRegularExpression();
}

QtObjectPtr QRegularExpression_NewQRegularExpression3(QtObjectPtr re){
	return new QRegularExpression(*static_cast<QRegularExpression*>(re));
}

QtObjectPtr QRegularExpression_NewQRegularExpression2(char* pattern, int options){
	return new QRegularExpression(QString(pattern), static_cast<QRegularExpression::PatternOption>(options));
}

int QRegularExpression_CaptureCount(QtObjectPtr ptr){
	return static_cast<QRegularExpression*>(ptr)->captureCount();
}

char* QRegularExpression_ErrorString(QtObjectPtr ptr){
	return static_cast<QRegularExpression*>(ptr)->errorString().toUtf8().data();
}

char* QRegularExpression_QRegularExpression_Escape(char* str){
	return QRegularExpression::escape(QString(str)).toUtf8().data();
}

int QRegularExpression_IsValid(QtObjectPtr ptr){
	return static_cast<QRegularExpression*>(ptr)->isValid();
}

char* QRegularExpression_NamedCaptureGroups(QtObjectPtr ptr){
	return static_cast<QRegularExpression*>(ptr)->namedCaptureGroups().join("|").toUtf8().data();
}

void QRegularExpression_Optimize(QtObjectPtr ptr){
	static_cast<QRegularExpression*>(ptr)->optimize();
}

char* QRegularExpression_Pattern(QtObjectPtr ptr){
	return static_cast<QRegularExpression*>(ptr)->pattern().toUtf8().data();
}

int QRegularExpression_PatternErrorOffset(QtObjectPtr ptr){
	return static_cast<QRegularExpression*>(ptr)->patternErrorOffset();
}

int QRegularExpression_PatternOptions(QtObjectPtr ptr){
	return static_cast<QRegularExpression*>(ptr)->patternOptions();
}

void QRegularExpression_SetPattern(QtObjectPtr ptr, char* pattern){
	static_cast<QRegularExpression*>(ptr)->setPattern(QString(pattern));
}

void QRegularExpression_SetPatternOptions(QtObjectPtr ptr, int options){
	static_cast<QRegularExpression*>(ptr)->setPatternOptions(static_cast<QRegularExpression::PatternOption>(options));
}

void QRegularExpression_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QRegularExpression*>(ptr)->swap(*static_cast<QRegularExpression*>(other));
}

void QRegularExpression_DestroyQRegularExpression(QtObjectPtr ptr){
	static_cast<QRegularExpression*>(ptr)->~QRegularExpression();
}

