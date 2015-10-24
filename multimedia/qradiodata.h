#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QRadioData_IsAlternativeFrequenciesEnabled(QtObjectPtr ptr);
int QRadioData_ProgramType(QtObjectPtr ptr);
char* QRadioData_ProgramTypeName(QtObjectPtr ptr);
char* QRadioData_RadioText(QtObjectPtr ptr);
void QRadioData_SetAlternativeFrequenciesEnabled(QtObjectPtr ptr, int enabled);
char* QRadioData_StationId(QtObjectPtr ptr);
char* QRadioData_StationName(QtObjectPtr ptr);
QtObjectPtr QRadioData_NewQRadioData(QtObjectPtr mediaObject, QtObjectPtr parent);
void QRadioData_ConnectAlternativeFrequenciesEnabledChanged(QtObjectPtr ptr);
void QRadioData_DisconnectAlternativeFrequenciesEnabledChanged(QtObjectPtr ptr);
int QRadioData_Error(QtObjectPtr ptr);
char* QRadioData_ErrorString(QtObjectPtr ptr);
QtObjectPtr QRadioData_MediaObject(QtObjectPtr ptr);
void QRadioData_ConnectProgramTypeChanged(QtObjectPtr ptr);
void QRadioData_DisconnectProgramTypeChanged(QtObjectPtr ptr);
void QRadioData_ConnectProgramTypeNameChanged(QtObjectPtr ptr);
void QRadioData_DisconnectProgramTypeNameChanged(QtObjectPtr ptr);
void QRadioData_ConnectRadioTextChanged(QtObjectPtr ptr);
void QRadioData_DisconnectRadioTextChanged(QtObjectPtr ptr);
void QRadioData_ConnectStationIdChanged(QtObjectPtr ptr);
void QRadioData_DisconnectStationIdChanged(QtObjectPtr ptr);
void QRadioData_ConnectStationNameChanged(QtObjectPtr ptr);
void QRadioData_DisconnectStationNameChanged(QtObjectPtr ptr);
void QRadioData_DestroyQRadioData(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif