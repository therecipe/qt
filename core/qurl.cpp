#include "qurl.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QUrlQuery>
#include <QUrl>
#include "_cgo_export.h"

class MyQUrl: public QUrl {
public:
};

void* QUrl_NewQUrl(){
	return new QUrl();
}

void* QUrl_NewQUrl4(void* other){
	return new QUrl(*static_cast<QUrl*>(other));
}

void* QUrl_NewQUrl3(char* url, int parsingMode){
	return new QUrl(QString(url), static_cast<QUrl::ParsingMode>(parsingMode));
}

void* QUrl_NewQUrl2(void* other){
	return new QUrl(*static_cast<QUrl*>(other));
}

char* QUrl_Authority(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->authority(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

void QUrl_Clear(void* ptr){
	static_cast<QUrl*>(ptr)->clear();
}

char* QUrl_ErrorString(void* ptr){
	return static_cast<QUrl*>(ptr)->errorString().toUtf8().data();
}

char* QUrl_FileName(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->fileName(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_Fragment(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->fragment(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_QUrl_FromAce(void* domain){
	return QUrl::fromAce(*static_cast<QByteArray*>(domain)).toUtf8().data();
}

char* QUrl_QUrl_FromPercentEncoding(void* input){
	return QUrl::fromPercentEncoding(*static_cast<QByteArray*>(input)).toUtf8().data();
}

int QUrl_HasFragment(void* ptr){
	return static_cast<QUrl*>(ptr)->hasFragment();
}

int QUrl_HasQuery(void* ptr){
	return static_cast<QUrl*>(ptr)->hasQuery();
}

char* QUrl_Host(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->host(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_QUrl_IdnWhitelist(){
	return QUrl::idnWhitelist().join("|").toUtf8().data();
}

int QUrl_IsEmpty(void* ptr){
	return static_cast<QUrl*>(ptr)->isEmpty();
}

int QUrl_IsLocalFile(void* ptr){
	return static_cast<QUrl*>(ptr)->isLocalFile();
}

int QUrl_IsParentOf(void* ptr, void* childUrl){
	return static_cast<QUrl*>(ptr)->isParentOf(*static_cast<QUrl*>(childUrl));
}

int QUrl_IsRelative(void* ptr){
	return static_cast<QUrl*>(ptr)->isRelative();
}

int QUrl_IsValid(void* ptr){
	return static_cast<QUrl*>(ptr)->isValid();
}

int QUrl_Matches(void* ptr, void* url, int options){
	return static_cast<QUrl*>(ptr)->matches(*static_cast<QUrl*>(url), static_cast<QUrl::UrlFormattingOption>(options));
}

char* QUrl_Password(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->password(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_Path(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->path(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

int QUrl_Port(void* ptr, int defaultPort){
	return static_cast<QUrl*>(ptr)->port(defaultPort);
}

char* QUrl_Query(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->query(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_Scheme(void* ptr){
	return static_cast<QUrl*>(ptr)->scheme().toUtf8().data();
}

void QUrl_SetAuthority(void* ptr, char* authority, int mode){
	static_cast<QUrl*>(ptr)->setAuthority(QString(authority), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_SetFragment(void* ptr, char* fragment, int mode){
	static_cast<QUrl*>(ptr)->setFragment(QString(fragment), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_SetHost(void* ptr, char* host, int mode){
	static_cast<QUrl*>(ptr)->setHost(QString(host), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_QUrl_SetIdnWhitelist(char* list){
	QUrl::setIdnWhitelist(QString(list).split("|", QString::SkipEmptyParts));
}

void QUrl_SetPassword(void* ptr, char* password, int mode){
	static_cast<QUrl*>(ptr)->setPassword(QString(password), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_SetPath(void* ptr, char* path, int mode){
	static_cast<QUrl*>(ptr)->setPath(QString(path), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_SetPort(void* ptr, int port){
	static_cast<QUrl*>(ptr)->setPort(port);
}

void QUrl_SetQuery(void* ptr, char* query, int mode){
	static_cast<QUrl*>(ptr)->setQuery(QString(query), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_SetQuery2(void* ptr, void* query){
	static_cast<QUrl*>(ptr)->setQuery(*static_cast<QUrlQuery*>(query));
}

void QUrl_SetScheme(void* ptr, char* scheme){
	static_cast<QUrl*>(ptr)->setScheme(QString(scheme));
}

void QUrl_SetUrl(void* ptr, char* url, int parsingMode){
	static_cast<QUrl*>(ptr)->setUrl(QString(url), static_cast<QUrl::ParsingMode>(parsingMode));
}

void QUrl_SetUserInfo(void* ptr, char* userInfo, int mode){
	static_cast<QUrl*>(ptr)->setUserInfo(QString(userInfo), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_SetUserName(void* ptr, char* userName, int mode){
	static_cast<QUrl*>(ptr)->setUserName(QString(userName), static_cast<QUrl::ParsingMode>(mode));
}

void QUrl_Swap(void* ptr, void* other){
	static_cast<QUrl*>(ptr)->swap(*static_cast<QUrl*>(other));
}

void* QUrl_QUrl_ToAce(char* domain){
	return new QByteArray(QUrl::toAce(QString(domain)));
}

char* QUrl_ToDisplayString(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->toDisplayString(static_cast<QUrl::UrlFormattingOption>(options)).toUtf8().data();
}

void* QUrl_ToEncoded(void* ptr, int options){
	return new QByteArray(static_cast<QUrl*>(ptr)->toEncoded(static_cast<QUrl::UrlFormattingOption>(options)));
}

char* QUrl_ToLocalFile(void* ptr){
	return static_cast<QUrl*>(ptr)->toLocalFile().toUtf8().data();
}

void* QUrl_QUrl_ToPercentEncoding(char* input, void* exclude, void* include){
	return new QByteArray(QUrl::toPercentEncoding(QString(input), *static_cast<QByteArray*>(exclude), *static_cast<QByteArray*>(include)));
}

char* QUrl_ToString(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->toString(static_cast<QUrl::UrlFormattingOption>(options)).toUtf8().data();
}

char* QUrl_TopLevelDomain(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->topLevelDomain(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_Url(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->url(static_cast<QUrl::UrlFormattingOption>(options)).toUtf8().data();
}

char* QUrl_UserInfo(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->userInfo(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

char* QUrl_UserName(void* ptr, int options){
	return static_cast<QUrl*>(ptr)->userName(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8().data();
}

void QUrl_DestroyQUrl(void* ptr){
	static_cast<QUrl*>(ptr)->~QUrl();
}

