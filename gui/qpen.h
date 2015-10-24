#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QPen_NewQPen5(QtObjectPtr pen);
void QPen_SetCapStyle(QtObjectPtr ptr, int style);
void QPen_SetColor(QtObjectPtr ptr, QtObjectPtr color);
void QPen_SetJoinStyle(QtObjectPtr ptr, int style);
void QPen_SetStyle(QtObjectPtr ptr, int style);
void QPen_SetWidth(QtObjectPtr ptr, int width);
int QPen_Style(QtObjectPtr ptr);
int QPen_Width(QtObjectPtr ptr);
QtObjectPtr QPen_NewQPen();
QtObjectPtr QPen_NewQPen6(QtObjectPtr pen);
QtObjectPtr QPen_NewQPen2(int style);
QtObjectPtr QPen_NewQPen3(QtObjectPtr color);
int QPen_CapStyle(QtObjectPtr ptr);
int QPen_IsCosmetic(QtObjectPtr ptr);
int QPen_IsSolid(QtObjectPtr ptr);
int QPen_JoinStyle(QtObjectPtr ptr);
void QPen_SetBrush(QtObjectPtr ptr, QtObjectPtr brush);
void QPen_SetCosmetic(QtObjectPtr ptr, int cosmetic);
void QPen_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QPen_DestroyQPen(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif