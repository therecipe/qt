#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QLocale_NewQLocale();
QtObjectPtr QLocale_NewQLocale3(int language, int country);
QtObjectPtr QLocale_NewQLocale4(int language, int script, int country);
QtObjectPtr QLocale_NewQLocale5(QtObjectPtr other);
QtObjectPtr QLocale_NewQLocale2(char* name);
char* QLocale_AmText(QtObjectPtr ptr);
char* QLocale_Bcp47Name(QtObjectPtr ptr);
int QLocale_Country(QtObjectPtr ptr);
char* QLocale_QLocale_CountryToString(int country);
char* QLocale_CreateSeparatedList(QtObjectPtr ptr, char* list);
char* QLocale_CurrencySymbol(QtObjectPtr ptr, int format);
char* QLocale_DateFormat(QtObjectPtr ptr, int format);
char* QLocale_DateTimeFormat(QtObjectPtr ptr, int format);
char* QLocale_DayName(QtObjectPtr ptr, int day, int ty);
int QLocale_FirstDayOfWeek(QtObjectPtr ptr);
int QLocale_Language(QtObjectPtr ptr);
char* QLocale_QLocale_LanguageToString(int language);
int QLocale_MeasurementSystem(QtObjectPtr ptr);
char* QLocale_MonthName(QtObjectPtr ptr, int month, int ty);
char* QLocale_Name(QtObjectPtr ptr);
char* QLocale_NativeCountryName(QtObjectPtr ptr);
char* QLocale_NativeLanguageName(QtObjectPtr ptr);
int QLocale_NumberOptions(QtObjectPtr ptr);
char* QLocale_PmText(QtObjectPtr ptr);
char* QLocale_QuoteString(QtObjectPtr ptr, char* str, int style);
char* QLocale_QuoteString2(QtObjectPtr ptr, QtObjectPtr str, int style);
int QLocale_Script(QtObjectPtr ptr);
char* QLocale_QLocale_ScriptToString(int script);
void QLocale_QLocale_SetDefault(QtObjectPtr locale);
void QLocale_SetNumberOptions(QtObjectPtr ptr, int options);
char* QLocale_StandaloneDayName(QtObjectPtr ptr, int day, int ty);
char* QLocale_StandaloneMonthName(QtObjectPtr ptr, int month, int ty);
int QLocale_TextDirection(QtObjectPtr ptr);
char* QLocale_TimeFormat(QtObjectPtr ptr, int format);
char* QLocale_ToCurrencyString6(QtObjectPtr ptr, int value, char* symbol);
int QLocale_ToInt(QtObjectPtr ptr, char* s, int ok);
int QLocale_ToInt2(QtObjectPtr ptr, QtObjectPtr s, int ok);
char* QLocale_ToLower(QtObjectPtr ptr, char* str);
char* QLocale_ToString3(QtObjectPtr ptr, QtObjectPtr date, int format);
char* QLocale_ToString2(QtObjectPtr ptr, QtObjectPtr date, char* format);
char* QLocale_ToString6(QtObjectPtr ptr, QtObjectPtr dateTime, int format);
char* QLocale_ToString7(QtObjectPtr ptr, QtObjectPtr dateTime, char* format);
char* QLocale_ToString5(QtObjectPtr ptr, QtObjectPtr time, int format);
char* QLocale_ToString4(QtObjectPtr ptr, QtObjectPtr time, char* format);
char* QLocale_ToString12(QtObjectPtr ptr, int i);
char* QLocale_ToUpper(QtObjectPtr ptr, char* str);
char* QLocale_UiLanguages(QtObjectPtr ptr);
void QLocale_DestroyQLocale(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif