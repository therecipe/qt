#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QAbstractSlider_HasTracking(QtObjectPtr ptr);
int QAbstractSlider_InvertedAppearance(QtObjectPtr ptr);
int QAbstractSlider_InvertedControls(QtObjectPtr ptr);
int QAbstractSlider_IsSliderDown(QtObjectPtr ptr);
int QAbstractSlider_Maximum(QtObjectPtr ptr);
int QAbstractSlider_Minimum(QtObjectPtr ptr);
int QAbstractSlider_Orientation(QtObjectPtr ptr);
int QAbstractSlider_PageStep(QtObjectPtr ptr);
void QAbstractSlider_SetInvertedAppearance(QtObjectPtr ptr, int v);
void QAbstractSlider_SetInvertedControls(QtObjectPtr ptr, int v);
void QAbstractSlider_SetMaximum(QtObjectPtr ptr, int v);
void QAbstractSlider_SetMinimum(QtObjectPtr ptr, int v);
void QAbstractSlider_SetOrientation(QtObjectPtr ptr, int v);
void QAbstractSlider_SetPageStep(QtObjectPtr ptr, int v);
void QAbstractSlider_SetSingleStep(QtObjectPtr ptr, int v);
void QAbstractSlider_SetSliderDown(QtObjectPtr ptr, int v);
void QAbstractSlider_SetSliderPosition(QtObjectPtr ptr, int v);
void QAbstractSlider_SetTracking(QtObjectPtr ptr, int enable);
void QAbstractSlider_SetValue(QtObjectPtr ptr, int v);
int QAbstractSlider_SingleStep(QtObjectPtr ptr);
int QAbstractSlider_SliderPosition(QtObjectPtr ptr);
int QAbstractSlider_Value(QtObjectPtr ptr);
QtObjectPtr QAbstractSlider_NewQAbstractSlider(QtObjectPtr parent);
void QAbstractSlider_ConnectActionTriggered(QtObjectPtr ptr);
void QAbstractSlider_DisconnectActionTriggered(QtObjectPtr ptr);
void QAbstractSlider_ConnectRangeChanged(QtObjectPtr ptr);
void QAbstractSlider_DisconnectRangeChanged(QtObjectPtr ptr);
void QAbstractSlider_SetRange(QtObjectPtr ptr, int min, int max);
void QAbstractSlider_ConnectSliderMoved(QtObjectPtr ptr);
void QAbstractSlider_DisconnectSliderMoved(QtObjectPtr ptr);
void QAbstractSlider_ConnectSliderPressed(QtObjectPtr ptr);
void QAbstractSlider_DisconnectSliderPressed(QtObjectPtr ptr);
void QAbstractSlider_ConnectSliderReleased(QtObjectPtr ptr);
void QAbstractSlider_DisconnectSliderReleased(QtObjectPtr ptr);
void QAbstractSlider_TriggerAction(QtObjectPtr ptr, int action);
void QAbstractSlider_ConnectValueChanged(QtObjectPtr ptr);
void QAbstractSlider_DisconnectValueChanged(QtObjectPtr ptr);
void QAbstractSlider_DestroyQAbstractSlider(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif