#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QRadioDataControl_ConnectAlternativeFrequenciesEnabledChanged(QtObjectPtr ptr);
void QRadioDataControl_DisconnectAlternativeFrequenciesEnabledChanged(QtObjectPtr ptr);
int QRadioDataControl_Error(QtObjectPtr ptr);
char* QRadioDataControl_ErrorString(QtObjectPtr ptr);
int QRadioDataControl_IsAlternativeFrequenciesEnabled(QtObjectPtr ptr);
int QRadioDataControl_ProgramType(QtObjectPtr ptr);
void QRadioDataControl_ConnectProgramTypeChanged(QtObjectPtr ptr);
void QRadioDataControl_DisconnectProgramTypeChanged(QtObjectPtr ptr);
char* QRadioDataControl_ProgramTypeName(QtObjectPtr ptr);
void QRadioDataControl_ConnectProgramTypeNameChanged(QtObjectPtr ptr);
void QRadioDataControl_DisconnectProgramTypeNameChanged(QtObjectPtr ptr);
char* QRadioDataControl_RadioText(QtObjectPtr ptr);
void QRadioDataControl_ConnectRadioTextChanged(QtObjectPtr ptr);
void QRadioDataControl_DisconnectRadioTextChanged(QtObjectPtr ptr);
void QRadioDataControl_SetAlternativeFrequenciesEnabled(QtObjectPtr ptr, int enabled);
char* QRadioDataControl_StationId(QtObjectPtr ptr);
void QRadioDataControl_ConnectStationIdChanged(QtObjectPtr ptr);
void QRadioDataControl_DisconnectStationIdChanged(QtObjectPtr ptr);
char* QRadioDataControl_StationName(QtObjectPtr ptr);
void QRadioDataControl_ConnectStationNameChanged(QtObjectPtr ptr);
void QRadioDataControl_DisconnectStationNameChanged(QtObjectPtr ptr);
void QRadioDataControl_DestroyQRadioDataControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif