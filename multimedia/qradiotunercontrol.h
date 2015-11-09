#ifdef __cplusplus
extern "C" {
#endif

void QRadioTunerControl_ConnectAntennaConnectedChanged(void* ptr);
void QRadioTunerControl_DisconnectAntennaConnectedChanged(void* ptr);
int QRadioTunerControl_Band(void* ptr);
void QRadioTunerControl_ConnectBandChanged(void* ptr);
void QRadioTunerControl_DisconnectBandChanged(void* ptr);
void QRadioTunerControl_CancelSearch(void* ptr);
int QRadioTunerControl_Error(void* ptr);
char* QRadioTunerControl_ErrorString(void* ptr);
int QRadioTunerControl_Frequency(void* ptr);
void QRadioTunerControl_ConnectFrequencyChanged(void* ptr);
void QRadioTunerControl_DisconnectFrequencyChanged(void* ptr);
int QRadioTunerControl_FrequencyStep(void* ptr, int band);
int QRadioTunerControl_IsAntennaConnected(void* ptr);
int QRadioTunerControl_IsBandSupported(void* ptr, int band);
int QRadioTunerControl_IsMuted(void* ptr);
int QRadioTunerControl_IsSearching(void* ptr);
int QRadioTunerControl_IsStereo(void* ptr);
void QRadioTunerControl_ConnectMutedChanged(void* ptr);
void QRadioTunerControl_DisconnectMutedChanged(void* ptr);
void QRadioTunerControl_SearchAllStations(void* ptr, int searchMode);
void QRadioTunerControl_SearchBackward(void* ptr);
void QRadioTunerControl_SearchForward(void* ptr);
void QRadioTunerControl_ConnectSearchingChanged(void* ptr);
void QRadioTunerControl_DisconnectSearchingChanged(void* ptr);
void QRadioTunerControl_SetBand(void* ptr, int band);
void QRadioTunerControl_SetFrequency(void* ptr, int frequency);
void QRadioTunerControl_SetMuted(void* ptr, int muted);
void QRadioTunerControl_SetStereoMode(void* ptr, int mode);
void QRadioTunerControl_SetVolume(void* ptr, int volume);
int QRadioTunerControl_SignalStrength(void* ptr);
void QRadioTunerControl_ConnectSignalStrengthChanged(void* ptr);
void QRadioTunerControl_DisconnectSignalStrengthChanged(void* ptr);
void QRadioTunerControl_Start(void* ptr);
void QRadioTunerControl_ConnectStateChanged(void* ptr);
void QRadioTunerControl_DisconnectStateChanged(void* ptr);
void QRadioTunerControl_ConnectStationFound(void* ptr);
void QRadioTunerControl_DisconnectStationFound(void* ptr);
int QRadioTunerControl_StereoMode(void* ptr);
void QRadioTunerControl_ConnectStereoStatusChanged(void* ptr);
void QRadioTunerControl_DisconnectStereoStatusChanged(void* ptr);
void QRadioTunerControl_Stop(void* ptr);
int QRadioTunerControl_Volume(void* ptr);
void QRadioTunerControl_ConnectVolumeChanged(void* ptr);
void QRadioTunerControl_DisconnectVolumeChanged(void* ptr);
void QRadioTunerControl_DestroyQRadioTunerControl(void* ptr);

#ifdef __cplusplus
}
#endif