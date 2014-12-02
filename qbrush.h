#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QBrush_New();
QtObjectPtr QBrush_New_BrushStyle(int style);
QtObjectPtr QBrush_New_GlobalColor_BrushStyle(int color, int style);
void QBrush_Destroy(QtObjectPtr ptr);
int QBrush_IsOpaque(QtObjectPtr ptr);
void QBrush_SetColor_GlobalColor(QtObjectPtr ptr, int color);
void QBrush_SetStyle_BrushStyle(QtObjectPtr ptr, int style);
int QBrush_Style(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
