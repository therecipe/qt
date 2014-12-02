#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QStackedLayout_New();
QtObjectPtr QStackedLayout_New_QWidget(QtObjectPtr parent);
QtObjectPtr QStackedLayout_New_QLayout(QtObjectPtr parentLayout);
void QStackedLayout_Destroy(QtObjectPtr ptr);
int QStackedLayout_CurrentIndex(QtObjectPtr ptr);
QtObjectPtr QStackedLayout_CurrentWidget(QtObjectPtr ptr);
int QStackedLayout_InsertWidget_Int_QWidget(QtObjectPtr ptr, int index, QtObjectPtr widget);
QtObjectPtr QStackedLayout_Widget_Int(QtObjectPtr ptr, int index);
//Public Slots
void QStackedLayout_ConnectSlotSetCurrentIndex(QtObjectPtr ptr);
void QStackedLayout_DisconnectSlotSetCurrentIndex(QtObjectPtr ptr);
void QStackedLayout_SetCurrentIndex_Int(QtObjectPtr ptr, int index);
//Signals
void QStackedLayout_ConnectSignalCurrentChanged(QtObjectPtr ptr);
void QStackedLayout_DisconnectSignalCurrentChanged(QtObjectPtr ptr);
void QStackedLayout_ConnectSignalWidgetRemoved(QtObjectPtr ptr);
void QStackedLayout_DisconnectSignalWidgetRemoved(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
