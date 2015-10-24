#include "qstringmatcher.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QChar>
#include <QStringMatcher>
#include "_cgo_export.h"

class MyQStringMatcher: public QStringMatcher {
public:
};

QtObjectPtr QStringMatcher_NewQStringMatcher3(QtObjectPtr uc, int length, int cs){
	return new QStringMatcher(static_cast<QChar*>(uc), length, static_cast<Qt::CaseSensitivity>(cs));
}

char* QStringMatcher_Pattern(QtObjectPtr ptr){
	return static_cast<QStringMatcher*>(ptr)->pattern().toUtf8().data();
}

QtObjectPtr QStringMatcher_NewQStringMatcher(){
	return new QStringMatcher();
}

QtObjectPtr QStringMatcher_NewQStringMatcher2(char* pattern, int cs){
	return new QStringMatcher(QString(pattern), static_cast<Qt::CaseSensitivity>(cs));
}

QtObjectPtr QStringMatcher_NewQStringMatcher4(QtObjectPtr other){
	return new QStringMatcher(*static_cast<QStringMatcher*>(other));
}

int QStringMatcher_CaseSensitivity(QtObjectPtr ptr){
	return static_cast<QStringMatcher*>(ptr)->caseSensitivity();
}

int QStringMatcher_IndexIn2(QtObjectPtr ptr, QtObjectPtr str, int length, int from){
	return static_cast<QStringMatcher*>(ptr)->indexIn(static_cast<QChar*>(str), length, from);
}

int QStringMatcher_IndexIn(QtObjectPtr ptr, char* str, int from){
	return static_cast<QStringMatcher*>(ptr)->indexIn(QString(str), from);
}

void QStringMatcher_SetCaseSensitivity(QtObjectPtr ptr, int cs){
	static_cast<QStringMatcher*>(ptr)->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QStringMatcher_SetPattern(QtObjectPtr ptr, char* pattern){
	static_cast<QStringMatcher*>(ptr)->setPattern(QString(pattern));
}

void QStringMatcher_DestroyQStringMatcher(QtObjectPtr ptr){
	static_cast<QStringMatcher*>(ptr)->~QStringMatcher();
}

