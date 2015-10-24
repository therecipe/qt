#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QStaticText_NewQStaticText();
QtObjectPtr QStaticText_NewQStaticText3(QtObjectPtr other);
QtObjectPtr QStaticText_NewQStaticText2(char* text);
int QStaticText_PerformanceHint(QtObjectPtr ptr);
void QStaticText_Prepare(QtObjectPtr ptr, QtObjectPtr matrix, QtObjectPtr font);
void QStaticText_SetPerformanceHint(QtObjectPtr ptr, int performanceHint);
void QStaticText_SetText(QtObjectPtr ptr, char* text);
void QStaticText_SetTextFormat(QtObjectPtr ptr, int textFormat);
void QStaticText_SetTextOption(QtObjectPtr ptr, QtObjectPtr textOption);
void QStaticText_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QStaticText_Text(QtObjectPtr ptr);
int QStaticText_TextFormat(QtObjectPtr ptr);
void QStaticText_DestroyQStaticText(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif