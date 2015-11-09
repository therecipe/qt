#include "qstringref.h"
#include <QLatin1String>
#include <QChar>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QStringRef>
#include "_cgo_export.h"

class MyQStringRef: public QStringRef {
public:
};

void* QStringRef_Left(void* ptr, int n){
	return new QStringRef(static_cast<QStringRef*>(ptr)->left(n));
}

void* QStringRef_Mid(void* ptr, int position, int n){
	return new QStringRef(static_cast<QStringRef*>(ptr)->mid(position, n));
}

void* QStringRef_Right(void* ptr, int n){
	return new QStringRef(static_cast<QStringRef*>(ptr)->right(n));
}

void* QStringRef_AppendTo(void* ptr, char* stri){
	return new QStringRef(static_cast<QStringRef*>(ptr)->appendTo(new QString(stri)));
}

void* QStringRef_Begin(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->begin());
}

void* QStringRef_Cbegin(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->cbegin());
}

void* QStringRef_Cend(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->cend());
}

void QStringRef_Clear(void* ptr){
	static_cast<QStringRef*>(ptr)->clear();
}

int QStringRef_QStringRef_Compare3(void* s1, void* s2, int cs){
	return QStringRef::compare(*static_cast<QStringRef*>(s1), *static_cast<QLatin1String*>(s2), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_QStringRef_Compare(void* s1, char* s2, int cs){
	return QStringRef::compare(*static_cast<QStringRef*>(s1), QString(s2), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_QStringRef_Compare2(void* s1, void* s2, int cs){
	return QStringRef::compare(*static_cast<QStringRef*>(s1), *static_cast<QStringRef*>(s2), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Compare6(void* ptr, void* other, int cs){
	return static_cast<QStringRef*>(ptr)->compare(*static_cast<QLatin1String*>(other), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Compare4(void* ptr, char* other, int cs){
	return static_cast<QStringRef*>(ptr)->compare(QString(other), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Compare5(void* ptr, void* other, int cs){
	return static_cast<QStringRef*>(ptr)->compare(*static_cast<QStringRef*>(other), static_cast<Qt::CaseSensitivity>(cs));
}

void* QStringRef_ConstData(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->constData());
}

int QStringRef_Contains2(void* ptr, void* ch, int cs){
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Contains4(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QLatin1String*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Contains(void* ptr, char* str, int cs){
	return static_cast<QStringRef*>(ptr)->contains(QString(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Contains3(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count(void* ptr){
	return static_cast<QStringRef*>(ptr)->count();
}

int QStringRef_Count3(void* ptr, void* ch, int cs){
	return static_cast<QStringRef*>(ptr)->count(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count2(void* ptr, char* str, int cs){
	return static_cast<QStringRef*>(ptr)->count(QString(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count4(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->count(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

void* QStringRef_Data(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->data());
}

void* QStringRef_End(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->end());
}

int QStringRef_EndsWith2(void* ptr, void* ch, int cs){
	return static_cast<QStringRef*>(ptr)->endsWith(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_EndsWith3(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->endsWith(*static_cast<QLatin1String*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_EndsWith(void* ptr, char* str, int cs){
	return static_cast<QStringRef*>(ptr)->endsWith(QString(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_EndsWith4(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->endsWith(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf3(void* ptr, void* ch, int from, int cs){
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QChar*>(ch), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf2(void* ptr, void* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QLatin1String*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf(void* ptr, char* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->indexOf(QString(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IndexOf4(void* ptr, void* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->indexOf(*static_cast<QStringRef*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_IsEmpty(void* ptr){
	return static_cast<QStringRef*>(ptr)->isEmpty();
}

int QStringRef_IsNull(void* ptr){
	return static_cast<QStringRef*>(ptr)->isNull();
}

int QStringRef_LastIndexOf2(void* ptr, void* ch, int from, int cs){
	return static_cast<QStringRef*>(ptr)->lastIndexOf(*static_cast<QChar*>(ch), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_LastIndexOf3(void* ptr, void* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->lastIndexOf(*static_cast<QLatin1String*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_LastIndexOf(void* ptr, char* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->lastIndexOf(QString(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_LastIndexOf4(void* ptr, void* str, int from, int cs){
	return static_cast<QStringRef*>(ptr)->lastIndexOf(*static_cast<QStringRef*>(str), from, static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Length(void* ptr){
	return static_cast<QStringRef*>(ptr)->length();
}

int QStringRef_QStringRef_LocaleAwareCompare(void* s1, char* s2){
	return QStringRef::localeAwareCompare(*static_cast<QStringRef*>(s1), QString(s2));
}

int QStringRef_QStringRef_LocaleAwareCompare2(void* s1, void* s2){
	return QStringRef::localeAwareCompare(*static_cast<QStringRef*>(s1), *static_cast<QStringRef*>(s2));
}

int QStringRef_LocaleAwareCompare3(void* ptr, char* other){
	return static_cast<QStringRef*>(ptr)->localeAwareCompare(QString(other));
}

int QStringRef_LocaleAwareCompare4(void* ptr, void* other){
	return static_cast<QStringRef*>(ptr)->localeAwareCompare(*static_cast<QStringRef*>(other));
}

int QStringRef_Position(void* ptr){
	return static_cast<QStringRef*>(ptr)->position();
}

int QStringRef_Size(void* ptr){
	return static_cast<QStringRef*>(ptr)->size();
}

int QStringRef_StartsWith4(void* ptr, void* ch, int cs){
	return static_cast<QStringRef*>(ptr)->startsWith(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_StartsWith2(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->startsWith(*static_cast<QLatin1String*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_StartsWith(void* ptr, char* str, int cs){
	return static_cast<QStringRef*>(ptr)->startsWith(QString(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_StartsWith3(void* ptr, void* str, int cs){
	return static_cast<QStringRef*>(ptr)->startsWith(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

char* QStringRef_String(void* ptr){
	return static_cast<QStringRef*>(ptr)->string()->toUtf8().data();
}

int QStringRef_ToInt(void* ptr, int ok, int base){
	return static_cast<QStringRef*>(ptr)->toInt(NULL, base);
}

void* QStringRef_ToLatin1(void* ptr){
	return new QByteArray(static_cast<QStringRef*>(ptr)->toLatin1());
}

void* QStringRef_ToLocal8Bit(void* ptr){
	return new QByteArray(static_cast<QStringRef*>(ptr)->toLocal8Bit());
}

char* QStringRef_ToString(void* ptr){
	return static_cast<QStringRef*>(ptr)->toString().toUtf8().data();
}

void* QStringRef_ToUtf8(void* ptr){
	return new QByteArray(static_cast<QStringRef*>(ptr)->toUtf8());
}

void* QStringRef_Trimmed(void* ptr){
	return new QStringRef(static_cast<QStringRef*>(ptr)->trimmed());
}

void* QStringRef_Unicode(void* ptr){
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->unicode());
}

void QStringRef_DestroyQStringRef(void* ptr){
	static_cast<QStringRef*>(ptr)->~QStringRef();
}

