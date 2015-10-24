#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QRadioTuner_Band(QtObjectPtr ptr);
int QRadioTuner_Frequency(QtObjectPtr ptr);
int QRadioTuner_IsAntennaConnected(QtObjectPtr ptr);
int QRadioTuner_IsMuted(QtObjectPtr ptr);
int QRadioTuner_IsSearching(QtObjectPtr ptr);
int QRadioTuner_IsStereo(QtObjectPtr ptr);
QtObjectPtr QRadioTuner_RadioData(QtObjectPtr ptr);
void QRadioTuner_SetMuted(QtObjectPtr ptr, int muted);
void QRadioTuner_SetStereoMode(QtObjectPtr ptr, int mode);
void QRadioTuner_SetVolume(QtObjectPtr ptr, int volume);
int QRadioTuner_SignalStrength(QtObjectPtr ptr);
int QRadioTuner_StereoMode(QtObjectPtr ptr);
int QRadioTuner_Volume(QtObjectPtr ptr);
QtObjectPtr QRadioTuner_NewQRadioTuner(QtObjectPtr parent);
void QRadioTuner_ConnectAntennaConnectedChanged(QtObjectPtr ptr);
void QRadioTuner_DisconnectAntennaConnectedChanged(QtObjectPtr ptr);
void QRadioTuner_ConnectBandChanged(QtObjectPtr ptr);
void QRadioTuner_DisconnectBandChanged(QtObjectPtr ptr);
void QRadioTuner_CancelSearch(QtObjectPtr ptr);
int QRadioTuner_Error(QtObjectPtr ptr);
char* QRadioTuner_ErrorString(QtObjectPtr ptr);
void QRadioTuner_ConnectFrequencyChanged(QtObjectPtr ptr);
void QRadioTuner_DisconnectFrequencyChanged(QtObjectPtr ptr);
int QRadioTuner_FrequencyStep(QtObjectPtr ptr, int band);
int QRadioTuner_IsBandSupported(QtObjectPtr ptr, int band);
void QRadioTuner_ConnectMutedChanged(QtObjectPtr ptr);
void QRadioTuner_DisconnectMutedChanged(QtObjectPtr ptr);
void QRadioTuner_SearchAllStations(QtObjectPtr ptr, int searchMode);
void QRadioTuner_SearchBackward(QtObjectPtr ptr);
void QRadioTuner_SearchForward(QtObjectPtr ptr);
void QRadioTuner_ConnectSearchingChanged(QtObjectPtr ptr);
void QRadioTuner_DisconnectSearchingChanged(QtObjectPtr ptr);
void QRadioTuner_SetBand(QtObjectPtr ptr, int band);
void QRadioTuner_SetFrequency(QtObjectPtr ptr, int frequency);
void QRadioTuner_ConnectSignalStrengthChanged(QtObjectPtr ptr);
void QRadioTuner_DisconnectSignalStrengthChanged(QtObjectPtr ptr);
void QRadioTuner_Start(QtObjectPtr ptr);
void QRadioTuner_ConnectStateChanged(QtObjectPtr ptr);
void QRadioTuner_DisconnectStateChanged(QtObjectPtr ptr);
void QRadioTuner_ConnectStationFound(QtObjectPtr ptr);
void QRadioTuner_DisconnectStationFound(QtObjectPtr ptr);
void QRadioTuner_ConnectStereoStatusChanged(QtObjectPtr ptr);
void QRadioTuner_DisconnectStereoStatusChanged(QtObjectPtr ptr);
void QRadioTuner_Stop(QtObjectPtr ptr);
void QRadioTuner_ConnectVolumeChanged(QtObjectPtr ptr);
void QRadioTuner_DisconnectVolumeChanged(QtObjectPtr ptr);
void QRadioTuner_DestroyQRadioTuner(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif