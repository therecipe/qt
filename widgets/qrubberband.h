#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QRubberBand_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect);
QtObjectPtr QRubberBand_NewQRubberBand(int s, QtObjectPtr p);
void QRubberBand_Move2(QtObjectPtr ptr, QtObjectPtr p);
void QRubberBand_Move(QtObjectPtr ptr, int x, int y);
void QRubberBand_Resize2(QtObjectPtr ptr, QtObjectPtr size);
void QRubberBand_Resize(QtObjectPtr ptr, int width, int height);
void QRubberBand_SetGeometry2(QtObjectPtr ptr, int x, int y, int width, int height);
int QRubberBand_Shape(QtObjectPtr ptr);
void QRubberBand_DestroyQRubberBand(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif