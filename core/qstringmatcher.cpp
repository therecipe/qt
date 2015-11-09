#include "qstringmatcher.h"
#include <QModelIndex>
#include <QChar>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStringMatcher>
#include "_cgo_export.h"

class MyQStringMatcher: public QStringMatcher {
public:
};

void* QStringMatcher_NewQStringMatcher3(void* uc, int length, int cs){
	return new QStringMatcher(static_cast<QChar*>(uc), length, static_cast<Qt::CaseSensitivity>(cs));
}

char* QStringMatcher_Pattern(void* ptr){
	return static_cast<QStringMatcher*>(ptr)->pattern().toUtf8().data();
}

void* QStringMatcher_NewQStringMatcher(){
	return new QStringMatcher();
}

void* QStringMatcher_NewQStringMatcher2(char* pattern, int cs){
	return new QStringMatcher(QString(pattern), static_cast<Qt::CaseSensitivity>(cs));
}

void* QStringMatcher_NewQStringMatcher4(void* other){
	return new QStringMatcher(*static_cast<QStringMatcher*>(other));
}

int QStringMatcher_CaseSensitivity(void* ptr){
	return static_cast<QStringMatcher*>(ptr)->caseSensitivity();
}

int QStringMatcher_IndexIn2(void* ptr, void* str, int length, int from){
	return static_cast<QStringMatcher*>(ptr)->indexIn(static_cast<QChar*>(str), length, from);
}

int QStringMatcher_IndexIn(void* ptr, char* str, int from){
	return static_cast<QStringMatcher*>(ptr)->indexIn(QString(str), from);
}

void QStringMatcher_SetCaseSensitivity(void* ptr, int cs){
	static_cast<QStringMatcher*>(ptr)->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QStringMatcher_SetPattern(void* ptr, char* pattern){
	static_cast<QStringMatcher*>(ptr)->setPattern(QString(pattern));
}

void QStringMatcher_DestroyQStringMatcher(void* ptr){
	static_cast<QStringMatcher*>(ptr)->~QStringMatcher();
}

