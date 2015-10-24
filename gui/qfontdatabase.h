#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QFontDatabase_Ogham_Type();
int QFontDatabase_Runic_Type();
int QFontDatabase_Nko_Type();
int QFontDatabase_WritingSystemsCount_Type();
int QFontDatabase_QFontDatabase_RemoveAllApplicationFonts();
int QFontDatabase_QFontDatabase_RemoveApplicationFont(int id);
QtObjectPtr QFontDatabase_NewQFontDatabase();
int QFontDatabase_QFontDatabase_AddApplicationFont(char* fileName);
int QFontDatabase_QFontDatabase_AddApplicationFontFromData(QtObjectPtr fontData);
char* QFontDatabase_QFontDatabase_ApplicationFontFamilies(int id);
int QFontDatabase_Bold(QtObjectPtr ptr, char* family, char* style);
char* QFontDatabase_Families(QtObjectPtr ptr, int writingSystem);
int QFontDatabase_IsBitmapScalable(QtObjectPtr ptr, char* family, char* style);
int QFontDatabase_IsFixedPitch(QtObjectPtr ptr, char* family, char* style);
int QFontDatabase_IsPrivateFamily(QtObjectPtr ptr, char* family);
int QFontDatabase_IsScalable(QtObjectPtr ptr, char* family, char* style);
int QFontDatabase_IsSmoothlyScalable(QtObjectPtr ptr, char* family, char* style);
int QFontDatabase_Italic(QtObjectPtr ptr, char* family, char* style);
char* QFontDatabase_StyleString(QtObjectPtr ptr, QtObjectPtr font);
char* QFontDatabase_StyleString2(QtObjectPtr ptr, QtObjectPtr fontInfo);
char* QFontDatabase_Styles(QtObjectPtr ptr, char* family);
int QFontDatabase_Weight(QtObjectPtr ptr, char* family, char* style);
char* QFontDatabase_QFontDatabase_WritingSystemName(int writingSystem);
char* QFontDatabase_QFontDatabase_WritingSystemSample(int writingSystem);

#ifdef __cplusplus
}
#endif