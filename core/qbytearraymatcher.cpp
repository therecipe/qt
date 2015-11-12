#include "qbytearraymatcher.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QString>
#include <QByteArrayMatcher>
#include "_cgo_export.h"

class MyQByteArrayMatcher: public QByteArrayMatcher {
public:
};

void* QByteArrayMatcher_NewQByteArrayMatcher(){
	return new QByteArrayMatcher();
}

void* QByteArrayMatcher_NewQByteArrayMatcher2(void* pattern){
	return new QByteArrayMatcher(*static_cast<QByteArray*>(pattern));
}

void* QByteArrayMatcher_NewQByteArrayMatcher4(void* other){
	return new QByteArrayMatcher(*static_cast<QByteArrayMatcher*>(other));
}

void* QByteArrayMatcher_NewQByteArrayMatcher3(char* pattern, int length){
	return new QByteArrayMatcher(const_cast<const char*>(pattern), length);
}

int QByteArrayMatcher_IndexIn(void* ptr, void* ba, int from){
	return static_cast<QByteArrayMatcher*>(ptr)->indexIn(*static_cast<QByteArray*>(ba), from);
}

int QByteArrayMatcher_IndexIn2(void* ptr, char* str, int len, int from){
	return static_cast<QByteArrayMatcher*>(ptr)->indexIn(const_cast<const char*>(str), len, from);
}

void* QByteArrayMatcher_Pattern(void* ptr){
	return new QByteArray(static_cast<QByteArrayMatcher*>(ptr)->pattern());
}

void QByteArrayMatcher_SetPattern(void* ptr, void* pattern){
	static_cast<QByteArrayMatcher*>(ptr)->setPattern(*static_cast<QByteArray*>(pattern));
}

void QByteArrayMatcher_DestroyQByteArrayMatcher(void* ptr){
	static_cast<QByteArrayMatcher*>(ptr)->~QByteArrayMatcher();
}

