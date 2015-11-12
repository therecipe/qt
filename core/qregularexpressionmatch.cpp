#include "qregularexpressionmatch.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRegularExpression>
#include <QStringRef>
#include <QRegularExpressionMatch>
#include "_cgo_export.h"

class MyQRegularExpressionMatch: public QRegularExpressionMatch {
public:
};

void* QRegularExpressionMatch_NewQRegularExpressionMatch(){
	return new QRegularExpressionMatch();
}

void* QRegularExpressionMatch_NewQRegularExpressionMatch2(void* match){
	return new QRegularExpressionMatch(*static_cast<QRegularExpressionMatch*>(match));
}

char* QRegularExpressionMatch_Captured2(void* ptr, char* name){
	return static_cast<QRegularExpressionMatch*>(ptr)->captured(QString(name)).toUtf8().data();
}

char* QRegularExpressionMatch_Captured(void* ptr, int nth){
	return static_cast<QRegularExpressionMatch*>(ptr)->captured(nth).toUtf8().data();
}

int QRegularExpressionMatch_CapturedEnd2(void* ptr, char* name){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedEnd(QString(name));
}

int QRegularExpressionMatch_CapturedEnd(void* ptr, int nth){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedEnd(nth);
}

int QRegularExpressionMatch_CapturedLength2(void* ptr, char* name){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedLength(QString(name));
}

int QRegularExpressionMatch_CapturedLength(void* ptr, int nth){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedLength(nth);
}

void* QRegularExpressionMatch_CapturedRef2(void* ptr, char* name){
	return new QStringRef(static_cast<QRegularExpressionMatch*>(ptr)->capturedRef(QString(name)));
}

void* QRegularExpressionMatch_CapturedRef(void* ptr, int nth){
	return new QStringRef(static_cast<QRegularExpressionMatch*>(ptr)->capturedRef(nth));
}

int QRegularExpressionMatch_CapturedStart2(void* ptr, char* name){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedStart(QString(name));
}

int QRegularExpressionMatch_CapturedStart(void* ptr, int nth){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedStart(nth);
}

char* QRegularExpressionMatch_CapturedTexts(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedTexts().join("|").toUtf8().data();
}

int QRegularExpressionMatch_HasMatch(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->hasMatch();
}

int QRegularExpressionMatch_HasPartialMatch(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->hasPartialMatch();
}

int QRegularExpressionMatch_IsValid(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->isValid();
}

int QRegularExpressionMatch_LastCapturedIndex(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->lastCapturedIndex();
}

int QRegularExpressionMatch_MatchOptions(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->matchOptions();
}

int QRegularExpressionMatch_MatchType(void* ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->matchType();
}

void* QRegularExpressionMatch_RegularExpression(void* ptr){
	return new QRegularExpression(static_cast<QRegularExpressionMatch*>(ptr)->regularExpression());
}

void QRegularExpressionMatch_Swap(void* ptr, void* other){
	static_cast<QRegularExpressionMatch*>(ptr)->swap(*static_cast<QRegularExpressionMatch*>(other));
}

void QRegularExpressionMatch_DestroyQRegularExpressionMatch(void* ptr){
	static_cast<QRegularExpressionMatch*>(ptr)->~QRegularExpressionMatch();
}

