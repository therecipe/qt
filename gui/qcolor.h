#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QColor_NewQColor();
QtObjectPtr QColor_NewQColor8(int color);
QtObjectPtr QColor_NewQColor6(QtObjectPtr color);
QtObjectPtr QColor_NewQColor4(char* name);
QtObjectPtr QColor_NewQColor5(char* name);
QtObjectPtr QColor_NewQColor2(int r, int g, int b, int a);
int QColor_Alpha(QtObjectPtr ptr);
int QColor_Black(QtObjectPtr ptr);
int QColor_Blue(QtObjectPtr ptr);
char* QColor_QColor_ColorNames();
int QColor_Cyan(QtObjectPtr ptr);
void QColor_GetCmyk(QtObjectPtr ptr, int c, int m, int y, int k, int a);
void QColor_GetHsl(QtObjectPtr ptr, int h, int s, int l, int a);
void QColor_GetHsv(QtObjectPtr ptr, int h, int s, int v, int a);
void QColor_GetRgb(QtObjectPtr ptr, int r, int g, int b, int a);
int QColor_Green(QtObjectPtr ptr);
int QColor_HslHue(QtObjectPtr ptr);
int QColor_HslSaturation(QtObjectPtr ptr);
int QColor_HsvHue(QtObjectPtr ptr);
int QColor_HsvSaturation(QtObjectPtr ptr);
int QColor_Hue(QtObjectPtr ptr);
int QColor_IsValid(QtObjectPtr ptr);
int QColor_QColor_IsValidColor(char* name);
int QColor_Lightness(QtObjectPtr ptr);
int QColor_Magenta(QtObjectPtr ptr);
char* QColor_Name(QtObjectPtr ptr);
char* QColor_Name2(QtObjectPtr ptr, int format);
int QColor_Red(QtObjectPtr ptr);
int QColor_Saturation(QtObjectPtr ptr);
void QColor_SetAlpha(QtObjectPtr ptr, int alpha);
void QColor_SetBlue(QtObjectPtr ptr, int blue);
void QColor_SetCmyk(QtObjectPtr ptr, int c, int m, int y, int k, int a);
void QColor_SetGreen(QtObjectPtr ptr, int green);
void QColor_SetHsl(QtObjectPtr ptr, int h, int s, int l, int a);
void QColor_SetHsv(QtObjectPtr ptr, int h, int s, int v, int a);
void QColor_SetNamedColor(QtObjectPtr ptr, char* name);
void QColor_SetRed(QtObjectPtr ptr, int red);
void QColor_SetRgb(QtObjectPtr ptr, int r, int g, int b, int a);
int QColor_Spec(QtObjectPtr ptr);
int QColor_Value(QtObjectPtr ptr);
int QColor_Yellow(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif