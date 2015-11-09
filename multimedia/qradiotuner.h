#ifdef __cplusplus
extern "C" {
#endif

int QRadioTuner_Band(void* ptr);
int QRadioTuner_Frequency(void* ptr);
int QRadioTuner_IsAntennaConnected(void* ptr);
int QRadioTuner_IsMuted(void* ptr);
int QRadioTuner_IsSearching(void* ptr);
int QRadioTuner_IsStereo(void* ptr);
void* QRadioTuner_RadioData(void* ptr);
void QRadioTuner_SetMuted(void* ptr, int muted);
void QRadioTuner_SetStereoMode(void* ptr, int mode);
void QRadioTuner_SetVolume(void* ptr, int volume);
int QRadioTuner_SignalStrength(void* ptr);
int QRadioTuner_StereoMode(void* ptr);
int QRadioTuner_Volume(void* ptr);
void* QRadioTuner_NewQRadioTuner(void* parent);
void QRadioTuner_ConnectAntennaConnectedChanged(void* ptr);
void QRadioTuner_DisconnectAntennaConnectedChanged(void* ptr);
void QRadioTuner_ConnectBandChanged(void* ptr);
void QRadioTuner_DisconnectBandChanged(void* ptr);
void QRadioTuner_CancelSearch(void* ptr);
int QRadioTuner_Error(void* ptr);
char* QRadioTuner_ErrorString(void* ptr);
void QRadioTuner_ConnectFrequencyChanged(void* ptr);
void QRadioTuner_DisconnectFrequencyChanged(void* ptr);
int QRadioTuner_FrequencyStep(void* ptr, int band);
int QRadioTuner_IsBandSupported(void* ptr, int band);
void QRadioTuner_ConnectMutedChanged(void* ptr);
void QRadioTuner_DisconnectMutedChanged(void* ptr);
void QRadioTuner_SearchAllStations(void* ptr, int searchMode);
void QRadioTuner_SearchBackward(void* ptr);
void QRadioTuner_SearchForward(void* ptr);
void QRadioTuner_ConnectSearchingChanged(void* ptr);
void QRadioTuner_DisconnectSearchingChanged(void* ptr);
void QRadioTuner_SetBand(void* ptr, int band);
void QRadioTuner_SetFrequency(void* ptr, int frequency);
void QRadioTuner_ConnectSignalStrengthChanged(void* ptr);
void QRadioTuner_DisconnectSignalStrengthChanged(void* ptr);
void QRadioTuner_Start(void* ptr);
void QRadioTuner_ConnectStateChanged(void* ptr);
void QRadioTuner_DisconnectStateChanged(void* ptr);
void QRadioTuner_ConnectStationFound(void* ptr);
void QRadioTuner_DisconnectStationFound(void* ptr);
void QRadioTuner_ConnectStereoStatusChanged(void* ptr);
void QRadioTuner_DisconnectStereoStatusChanged(void* ptr);
void QRadioTuner_Stop(void* ptr);
void QRadioTuner_ConnectVolumeChanged(void* ptr);
void QRadioTuner_DisconnectVolumeChanged(void* ptr);
void QRadioTuner_DestroyQRadioTuner(void* ptr);

#ifdef __cplusplus
}
#endif