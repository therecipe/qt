#ifdef __cplusplus
extern "C" {
#endif

void QRadioDataControl_ConnectAlternativeFrequenciesEnabledChanged(void* ptr);
void QRadioDataControl_DisconnectAlternativeFrequenciesEnabledChanged(void* ptr);
int QRadioDataControl_Error(void* ptr);
char* QRadioDataControl_ErrorString(void* ptr);
int QRadioDataControl_IsAlternativeFrequenciesEnabled(void* ptr);
int QRadioDataControl_ProgramType(void* ptr);
void QRadioDataControl_ConnectProgramTypeChanged(void* ptr);
void QRadioDataControl_DisconnectProgramTypeChanged(void* ptr);
char* QRadioDataControl_ProgramTypeName(void* ptr);
void QRadioDataControl_ConnectProgramTypeNameChanged(void* ptr);
void QRadioDataControl_DisconnectProgramTypeNameChanged(void* ptr);
char* QRadioDataControl_RadioText(void* ptr);
void QRadioDataControl_ConnectRadioTextChanged(void* ptr);
void QRadioDataControl_DisconnectRadioTextChanged(void* ptr);
void QRadioDataControl_SetAlternativeFrequenciesEnabled(void* ptr, int enabled);
char* QRadioDataControl_StationId(void* ptr);
void QRadioDataControl_ConnectStationIdChanged(void* ptr);
void QRadioDataControl_DisconnectStationIdChanged(void* ptr);
char* QRadioDataControl_StationName(void* ptr);
void QRadioDataControl_ConnectStationNameChanged(void* ptr);
void QRadioDataControl_DisconnectStationNameChanged(void* ptr);
void QRadioDataControl_DestroyQRadioDataControl(void* ptr);

#ifdef __cplusplus
}
#endif