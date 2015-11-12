#include "qregularexpression.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStringRef>
#include <QRegularExpressionMatchIterator>
#include <QRegularExpressionMatch>
#include <QString>
#include <QRegularExpression>
#include "_cgo_export.h"

class MyQRegularExpression: public QRegularExpression {
public:
};

void* QRegularExpression_NewQRegularExpression(){
	return new QRegularExpression();
}

void* QRegularExpression_NewQRegularExpression3(void* re){
	return new QRegularExpression(*static_cast<QRegularExpression*>(re));
}

void* QRegularExpression_NewQRegularExpression2(char* pattern, int options){
	return new QRegularExpression(QString(pattern), static_cast<QRegularExpression::PatternOption>(options));
}

int QRegularExpression_CaptureCount(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->captureCount();
}

char* QRegularExpression_ErrorString(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->errorString().toUtf8().data();
}

char* QRegularExpression_QRegularExpression_Escape(char* str){
	return QRegularExpression::escape(QString(str)).toUtf8().data();
}

void* QRegularExpression_GlobalMatch(void* ptr, char* subject, int offset, int matchType, int matchOptions){
	return new QRegularExpressionMatchIterator(static_cast<QRegularExpression*>(ptr)->globalMatch(QString(subject), offset, static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOption>(matchOptions)));
}

void* QRegularExpression_GlobalMatch2(void* ptr, void* subjectRef, int offset, int matchType, int matchOptions){
	return new QRegularExpressionMatchIterator(static_cast<QRegularExpression*>(ptr)->globalMatch(*static_cast<QStringRef*>(subjectRef), offset, static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOption>(matchOptions)));
}

int QRegularExpression_IsValid(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->isValid();
}

void* QRegularExpression_Match(void* ptr, char* subject, int offset, int matchType, int matchOptions){
	return new QRegularExpressionMatch(static_cast<QRegularExpression*>(ptr)->match(QString(subject), offset, static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOption>(matchOptions)));
}

void* QRegularExpression_Match2(void* ptr, void* subjectRef, int offset, int matchType, int matchOptions){
	return new QRegularExpressionMatch(static_cast<QRegularExpression*>(ptr)->match(*static_cast<QStringRef*>(subjectRef), offset, static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOption>(matchOptions)));
}

char* QRegularExpression_NamedCaptureGroups(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->namedCaptureGroups().join("|").toUtf8().data();
}

void QRegularExpression_Optimize(void* ptr){
	static_cast<QRegularExpression*>(ptr)->optimize();
}

char* QRegularExpression_Pattern(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->pattern().toUtf8().data();
}

int QRegularExpression_PatternErrorOffset(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->patternErrorOffset();
}

int QRegularExpression_PatternOptions(void* ptr){
	return static_cast<QRegularExpression*>(ptr)->patternOptions();
}

void QRegularExpression_SetPattern(void* ptr, char* pattern){
	static_cast<QRegularExpression*>(ptr)->setPattern(QString(pattern));
}

void QRegularExpression_SetPatternOptions(void* ptr, int options){
	static_cast<QRegularExpression*>(ptr)->setPatternOptions(static_cast<QRegularExpression::PatternOption>(options));
}

void QRegularExpression_Swap(void* ptr, void* other){
	static_cast<QRegularExpression*>(ptr)->swap(*static_cast<QRegularExpression*>(other));
}

void QRegularExpression_DestroyQRegularExpression(void* ptr){
	static_cast<QRegularExpression*>(ptr)->~QRegularExpression();
}

