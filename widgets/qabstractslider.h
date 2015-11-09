#ifdef __cplusplus
extern "C" {
#endif

int QAbstractSlider_HasTracking(void* ptr);
int QAbstractSlider_InvertedAppearance(void* ptr);
int QAbstractSlider_InvertedControls(void* ptr);
int QAbstractSlider_IsSliderDown(void* ptr);
int QAbstractSlider_Maximum(void* ptr);
int QAbstractSlider_Minimum(void* ptr);
int QAbstractSlider_Orientation(void* ptr);
int QAbstractSlider_PageStep(void* ptr);
void QAbstractSlider_SetInvertedAppearance(void* ptr, int v);
void QAbstractSlider_SetInvertedControls(void* ptr, int v);
void QAbstractSlider_SetMaximum(void* ptr, int v);
void QAbstractSlider_SetMinimum(void* ptr, int v);
void QAbstractSlider_SetOrientation(void* ptr, int v);
void QAbstractSlider_SetPageStep(void* ptr, int v);
void QAbstractSlider_SetSingleStep(void* ptr, int v);
void QAbstractSlider_SetSliderDown(void* ptr, int v);
void QAbstractSlider_SetSliderPosition(void* ptr, int v);
void QAbstractSlider_SetTracking(void* ptr, int enable);
void QAbstractSlider_SetValue(void* ptr, int v);
int QAbstractSlider_SingleStep(void* ptr);
int QAbstractSlider_SliderPosition(void* ptr);
int QAbstractSlider_Value(void* ptr);
void* QAbstractSlider_NewQAbstractSlider(void* parent);
void QAbstractSlider_ConnectActionTriggered(void* ptr);
void QAbstractSlider_DisconnectActionTriggered(void* ptr);
void QAbstractSlider_ConnectRangeChanged(void* ptr);
void QAbstractSlider_DisconnectRangeChanged(void* ptr);
void QAbstractSlider_SetRange(void* ptr, int min, int max);
void QAbstractSlider_ConnectSliderMoved(void* ptr);
void QAbstractSlider_DisconnectSliderMoved(void* ptr);
void QAbstractSlider_ConnectSliderPressed(void* ptr);
void QAbstractSlider_DisconnectSliderPressed(void* ptr);
void QAbstractSlider_ConnectSliderReleased(void* ptr);
void QAbstractSlider_DisconnectSliderReleased(void* ptr);
void QAbstractSlider_TriggerAction(void* ptr, int action);
void QAbstractSlider_ConnectValueChanged(void* ptr);
void QAbstractSlider_DisconnectValueChanged(void* ptr);
void QAbstractSlider_DestroyQAbstractSlider(void* ptr);

#ifdef __cplusplus
}
#endif