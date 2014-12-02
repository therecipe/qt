#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QSlider_New_QWidget(QtObjectPtr parent);
QtObjectPtr QSlider_New_Orientation_QWidget(int orientation, QtObjectPtr parent);
void QSlider_Destroy(QtObjectPtr ptr);
void QSlider_SetTickInterval_Int(QtObjectPtr ptr, int ti);
int QSlider_TickInterval(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
