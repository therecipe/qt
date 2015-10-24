#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QBrush_NewQBrush4(int color, int style);
void QBrush_SetColor(QtObjectPtr ptr, QtObjectPtr color);
QtObjectPtr QBrush_NewQBrush();
QtObjectPtr QBrush_NewQBrush2(int style);
QtObjectPtr QBrush_NewQBrush6(int color, QtObjectPtr pixmap);
QtObjectPtr QBrush_NewQBrush9(QtObjectPtr other);
QtObjectPtr QBrush_NewQBrush3(QtObjectPtr color, int style);
QtObjectPtr QBrush_NewQBrush5(QtObjectPtr color, QtObjectPtr pixmap);
QtObjectPtr QBrush_NewQBrush10(QtObjectPtr gradient);
QtObjectPtr QBrush_NewQBrush8(QtObjectPtr image);
QtObjectPtr QBrush_NewQBrush7(QtObjectPtr pixmap);
QtObjectPtr QBrush_Gradient(QtObjectPtr ptr);
int QBrush_IsOpaque(QtObjectPtr ptr);
void QBrush_SetColor2(QtObjectPtr ptr, int color);
void QBrush_SetStyle(QtObjectPtr ptr, int style);
void QBrush_SetTexture(QtObjectPtr ptr, QtObjectPtr pixmap);
void QBrush_SetTextureImage(QtObjectPtr ptr, QtObjectPtr image);
void QBrush_SetTransform(QtObjectPtr ptr, QtObjectPtr matrix);
int QBrush_Style(QtObjectPtr ptr);
void QBrush_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QBrush_DestroyQBrush(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif