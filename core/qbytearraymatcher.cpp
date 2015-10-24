#include "qbytearraymatcher.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QByteArrayMatcher>
#include "_cgo_export.h"

class MyQByteArrayMatcher: public QByteArrayMatcher {
public:
};

QtObjectPtr QByteArrayMatcher_NewQByteArrayMatcher(){
	return new QByteArrayMatcher();
}

QtObjectPtr QByteArrayMatcher_NewQByteArrayMatcher2(QtObjectPtr pattern){
	return new QByteArrayMatcher(*static_cast<QByteArray*>(pattern));
}

QtObjectPtr QByteArrayMatcher_NewQByteArrayMatcher4(QtObjectPtr other){
	return new QByteArrayMatcher(*static_cast<QByteArrayMatcher*>(other));
}

QtObjectPtr QByteArrayMatcher_NewQByteArrayMatcher3(char* pattern, int length){
	return new QByteArrayMatcher(const_cast<const char*>(pattern), length);
}

int QByteArrayMatcher_IndexIn(QtObjectPtr ptr, QtObjectPtr ba, int from){
	return static_cast<QByteArrayMatcher*>(ptr)->indexIn(*static_cast<QByteArray*>(ba), from);
}

int QByteArrayMatcher_IndexIn2(QtObjectPtr ptr, char* str, int len, int from){
	return static_cast<QByteArrayMatcher*>(ptr)->indexIn(const_cast<const char*>(str), len, from);
}

void QByteArrayMatcher_SetPattern(QtObjectPtr ptr, QtObjectPtr pattern){
	static_cast<QByteArrayMatcher*>(ptr)->setPattern(*static_cast<QByteArray*>(pattern));
}

void QByteArrayMatcher_DestroyQByteArrayMatcher(QtObjectPtr ptr){
	static_cast<QByteArrayMatcher*>(ptr)->~QByteArrayMatcher();
}

