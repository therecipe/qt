#ifdef __cplusplus
extern "C" {
#endif

int QRadioData_IsAlternativeFrequenciesEnabled(void* ptr);
int QRadioData_ProgramType(void* ptr);
char* QRadioData_ProgramTypeName(void* ptr);
char* QRadioData_RadioText(void* ptr);
void QRadioData_SetAlternativeFrequenciesEnabled(void* ptr, int enabled);
char* QRadioData_StationId(void* ptr);
char* QRadioData_StationName(void* ptr);
void* QRadioData_NewQRadioData(void* mediaObject, void* parent);
void QRadioData_ConnectAlternativeFrequenciesEnabledChanged(void* ptr);
void QRadioData_DisconnectAlternativeFrequenciesEnabledChanged(void* ptr);
int QRadioData_Error(void* ptr);
char* QRadioData_ErrorString(void* ptr);
void* QRadioData_MediaObject(void* ptr);
void QRadioData_ConnectProgramTypeChanged(void* ptr);
void QRadioData_DisconnectProgramTypeChanged(void* ptr);
void QRadioData_ConnectProgramTypeNameChanged(void* ptr);
void QRadioData_DisconnectProgramTypeNameChanged(void* ptr);
void QRadioData_ConnectRadioTextChanged(void* ptr);
void QRadioData_DisconnectRadioTextChanged(void* ptr);
void QRadioData_ConnectStationIdChanged(void* ptr);
void QRadioData_DisconnectStationIdChanged(void* ptr);
void QRadioData_ConnectStationNameChanged(void* ptr);
void QRadioData_DisconnectStationNameChanged(void* ptr);
void QRadioData_DestroyQRadioData(void* ptr);

#ifdef __cplusplus
}
#endif