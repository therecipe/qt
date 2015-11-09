#ifdef __cplusplus
extern "C" {
#endif

void* QRawFont_NewQRawFont();
void* QRawFont_NewQRawFont3(void* fontData, double pixelSize, int hintingPreference);
void* QRawFont_NewQRawFont4(void* other);
void* QRawFont_NewQRawFont2(char* fileName, double pixelSize, int hintingPreference);
double QRawFont_Ascent(void* ptr);
double QRawFont_AverageCharWidth(void* ptr);
double QRawFont_Descent(void* ptr);
char* QRawFont_FamilyName(void* ptr);
void* QRawFont_FontTable(void* ptr, char* tagName);
int QRawFont_HintingPreference(void* ptr);
int QRawFont_IsValid(void* ptr);
double QRawFont_Leading(void* ptr);
double QRawFont_LineThickness(void* ptr);
void QRawFont_LoadFromData(void* ptr, void* fontData, double pixelSize, int hintingPreference);
void QRawFont_LoadFromFile(void* ptr, char* fileName, double pixelSize, int hintingPreference);
double QRawFont_MaxCharWidth(void* ptr);
double QRawFont_PixelSize(void* ptr);
void QRawFont_SetPixelSize(void* ptr, double pixelSize);
int QRawFont_Style(void* ptr);
char* QRawFont_StyleName(void* ptr);
int QRawFont_SupportsCharacter(void* ptr, void* character);
void QRawFont_Swap(void* ptr, void* other);
double QRawFont_UnderlinePosition(void* ptr);
double QRawFont_UnitsPerEm(void* ptr);
int QRawFont_Weight(void* ptr);
double QRawFont_XHeight(void* ptr);
void QRawFont_DestroyQRawFont(void* ptr);

#ifdef __cplusplus
}
#endif