#ifdef __cplusplus
extern "C" {
#endif

void* QLocale_NewQLocale();
void* QLocale_NewQLocale3(int language, int country);
void* QLocale_NewQLocale4(int language, int script, int country);
void* QLocale_NewQLocale5(void* other);
void* QLocale_NewQLocale2(char* name);
char* QLocale_AmText(void* ptr);
char* QLocale_Bcp47Name(void* ptr);
int QLocale_Country(void* ptr);
char* QLocale_QLocale_CountryToString(int country);
char* QLocale_CreateSeparatedList(void* ptr, char* list);
char* QLocale_CurrencySymbol(void* ptr, int format);
char* QLocale_DateFormat(void* ptr, int format);
char* QLocale_DateTimeFormat(void* ptr, int format);
char* QLocale_DayName(void* ptr, int day, int ty);
int QLocale_FirstDayOfWeek(void* ptr);
int QLocale_Language(void* ptr);
char* QLocale_QLocale_LanguageToString(int language);
int QLocale_MeasurementSystem(void* ptr);
char* QLocale_MonthName(void* ptr, int month, int ty);
char* QLocale_Name(void* ptr);
char* QLocale_NativeCountryName(void* ptr);
char* QLocale_NativeLanguageName(void* ptr);
int QLocale_NumberOptions(void* ptr);
char* QLocale_PmText(void* ptr);
char* QLocale_QuoteString(void* ptr, char* str, int style);
char* QLocale_QuoteString2(void* ptr, void* str, int style);
int QLocale_Script(void* ptr);
char* QLocale_QLocale_ScriptToString(int script);
void QLocale_QLocale_SetDefault(void* locale);
void QLocale_SetNumberOptions(void* ptr, int options);
char* QLocale_StandaloneDayName(void* ptr, int day, int ty);
char* QLocale_StandaloneMonthName(void* ptr, int month, int ty);
int QLocale_TextDirection(void* ptr);
char* QLocale_TimeFormat(void* ptr, int format);
char* QLocale_ToCurrencyString6(void* ptr, int value, char* symbol);
void* QLocale_ToDateTime(void* ptr, char* stri, int format);
void* QLocale_ToDateTime2(void* ptr, char* stri, char* format);
int QLocale_ToInt(void* ptr, char* s, int ok);
int QLocale_ToInt2(void* ptr, void* s, int ok);
char* QLocale_ToLower(void* ptr, char* str);
char* QLocale_ToString3(void* ptr, void* date, int format);
char* QLocale_ToString2(void* ptr, void* date, char* format);
char* QLocale_ToString6(void* ptr, void* dateTime, int format);
char* QLocale_ToString7(void* ptr, void* dateTime, char* format);
char* QLocale_ToString5(void* ptr, void* time, int format);
char* QLocale_ToString4(void* ptr, void* time, char* format);
char* QLocale_ToString12(void* ptr, int i);
char* QLocale_ToUpper(void* ptr, char* str);
char* QLocale_UiLanguages(void* ptr);
void QLocale_DestroyQLocale(void* ptr);

#ifdef __cplusplus
}
#endif