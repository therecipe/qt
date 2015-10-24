#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QRawFont_NewQRawFont();
QtObjectPtr QRawFont_NewQRawFont4(QtObjectPtr other);
char* QRawFont_FamilyName(QtObjectPtr ptr);
int QRawFont_HintingPreference(QtObjectPtr ptr);
int QRawFont_IsValid(QtObjectPtr ptr);
int QRawFont_Style(QtObjectPtr ptr);
char* QRawFont_StyleName(QtObjectPtr ptr);
int QRawFont_SupportsCharacter(QtObjectPtr ptr, QtObjectPtr character);
void QRawFont_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QRawFont_Weight(QtObjectPtr ptr);
void QRawFont_DestroyQRawFont(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif