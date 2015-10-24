#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QDial_NotchSize(QtObjectPtr ptr);
int QDial_NotchesVisible(QtObjectPtr ptr);
void QDial_SetNotchesVisible(QtObjectPtr ptr, int visible);
void QDial_SetWrapping(QtObjectPtr ptr, int on);
int QDial_Wrapping(QtObjectPtr ptr);
QtObjectPtr QDial_NewQDial(QtObjectPtr parent);
void QDial_DestroyQDial(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif