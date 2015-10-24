#include "qregularexpressionmatch.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRegularExpression>
#include <QRegularExpressionMatch>
#include "_cgo_export.h"

class MyQRegularExpressionMatch: public QRegularExpressionMatch {
public:
};

QtObjectPtr QRegularExpressionMatch_NewQRegularExpressionMatch(){
	return new QRegularExpressionMatch();
}

QtObjectPtr QRegularExpressionMatch_NewQRegularExpressionMatch2(QtObjectPtr match){
	return new QRegularExpressionMatch(*static_cast<QRegularExpressionMatch*>(match));
}

char* QRegularExpressionMatch_Captured2(QtObjectPtr ptr, char* name){
	return static_cast<QRegularExpressionMatch*>(ptr)->captured(QString(name)).toUtf8().data();
}

char* QRegularExpressionMatch_Captured(QtObjectPtr ptr, int nth){
	return static_cast<QRegularExpressionMatch*>(ptr)->captured(nth).toUtf8().data();
}

int QRegularExpressionMatch_CapturedEnd2(QtObjectPtr ptr, char* name){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedEnd(QString(name));
}

int QRegularExpressionMatch_CapturedEnd(QtObjectPtr ptr, int nth){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedEnd(nth);
}

int QRegularExpressionMatch_CapturedLength2(QtObjectPtr ptr, char* name){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedLength(QString(name));
}

int QRegularExpressionMatch_CapturedLength(QtObjectPtr ptr, int nth){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedLength(nth);
}

int QRegularExpressionMatch_CapturedStart2(QtObjectPtr ptr, char* name){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedStart(QString(name));
}

int QRegularExpressionMatch_CapturedStart(QtObjectPtr ptr, int nth){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedStart(nth);
}

char* QRegularExpressionMatch_CapturedTexts(QtObjectPtr ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->capturedTexts().join("|").toUtf8().data();
}

int QRegularExpressionMatch_HasMatch(QtObjectPtr ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->hasMatch();
}

int QRegularExpressionMatch_HasPartialMatch(QtObjectPtr ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->hasPartialMatch();
}

int QRegularExpressionMatch_IsValid(QtObjectPtr ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->isValid();
}

int QRegularExpressionMatch_LastCapturedIndex(QtObjectPtr ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->lastCapturedIndex();
}

int QRegularExpressionMatch_MatchOptions(QtObjectPtr ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->matchOptions();
}

int QRegularExpressionMatch_MatchType(QtObjectPtr ptr){
	return static_cast<QRegularExpressionMatch*>(ptr)->matchType();
}

void QRegularExpressionMatch_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QRegularExpressionMatch*>(ptr)->swap(*static_cast<QRegularExpressionMatch*>(other));
}

void QRegularExpressionMatch_DestroyQRegularExpressionMatch(QtObjectPtr ptr){
	static_cast<QRegularExpressionMatch*>(ptr)->~QRegularExpressionMatch();
}

