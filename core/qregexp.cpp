#include "qregexp.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRegExp>
#include "_cgo_export.h"

class MyQRegExp: public QRegExp {
public:
};

void* QRegExp_NewQRegExp(){
	return new QRegExp();
}

void* QRegExp_NewQRegExp3(void* rx){
	return new QRegExp(*static_cast<QRegExp*>(rx));
}

void* QRegExp_NewQRegExp2(char* pattern, int cs, int syntax){
	return new QRegExp(QString(pattern), static_cast<Qt::CaseSensitivity>(cs), static_cast<QRegExp::PatternSyntax>(syntax));
}

char* QRegExp_Cap(void* ptr, int nth){
	return static_cast<QRegExp*>(ptr)->cap(nth).toUtf8().data();
}

char* QRegExp_ErrorString(void* ptr){
	return static_cast<QRegExp*>(ptr)->errorString().toUtf8().data();
}

int QRegExp_CaptureCount(void* ptr){
	return static_cast<QRegExp*>(ptr)->captureCount();
}

char* QRegExp_CapturedTexts(void* ptr){
	return static_cast<QRegExp*>(ptr)->capturedTexts().join("|").toUtf8().data();
}

int QRegExp_CaseSensitivity(void* ptr){
	return static_cast<QRegExp*>(ptr)->caseSensitivity();
}

char* QRegExp_QRegExp_Escape(char* str){
	return QRegExp::escape(QString(str)).toUtf8().data();
}

int QRegExp_ExactMatch(void* ptr, char* str){
	return static_cast<QRegExp*>(ptr)->exactMatch(QString(str));
}

int QRegExp_IndexIn(void* ptr, char* str, int offset, int caretMode){
	return static_cast<QRegExp*>(ptr)->indexIn(QString(str), offset, static_cast<QRegExp::CaretMode>(caretMode));
}

int QRegExp_IsEmpty(void* ptr){
	return static_cast<QRegExp*>(ptr)->isEmpty();
}

int QRegExp_IsMinimal(void* ptr){
	return static_cast<QRegExp*>(ptr)->isMinimal();
}

int QRegExp_IsValid(void* ptr){
	return static_cast<QRegExp*>(ptr)->isValid();
}

int QRegExp_LastIndexIn(void* ptr, char* str, int offset, int caretMode){
	return static_cast<QRegExp*>(ptr)->lastIndexIn(QString(str), offset, static_cast<QRegExp::CaretMode>(caretMode));
}

int QRegExp_MatchedLength(void* ptr){
	return static_cast<QRegExp*>(ptr)->matchedLength();
}

char* QRegExp_Pattern(void* ptr){
	return static_cast<QRegExp*>(ptr)->pattern().toUtf8().data();
}

int QRegExp_PatternSyntax(void* ptr){
	return static_cast<QRegExp*>(ptr)->patternSyntax();
}

int QRegExp_Pos(void* ptr, int nth){
	return static_cast<QRegExp*>(ptr)->pos(nth);
}

void QRegExp_SetCaseSensitivity(void* ptr, int cs){
	static_cast<QRegExp*>(ptr)->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QRegExp_SetMinimal(void* ptr, int minimal){
	static_cast<QRegExp*>(ptr)->setMinimal(minimal != 0);
}

void QRegExp_SetPattern(void* ptr, char* pattern){
	static_cast<QRegExp*>(ptr)->setPattern(QString(pattern));
}

void QRegExp_SetPatternSyntax(void* ptr, int syntax){
	static_cast<QRegExp*>(ptr)->setPatternSyntax(static_cast<QRegExp::PatternSyntax>(syntax));
}

void QRegExp_Swap(void* ptr, void* other){
	static_cast<QRegExp*>(ptr)->swap(*static_cast<QRegExp*>(other));
}

void QRegExp_DestroyQRegExp(void* ptr){
	static_cast<QRegExp*>(ptr)->~QRegExp();
}

