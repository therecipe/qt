#include "qlocale.h"
#include <QStringRef>
#include <QDate>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDateTime>
#include <QTime>
#include <QLocale>
#include "_cgo_export.h"

class MyQLocale: public QLocale {
public:
};

void* QLocale_NewQLocale(){
	return new QLocale();
}

void* QLocale_NewQLocale3(int language, int country){
	return new QLocale(static_cast<QLocale::Language>(language), static_cast<QLocale::Country>(country));
}

void* QLocale_NewQLocale4(int language, int script, int country){
	return new QLocale(static_cast<QLocale::Language>(language), static_cast<QLocale::Script>(script), static_cast<QLocale::Country>(country));
}

void* QLocale_NewQLocale5(void* other){
	return new QLocale(*static_cast<QLocale*>(other));
}

void* QLocale_NewQLocale2(char* name){
	return new QLocale(QString(name));
}

char* QLocale_AmText(void* ptr){
	return static_cast<QLocale*>(ptr)->amText().toUtf8().data();
}

char* QLocale_Bcp47Name(void* ptr){
	return static_cast<QLocale*>(ptr)->bcp47Name().toUtf8().data();
}

int QLocale_Country(void* ptr){
	return static_cast<QLocale*>(ptr)->country();
}

char* QLocale_QLocale_CountryToString(int country){
	return QLocale::countryToString(static_cast<QLocale::Country>(country)).toUtf8().data();
}

char* QLocale_CreateSeparatedList(void* ptr, char* list){
	return static_cast<QLocale*>(ptr)->createSeparatedList(QString(list).split("|", QString::SkipEmptyParts)).toUtf8().data();
}

char* QLocale_CurrencySymbol(void* ptr, int format){
	return static_cast<QLocale*>(ptr)->currencySymbol(static_cast<QLocale::CurrencySymbolFormat>(format)).toUtf8().data();
}

char* QLocale_DateFormat(void* ptr, int format){
	return static_cast<QLocale*>(ptr)->dateFormat(static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_DateTimeFormat(void* ptr, int format){
	return static_cast<QLocale*>(ptr)->dateTimeFormat(static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_DayName(void* ptr, int day, int ty){
	return static_cast<QLocale*>(ptr)->dayName(day, static_cast<QLocale::FormatType>(ty)).toUtf8().data();
}

int QLocale_FirstDayOfWeek(void* ptr){
	return static_cast<QLocale*>(ptr)->firstDayOfWeek();
}

int QLocale_Language(void* ptr){
	return static_cast<QLocale*>(ptr)->language();
}

char* QLocale_QLocale_LanguageToString(int language){
	return QLocale::languageToString(static_cast<QLocale::Language>(language)).toUtf8().data();
}

int QLocale_MeasurementSystem(void* ptr){
	return static_cast<QLocale*>(ptr)->measurementSystem();
}

char* QLocale_MonthName(void* ptr, int month, int ty){
	return static_cast<QLocale*>(ptr)->monthName(month, static_cast<QLocale::FormatType>(ty)).toUtf8().data();
}

char* QLocale_Name(void* ptr){
	return static_cast<QLocale*>(ptr)->name().toUtf8().data();
}

char* QLocale_NativeCountryName(void* ptr){
	return static_cast<QLocale*>(ptr)->nativeCountryName().toUtf8().data();
}

char* QLocale_NativeLanguageName(void* ptr){
	return static_cast<QLocale*>(ptr)->nativeLanguageName().toUtf8().data();
}

int QLocale_NumberOptions(void* ptr){
	return static_cast<QLocale*>(ptr)->numberOptions();
}

char* QLocale_PmText(void* ptr){
	return static_cast<QLocale*>(ptr)->pmText().toUtf8().data();
}

char* QLocale_QuoteString(void* ptr, char* str, int style){
	return static_cast<QLocale*>(ptr)->quoteString(QString(str), static_cast<QLocale::QuotationStyle>(style)).toUtf8().data();
}

char* QLocale_QuoteString2(void* ptr, void* str, int style){
	return static_cast<QLocale*>(ptr)->quoteString(*static_cast<QStringRef*>(str), static_cast<QLocale::QuotationStyle>(style)).toUtf8().data();
}

int QLocale_Script(void* ptr){
	return static_cast<QLocale*>(ptr)->script();
}

char* QLocale_QLocale_ScriptToString(int script){
	return QLocale::scriptToString(static_cast<QLocale::Script>(script)).toUtf8().data();
}

void QLocale_QLocale_SetDefault(void* locale){
	QLocale::setDefault(*static_cast<QLocale*>(locale));
}

void QLocale_SetNumberOptions(void* ptr, int options){
	static_cast<QLocale*>(ptr)->setNumberOptions(static_cast<QLocale::NumberOption>(options));
}

char* QLocale_StandaloneDayName(void* ptr, int day, int ty){
	return static_cast<QLocale*>(ptr)->standaloneDayName(day, static_cast<QLocale::FormatType>(ty)).toUtf8().data();
}

char* QLocale_StandaloneMonthName(void* ptr, int month, int ty){
	return static_cast<QLocale*>(ptr)->standaloneMonthName(month, static_cast<QLocale::FormatType>(ty)).toUtf8().data();
}

int QLocale_TextDirection(void* ptr){
	return static_cast<QLocale*>(ptr)->textDirection();
}

char* QLocale_TimeFormat(void* ptr, int format){
	return static_cast<QLocale*>(ptr)->timeFormat(static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_ToCurrencyString6(void* ptr, int value, char* symbol){
	return static_cast<QLocale*>(ptr)->toCurrencyString(value, QString(symbol)).toUtf8().data();
}

void* QLocale_ToDateTime(void* ptr, char* stri, int format){
	return new QDateTime(static_cast<QLocale*>(ptr)->toDateTime(QString(stri), static_cast<QLocale::FormatType>(format)));
}

void* QLocale_ToDateTime2(void* ptr, char* stri, char* format){
	return new QDateTime(static_cast<QLocale*>(ptr)->toDateTime(QString(stri), QString(format)));
}

int QLocale_ToInt(void* ptr, char* s, int ok){
	return static_cast<QLocale*>(ptr)->toInt(QString(s), NULL);
}

int QLocale_ToInt2(void* ptr, void* s, int ok){
	return static_cast<QLocale*>(ptr)->toInt(*static_cast<QStringRef*>(s), NULL);
}

char* QLocale_ToLower(void* ptr, char* str){
	return static_cast<QLocale*>(ptr)->toLower(QString(str)).toUtf8().data();
}

char* QLocale_ToString3(void* ptr, void* date, int format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QDate*>(date), static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_ToString2(void* ptr, void* date, char* format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QDate*>(date), QString(format)).toUtf8().data();
}

char* QLocale_ToString6(void* ptr, void* dateTime, int format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QDateTime*>(dateTime), static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_ToString7(void* ptr, void* dateTime, char* format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QDateTime*>(dateTime), QString(format)).toUtf8().data();
}

char* QLocale_ToString5(void* ptr, void* time, int format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QTime*>(time), static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_ToString4(void* ptr, void* time, char* format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QTime*>(time), QString(format)).toUtf8().data();
}

char* QLocale_ToString12(void* ptr, int i){
	return static_cast<QLocale*>(ptr)->toString(i).toUtf8().data();
}

char* QLocale_ToUpper(void* ptr, char* str){
	return static_cast<QLocale*>(ptr)->toUpper(QString(str)).toUtf8().data();
}

char* QLocale_UiLanguages(void* ptr){
	return static_cast<QLocale*>(ptr)->uiLanguages().join("|").toUtf8().data();
}

void QLocale_DestroyQLocale(void* ptr){
	static_cast<QLocale*>(ptr)->~QLocale();
}

