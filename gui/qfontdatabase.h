#ifdef __cplusplus
extern "C" {
#endif

int QFontDatabase_Ogham_Type();
int QFontDatabase_Runic_Type();
int QFontDatabase_Nko_Type();
int QFontDatabase_WritingSystemsCount_Type();
int QFontDatabase_QFontDatabase_RemoveAllApplicationFonts();
int QFontDatabase_QFontDatabase_RemoveApplicationFont(int id);
void* QFontDatabase_NewQFontDatabase();
int QFontDatabase_QFontDatabase_AddApplicationFont(char* fileName);
int QFontDatabase_QFontDatabase_AddApplicationFontFromData(void* fontData);
char* QFontDatabase_QFontDatabase_ApplicationFontFamilies(int id);
int QFontDatabase_Bold(void* ptr, char* family, char* style);
char* QFontDatabase_Families(void* ptr, int writingSystem);
int QFontDatabase_IsBitmapScalable(void* ptr, char* family, char* style);
int QFontDatabase_IsFixedPitch(void* ptr, char* family, char* style);
int QFontDatabase_IsPrivateFamily(void* ptr, char* family);
int QFontDatabase_IsScalable(void* ptr, char* family, char* style);
int QFontDatabase_IsSmoothlyScalable(void* ptr, char* family, char* style);
int QFontDatabase_Italic(void* ptr, char* family, char* style);
char* QFontDatabase_StyleString(void* ptr, void* font);
char* QFontDatabase_StyleString2(void* ptr, void* fontInfo);
char* QFontDatabase_Styles(void* ptr, char* family);
int QFontDatabase_Weight(void* ptr, char* family, char* style);
char* QFontDatabase_QFontDatabase_WritingSystemName(int writingSystem);
char* QFontDatabase_QFontDatabase_WritingSystemSample(int writingSystem);

#ifdef __cplusplus
}
#endif