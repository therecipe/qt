#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QSlider_SetTickInterval(QtObjectPtr ptr, int ti);
void QSlider_SetTickPosition(QtObjectPtr ptr, int position);
int QSlider_TickInterval(QtObjectPtr ptr);
int QSlider_TickPosition(QtObjectPtr ptr);
QtObjectPtr QSlider_NewQSlider(QtObjectPtr parent);
QtObjectPtr QSlider_NewQSlider2(int orientation, QtObjectPtr parent);
int QSlider_Event(QtObjectPtr ptr, QtObjectPtr event);
void QSlider_DestroyQSlider(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif