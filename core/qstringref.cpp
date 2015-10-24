#include "qstringref.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLatin1String>
#include <QChar>
#include <QString>
#include <QStringRef>
#include "_cgo_export.h"

class MyQStringRef: public QStringRef {
public:
};

QtObjectPtr QStringRef_Begin(QtObjectPtr ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->begin());
}

QtObjectPtr QStringRef_Cbegin(QtObjectPtr ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->cbegin());
}

QtObjectPtr QStringRef_Cend(QtObjectPtr ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->cend());
}

void QStringRef_Clear(QtObjectPtr ptr){
	static_cast<QStringRef*>(ptr)->clear();
}

int QStringRef_QStringRef_Compare3(QtObjectPtr s1, QtObjectPtr s2, int cs){
	return QStringRef::compare(*static_cast<QStringRef*>(s1), *static_cast<QLatin1String*>(s2), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_QStringRef_Compare(QtObjectPtr s1, char* s2, int cs){
	return QStringRef::compare(*static_cast<QStringRef*>(s1), QString(s2), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_QStringRef_Compare2(QtObjectPtr s1, QtObjectPtr s2, int cs){
	return QStringRef::compare(*static_cast<QStringRef*>(s1), *static_cast<QStringRef*>(s2), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Compare6(QtObjectPtr ptr, QtObjectPtr other, int cs){
	return static_cast<QStringRef*>(ptr)->compare(*static_cast<QLatin1String*>(other), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Compare4(QtObjectPtr ptr, char* other, int cs){
	return static_cast<QStringRef*>(ptr)->compare(QString(other), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Compare5(QtObjectPtr ptr, QtObjectPtr other, int cs){
	return static_cast<QStringRef*>(ptr)->compare(*static_cast<QStringRef*>(other), static_cast<Qt::CaseSensitivity>(cs));
}

QtObjectPtr QStringRef_ConstData(QtObjectPtr ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->constData());
}

int QStringRef_Contains2(QtObjectPtr ptr, QtObjectPtr ch, int cs){
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Contains4(QtObjectPtr ptr, QtObjectPtr str, int cs){
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QLatin1String*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Contains(QtObjectPtr ptr, char* str, int cs){
	return static_cast<QStringRef*>(ptr)->contains(QString(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Contains3(QtObjectPtr ptr, QtObjectPtr str, int cs){
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count(QtObjectPtr ptr){
	return static_cast<QStringRef*>(ptr)->count();
}

int QStringRef_Count3(QtObjectPtr ptr, QtObjectPtr ch, int cs){
	return static_cast<QStringRef*>(ptr)->count(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count2(QtObjectPtr ptr, char* str, int cs){
	return static_cast<QStringRef*>(ptr)->count(QString(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count4(QtObjectPtr ptr, QtObjectPtr str, int cs){
	return static_cast<QStringRef*>(ptr)->count(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

QtObjectPtr QStringRef_Data(QtObjectPtr ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->data());
}

QtObjectPtr QStringRef_End(QtObjectPtr ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->end());
}

int QStringRef_EndsWith2(QtObjectPtr ptr, QtObjectPtr ch, int cs){
	return static_cast<QStringRef*>(ptr)->endsWith(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_EndsWith3(QtObjectPtr ptr, QtObjectPtr str, int cs){
	return static_cast<QStringRef*>(ptr)->endsWith(*static_cast<QLatin1String*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_EndsWith(QtObjectPtr ptr, char* str, int cs){
	return static_cast<QStringRef*>(ptr)->endsWith(QString(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_EndsWith4(QtObjectPtr ptr, QtObjectPtr str, int cs){
	return static_cast<QStringRef*>(ptr)->endsWith(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf3(QtObjectPtr ptr, QtObjectPtr ch, int from, int cs){
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QChar*>(ch), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf2(QtObjectPtr ptr, QtObjectPtr str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QLatin1String*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf(QtObjectPtr ptr, char* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->indexOf(QString(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf4(QtObjectPtr ptr, QtObjectPtr str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QStringRef*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IsEmpty(QtObjectPtr ptr){
	return static_cast<QStringRef*>(ptr)->isEmpty();
}

int QStringRef_IsNull(QtObjectPtr ptr){
	return static_cast<QStringRef*>(ptr)->isNull();
}

int QStringRef_LastIndexOf2(QtObjectPtr ptr, QtObjectPtr ch, int from, int cs){
	return static_cast<QStringRef*>(ptr)->lastIndexOf(*static_cast<QChar*>(ch), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_LastIndexOf3(QtObjectPtr ptr, QtObjectPtr str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->lastIndexOf(*static_cast<QLatin1String*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_LastIndexOf(QtObjectPtr ptr, char* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->lastIndexOf(QString(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_LastIndexOf4(QtObjectPtr ptr, QtObjectPtr str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->lastIndexOf(*static_cast<QStringRef*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Length(QtObjectPtr ptr){
	return static_cast<QStringRef*>(ptr)->length();
}

int QStringRef_QStringRef_LocaleAwareCompare(QtObjectPtr s1, char* s2){
	return QStringRef::localeAwareCompare(*static_cast<QStringRef*>(s1), QString(s2));
}

int QStringRef_QStringRef_LocaleAwareCompare2(QtObjectPtr s1, QtObjectPtr s2){
	return QStringRef::localeAwareCompare(*static_cast<QStringRef*>(s1), *static_cast<QStringRef*>(s2));
}

int QStringRef_LocaleAwareCompare3(QtObjectPtr ptr, char* other){
	return static_cast<QStringRef*>(ptr)->localeAwareCompare(QString(other));
}

int QStringRef_LocaleAwareCompare4(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QStringRef*>(ptr)->localeAwareCompare(*static_cast<QStringRef*>(other));
}

int QStringRef_Position(QtObjectPtr ptr){
	return static_cast<QStringRef*>(ptr)->position();
}

int QStringRef_Size(QtObjectPtr ptr){
	return static_cast<QStringRef*>(ptr)->size();
}

int QStringRef_StartsWith4(QtObjectPtr ptr, QtObjectPtr ch, int cs){
	return static_cast<QStringRef*>(ptr)->startsWith(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_StartsWith2(QtObjectPtr ptr, QtObjectPtr str, int cs){
	return static_cast<QStringRef*>(ptr)->startsWith(*static_cast<QLatin1String*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_StartsWith(QtObjectPtr ptr, char* str, int cs){
	return static_cast<QStringRef*>(ptr)->startsWith(QString(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_StartsWith3(QtObjectPtr ptr, QtObjectPtr str, int cs){
	return static_cast<QStringRef*>(ptr)->startsWith(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

char* QStringRef_String(QtObjectPtr ptr){
	return static_cast<QStringRef*>(ptr)->string()->toUtf8().data();
}

int QStringRef_ToInt(QtObjectPtr ptr, int ok, int base){
	return static_cast<QStringRef*>(ptr)->toInt(NULL, base);
}

char* QStringRef_ToString(QtObjectPtr ptr){
	return static_cast<QStringRef*>(ptr)->toString().toUtf8().data();
}

QtObjectPtr QStringRef_Unicode(QtObjectPtr ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->unicode());
}

void QStringRef_DestroyQStringRef(QtObjectPtr ptr){
	static_cast<QStringRef*>(ptr)->~QStringRef();
}

