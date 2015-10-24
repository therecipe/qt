#include "qregexp.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QRegExp>
#include "_cgo_export.h"

class MyQRegExp: public QRegExp {
public:
};

QtObjectPtr QRegExp_NewQRegExp(){
	return new QRegExp();
}

QtObjectPtr QRegExp_NewQRegExp3(QtObjectPtr rx){
	return new QRegExp(*static_cast<QRegExp*>(rx));
}

QtObjectPtr QRegExp_NewQRegExp2(char* pattern, int cs, int syntax){
	return new QRegExp(QString(pattern), static_cast<Qt::CaseSensitivity>(cs), static_cast<QRegExp::PatternSyntax>(syntax));
}

char* QRegExp_Cap(QtObjectPtr ptr, int nth){
	return static_cast<QRegExp*>(ptr)->cap(nth).toUtf8().data();
}

char* QRegExp_ErrorString(QtObjectPtr ptr){
	return static_cast<QRegExp*>(ptr)->errorString().toUtf8().data();
}

int QRegExp_CaptureCount(QtObjectPtr ptr){
	return static_cast<QRegExp*>(ptr)->captureCount();
}

char* QRegExp_CapturedTexts(QtObjectPtr ptr){
	return static_cast<QRegExp*>(ptr)->capturedTexts().join("|").toUtf8().data();
}

int QRegExp_CaseSensitivity(QtObjectPtr ptr){
	return static_cast<QRegExp*>(ptr)->caseSensitivity();
}

char* QRegExp_QRegExp_Escape(char* str){
	return QRegExp::escape(QString(str)).toUtf8().data();
}

int QRegExp_ExactMatch(QtObjectPtr ptr, char* str){
	return static_cast<QRegExp*>(ptr)->exactMatch(QString(str));
}

int QRegExp_IndexIn(QtObjectPtr ptr, char* str, int offset, int caretMode){
	return static_cast<QRegExp*>(ptr)->indexIn(QString(str), offset, static_cast<QRegExp::CaretMode>(caretMode));
}

int QRegExp_IsEmpty(QtObjectPtr ptr){
	return static_cast<QRegExp*>(ptr)->isEmpty();
}

int QRegExp_IsMinimal(QtObjectPtr ptr){
	return static_cast<QRegExp*>(ptr)->isMinimal();
}

int QRegExp_IsValid(QtObjectPtr ptr){
	return static_cast<QRegExp*>(ptr)->isValid();
}

int QRegExp_LastIndexIn(QtObjectPtr ptr, char* str, int offset, int caretMode){
	return static_cast<QRegExp*>(ptr)->lastIndexIn(QString(str), offset, static_cast<QRegExp::CaretMode>(caretMode));
}

int QRegExp_MatchedLength(QtObjectPtr ptr){
	return static_cast<QRegExp*>(ptr)->matchedLength();
}

char* QRegExp_Pattern(QtObjectPtr ptr){
	return static_cast<QRegExp*>(ptr)->pattern().toUtf8().data();
}

int QRegExp_PatternSyntax(QtObjectPtr ptr){
	return static_cast<QRegExp*>(ptr)->patternSyntax();
}

int QRegExp_Pos(QtObjectPtr ptr, int nth){
	return static_cast<QRegExp*>(ptr)->pos(nth);
}

void QRegExp_SetCaseSensitivity(QtObjectPtr ptr, int cs){
	static_cast<QRegExp*>(ptr)->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QRegExp_SetMinimal(QtObjectPtr ptr, int minimal){
	static_cast<QRegExp*>(ptr)->setMinimal(minimal != 0);
}

void QRegExp_SetPattern(QtObjectPtr ptr, char* pattern){
	static_cast<QRegExp*>(ptr)->setPattern(QString(pattern));
}

void QRegExp_SetPatternSyntax(QtObjectPtr ptr, int syntax){
	static_cast<QRegExp*>(ptr)->setPatternSyntax(static_cast<QRegExp::PatternSyntax>(syntax));
}

void QRegExp_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QRegExp*>(ptr)->swap(*static_cast<QRegExp*>(other));
}

void QRegExp_DestroyQRegExp(QtObjectPtr ptr){
	static_cast<QRegExp*>(ptr)->~QRegExp();
}

