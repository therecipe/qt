#ifdef __cplusplus
extern "C" {
#endif

void* QColor_ConvertTo(void* ptr, int colorSpec);
void QColor_SetRgbF(void* ptr, double r, double g, double b, double a);
void* QColor_NewQColor();
void* QColor_NewQColor8(int color);
void* QColor_NewQColor6(void* color);
void* QColor_NewQColor4(char* name);
void* QColor_NewQColor5(char* name);
void* QColor_NewQColor2(int r, int g, int b, int a);
int QColor_Alpha(void* ptr);
double QColor_AlphaF(void* ptr);
int QColor_Black(void* ptr);
double QColor_BlackF(void* ptr);
int QColor_Blue(void* ptr);
double QColor_BlueF(void* ptr);
char* QColor_QColor_ColorNames();
int QColor_Cyan(void* ptr);
double QColor_CyanF(void* ptr);
void* QColor_Darker(void* ptr, int factor);
void* QColor_QColor_FromCmyk(int c, int m, int y, int k, int a);
void* QColor_QColor_FromCmykF(double c, double m, double y, double k, double a);
void* QColor_QColor_FromHsl(int h, int s, int l, int a);
void* QColor_QColor_FromHslF(double h, double s, double l, double a);
void* QColor_QColor_FromHsv(int h, int s, int v, int a);
void* QColor_QColor_FromHsvF(double h, double s, double v, double a);
void* QColor_QColor_FromRgb2(int r, int g, int b, int a);
void* QColor_QColor_FromRgbF(double r, double g, double b, double a);
void QColor_GetCmyk(void* ptr, int c, int m, int y, int k, int a);
void QColor_GetHsl(void* ptr, int h, int s, int l, int a);
void QColor_GetHsv(void* ptr, int h, int s, int v, int a);
void QColor_GetRgb(void* ptr, int r, int g, int b, int a);
int QColor_Green(void* ptr);
double QColor_GreenF(void* ptr);
int QColor_HslHue(void* ptr);
double QColor_HslHueF(void* ptr);
int QColor_HslSaturation(void* ptr);
double QColor_HslSaturationF(void* ptr);
int QColor_HsvHue(void* ptr);
double QColor_HsvHueF(void* ptr);
int QColor_HsvSaturation(void* ptr);
double QColor_HsvSaturationF(void* ptr);
int QColor_Hue(void* ptr);
double QColor_HueF(void* ptr);
int QColor_IsValid(void* ptr);
int QColor_QColor_IsValidColor(char* name);
void* QColor_Lighter(void* ptr, int factor);
int QColor_Lightness(void* ptr);
double QColor_LightnessF(void* ptr);
int QColor_Magenta(void* ptr);
double QColor_MagentaF(void* ptr);
char* QColor_Name(void* ptr);
char* QColor_Name2(void* ptr, int format);
int QColor_Red(void* ptr);
double QColor_RedF(void* ptr);
int QColor_Saturation(void* ptr);
double QColor_SaturationF(void* ptr);
void QColor_SetAlpha(void* ptr, int alpha);
void QColor_SetAlphaF(void* ptr, double alpha);
void QColor_SetBlue(void* ptr, int blue);
void QColor_SetBlueF(void* ptr, double blue);
void QColor_SetCmyk(void* ptr, int c, int m, int y, int k, int a);
void QColor_SetCmykF(void* ptr, double c, double m, double y, double k, double a);
void QColor_SetGreen(void* ptr, int green);
void QColor_SetGreenF(void* ptr, double green);
void QColor_SetHsl(void* ptr, int h, int s, int l, int a);
void QColor_SetHslF(void* ptr, double h, double s, double l, double a);
void QColor_SetHsv(void* ptr, int h, int s, int v, int a);
void QColor_SetHsvF(void* ptr, double h, double s, double v, double a);
void QColor_SetNamedColor(void* ptr, char* name);
void QColor_SetRed(void* ptr, int red);
void QColor_SetRedF(void* ptr, double red);
void QColor_SetRgb(void* ptr, int r, int g, int b, int a);
int QColor_Spec(void* ptr);
void* QColor_ToCmyk(void* ptr);
void* QColor_ToHsl(void* ptr);
void* QColor_ToHsv(void* ptr);
void* QColor_ToRgb(void* ptr);
int QColor_Value(void* ptr);
double QColor_ValueF(void* ptr);
int QColor_Yellow(void* ptr);
double QColor_YellowF(void* ptr);

#ifdef __cplusplus
}
#endif