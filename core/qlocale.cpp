#include "qlocale.h"
#include <QModelIndex>
#include <QDateTime>
#include <QStringRef>
#include <QTime>
#include <QDate>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QLocale>
#include "_cgo_export.h"

class MyQLocale: public QLocale {
public:
};

QtObjectPtr QLocale_NewQLocale(){
	return new QLocale();
}

QtObjectPtr QLocale_NewQLocale3(int language, int country){
	return new QLocale(static_cast<QLocale::Language>(language), static_cast<QLocale::Country>(country));
}

QtObjectPtr QLocale_NewQLocale4(int language, int script, int country){
	return new QLocale(static_cast<QLocale::Language>(language), static_cast<QLocale::Script>(script), static_cast<QLocale::Country>(country));
}

QtObjectPtr QLocale_NewQLocale5(QtObjectPtr other){
	return new QLocale(*static_cast<QLocale*>(other));
}

QtObjectPtr QLocale_NewQLocale2(char* name){
	return new QLocale(QString(name));
}

char* QLocale_AmText(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->amText().toUtf8().data();
}

char* QLocale_Bcp47Name(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->bcp47Name().toUtf8().data();
}

int QLocale_Country(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->country();
}

char* QLocale_QLocale_CountryToString(int country){
	return QLocale::countryToString(static_cast<QLocale::Country>(country)).toUtf8().data();
}

char* QLocale_CreateSeparatedList(QtObjectPtr ptr, char* list){
	return static_cast<QLocale*>(ptr)->createSeparatedList(QString(list).split("|", QString::SkipEmptyParts)).toUtf8().data();
}

char* QLocale_CurrencySymbol(QtObjectPtr ptr, int format){
	return static_cast<QLocale*>(ptr)->currencySymbol(static_cast<QLocale::CurrencySymbolFormat>(format)).toUtf8().data();
}

char* QLocale_DateFormat(QtObjectPtr ptr, int format){
	return static_cast<QLocale*>(ptr)->dateFormat(static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_DateTimeFormat(QtObjectPtr ptr, int format){
	return static_cast<QLocale*>(ptr)->dateTimeFormat(static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_DayName(QtObjectPtr ptr, int day, int ty){
	return static_cast<QLocale*>(ptr)->dayName(day, static_cast<QLocale::FormatType>(ty)).toUtf8().data();
}

int QLocale_FirstDayOfWeek(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->firstDayOfWeek();
}

int QLocale_Language(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->language();
}

char* QLocale_QLocale_LanguageToString(int language){
	return QLocale::languageToString(static_cast<QLocale::Language>(language)).toUtf8().data();
}

int QLocale_MeasurementSystem(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->measurementSystem();
}

char* QLocale_MonthName(QtObjectPtr ptr, int month, int ty){
	return static_cast<QLocale*>(ptr)->monthName(month, static_cast<QLocale::FormatType>(ty)).toUtf8().data();
}

char* QLocale_Name(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->name().toUtf8().data();
}

char* QLocale_NativeCountryName(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->nativeCountryName().toUtf8().data();
}

char* QLocale_NativeLanguageName(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->nativeLanguageName().toUtf8().data();
}

int QLocale_NumberOptions(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->numberOptions();
}

char* QLocale_PmText(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->pmText().toUtf8().data();
}

char* QLocale_QuoteString(QtObjectPtr ptr, char* str, int style){
	return static_cast<QLocale*>(ptr)->quoteString(QString(str), static_cast<QLocale::QuotationStyle>(style)).toUtf8().data();
}

char* QLocale_QuoteString2(QtObjectPtr ptr, QtObjectPtr str, int style){
	return static_cast<QLocale*>(ptr)->quoteString(*static_cast<QStringRef*>(str), static_cast<QLocale::QuotationStyle>(style)).toUtf8().data();
}

int QLocale_Script(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->script();
}

char* QLocale_QLocale_ScriptToString(int script){
	return QLocale::scriptToString(static_cast<QLocale::Script>(script)).toUtf8().data();
}

void QLocale_QLocale_SetDefault(QtObjectPtr locale){
	QLocale::setDefault(*static_cast<QLocale*>(locale));
}

void QLocale_SetNumberOptions(QtObjectPtr ptr, int options){
	static_cast<QLocale*>(ptr)->setNumberOptions(static_cast<QLocale::NumberOption>(options));
}

char* QLocale_StandaloneDayName(QtObjectPtr ptr, int day, int ty){
	return static_cast<QLocale*>(ptr)->standaloneDayName(day, static_cast<QLocale::FormatType>(ty)).toUtf8().data();
}

char* QLocale_StandaloneMonthName(QtObjectPtr ptr, int month, int ty){
	return static_cast<QLocale*>(ptr)->standaloneMonthName(month, static_cast<QLocale::FormatType>(ty)).toUtf8().data();
}

int QLocale_TextDirection(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->textDirection();
}

char* QLocale_TimeFormat(QtObjectPtr ptr, int format){
	return static_cast<QLocale*>(ptr)->timeFormat(static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_ToCurrencyString6(QtObjectPtr ptr, int value, char* symbol){
	return static_cast<QLocale*>(ptr)->toCurrencyString(value, QString(symbol)).toUtf8().data();
}

int QLocale_ToInt(QtObjectPtr ptr, char* s, int ok){
	return static_cast<QLocale*>(ptr)->toInt(QString(s), NULL);
}

int QLocale_ToInt2(QtObjectPtr ptr, QtObjectPtr s, int ok){
	return static_cast<QLocale*>(ptr)->toInt(*static_cast<QStringRef*>(s), NULL);
}

char* QLocale_ToLower(QtObjectPtr ptr, char* str){
	return static_cast<QLocale*>(ptr)->toLower(QString(str)).toUtf8().data();
}

char* QLocale_ToString3(QtObjectPtr ptr, QtObjectPtr date, int format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QDate*>(date), static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_ToString2(QtObjectPtr ptr, QtObjectPtr date, char* format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QDate*>(date), QString(format)).toUtf8().data();
}

char* QLocale_ToString6(QtObjectPtr ptr, QtObjectPtr dateTime, int format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QDateTime*>(dateTime), static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_ToString7(QtObjectPtr ptr, QtObjectPtr dateTime, char* format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QDateTime*>(dateTime), QString(format)).toUtf8().data();
}

char* QLocale_ToString5(QtObjectPtr ptr, QtObjectPtr time, int format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QTime*>(time), static_cast<QLocale::FormatType>(format)).toUtf8().data();
}

char* QLocale_ToString4(QtObjectPtr ptr, QtObjectPtr time, char* format){
	return static_cast<QLocale*>(ptr)->toString(*static_cast<QTime*>(time), QString(format)).toUtf8().data();
}

char* QLocale_ToString12(QtObjectPtr ptr, int i){
	return static_cast<QLocale*>(ptr)->toString(i).toUtf8().data();
}

char* QLocale_ToUpper(QtObjectPtr ptr, char* str){
	return static_cast<QLocale*>(ptr)->toUpper(QString(str)).toUtf8().data();
}

char* QLocale_UiLanguages(QtObjectPtr ptr){
	return static_cast<QLocale*>(ptr)->uiLanguages().join("|").toUtf8().data();
}

void QLocale_DestroyQLocale(QtObjectPtr ptr){
	static_cast<QLocale*>(ptr)->~QLocale();
}

