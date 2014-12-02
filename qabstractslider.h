#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QAbstractSlider_New_QWidget(QtObjectPtr parent);
void QAbstractSlider_Destroy(QtObjectPtr ptr);
int QAbstractSlider_HasTracking(QtObjectPtr ptr);
int QAbstractSlider_InvertedAppearance(QtObjectPtr ptr);
int QAbstractSlider_InvertedControls(QtObjectPtr ptr);
int QAbstractSlider_IsSliderDown(QtObjectPtr ptr);
int QAbstractSlider_Maximum(QtObjectPtr ptr);
int QAbstractSlider_Minimum(QtObjectPtr ptr);
int QAbstractSlider_Orientation(QtObjectPtr ptr);
int QAbstractSlider_PageStep(QtObjectPtr ptr);
void QAbstractSlider_SetInvertedAppearance_Bool(QtObjectPtr ptr, int invertedAppearance);
void QAbstractSlider_SetInvertedControls_Bool(QtObjectPtr ptr, int invertedControls);
void QAbstractSlider_SetMaximum_Int(QtObjectPtr ptr, int maximum);
void QAbstractSlider_SetMinimum_Int(QtObjectPtr ptr, int minimum);
void QAbstractSlider_SetPageStep_Int(QtObjectPtr ptr, int pageStep);
void QAbstractSlider_SetSingleStep_Int(QtObjectPtr ptr, int singleStep);
void QAbstractSlider_SetSliderDown_Bool(QtObjectPtr ptr, int sliderDown);
void QAbstractSlider_SetSliderPosition_Int(QtObjectPtr ptr, int sliderPosition);
void QAbstractSlider_SetTracking_Bool(QtObjectPtr ptr, int enable);
int QAbstractSlider_SingleStep(QtObjectPtr ptr);
int QAbstractSlider_SliderPosition(QtObjectPtr ptr);
int QAbstractSlider_Value(QtObjectPtr ptr);
//Public Slots
void QAbstractSlider_ConnectSlotSetOrientation(QtObjectPtr ptr);
void QAbstractSlider_DisconnectSlotSetOrientation(QtObjectPtr ptr);
void QAbstractSlider_SetOrientation_Orientation(QtObjectPtr ptr, int orientation);
void QAbstractSlider_ConnectSlotSetRange(QtObjectPtr ptr);
void QAbstractSlider_DisconnectSlotSetRange(QtObjectPtr ptr);
void QAbstractSlider_SetRange_Int_Int(QtObjectPtr ptr, int min, int max);
void QAbstractSlider_ConnectSlotSetValue(QtObjectPtr ptr);
void QAbstractSlider_DisconnectSlotSetValue(QtObjectPtr ptr);
void QAbstractSlider_SetValue_Int(QtObjectPtr ptr, int value);
//Signals
void QAbstractSlider_ConnectSignalActionTriggered(QtObjectPtr ptr);
void QAbstractSlider_DisconnectSignalActionTriggered(QtObjectPtr ptr);
void QAbstractSlider_ConnectSignalRangeChanged(QtObjectPtr ptr);
void QAbstractSlider_DisconnectSignalRangeChanged(QtObjectPtr ptr);
void QAbstractSlider_ConnectSignalSliderMoved(QtObjectPtr ptr);
void QAbstractSlider_DisconnectSignalSliderMoved(QtObjectPtr ptr);
void QAbstractSlider_ConnectSignalSliderPressed(QtObjectPtr ptr);
void QAbstractSlider_DisconnectSignalSliderPressed(QtObjectPtr ptr);
void QAbstractSlider_ConnectSignalSliderReleased(QtObjectPtr ptr);
void QAbstractSlider_DisconnectSignalSliderReleased(QtObjectPtr ptr);
void QAbstractSlider_ConnectSignalValueChanged(QtObjectPtr ptr);
void QAbstractSlider_DisconnectSignalValueChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
