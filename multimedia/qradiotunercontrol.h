#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QRadioTunerControl_ConnectAntennaConnectedChanged(QtObjectPtr ptr);
void QRadioTunerControl_DisconnectAntennaConnectedChanged(QtObjectPtr ptr);
int QRadioTunerControl_Band(QtObjectPtr ptr);
void QRadioTunerControl_ConnectBandChanged(QtObjectPtr ptr);
void QRadioTunerControl_DisconnectBandChanged(QtObjectPtr ptr);
void QRadioTunerControl_CancelSearch(QtObjectPtr ptr);
int QRadioTunerControl_Error(QtObjectPtr ptr);
char* QRadioTunerControl_ErrorString(QtObjectPtr ptr);
int QRadioTunerControl_Frequency(QtObjectPtr ptr);
void QRadioTunerControl_ConnectFrequencyChanged(QtObjectPtr ptr);
void QRadioTunerControl_DisconnectFrequencyChanged(QtObjectPtr ptr);
int QRadioTunerControl_FrequencyStep(QtObjectPtr ptr, int band);
int QRadioTunerControl_IsAntennaConnected(QtObjectPtr ptr);
int QRadioTunerControl_IsBandSupported(QtObjectPtr ptr, int band);
int QRadioTunerControl_IsMuted(QtObjectPtr ptr);
int QRadioTunerControl_IsSearching(QtObjectPtr ptr);
int QRadioTunerControl_IsStereo(QtObjectPtr ptr);
void QRadioTunerControl_ConnectMutedChanged(QtObjectPtr ptr);
void QRadioTunerControl_DisconnectMutedChanged(QtObjectPtr ptr);
void QRadioTunerControl_SearchAllStations(QtObjectPtr ptr, int searchMode);
void QRadioTunerControl_SearchBackward(QtObjectPtr ptr);
void QRadioTunerControl_SearchForward(QtObjectPtr ptr);
void QRadioTunerControl_ConnectSearchingChanged(QtObjectPtr ptr);
void QRadioTunerControl_DisconnectSearchingChanged(QtObjectPtr ptr);
void QRadioTunerControl_SetBand(QtObjectPtr ptr, int band);
void QRadioTunerControl_SetFrequency(QtObjectPtr ptr, int frequency);
void QRadioTunerControl_SetMuted(QtObjectPtr ptr, int muted);
void QRadioTunerControl_SetStereoMode(QtObjectPtr ptr, int mode);
void QRadioTunerControl_SetVolume(QtObjectPtr ptr, int volume);
int QRadioTunerControl_SignalStrength(QtObjectPtr ptr);
void QRadioTunerControl_ConnectSignalStrengthChanged(QtObjectPtr ptr);
void QRadioTunerControl_DisconnectSignalStrengthChanged(QtObjectPtr ptr);
void QRadioTunerControl_Start(QtObjectPtr ptr);
void QRadioTunerControl_ConnectStateChanged(QtObjectPtr ptr);
void QRadioTunerControl_DisconnectStateChanged(QtObjectPtr ptr);
void QRadioTunerControl_ConnectStationFound(QtObjectPtr ptr);
void QRadioTunerControl_DisconnectStationFound(QtObjectPtr ptr);
int QRadioTunerControl_StereoMode(QtObjectPtr ptr);
void QRadioTunerControl_ConnectStereoStatusChanged(QtObjectPtr ptr);
void QRadioTunerControl_DisconnectStereoStatusChanged(QtObjectPtr ptr);
void QRadioTunerControl_Stop(QtObjectPtr ptr);
int QRadioTunerControl_Volume(QtObjectPtr ptr);
void QRadioTunerControl_ConnectVolumeChanged(QtObjectPtr ptr);
void QRadioTunerControl_DisconnectVolumeChanged(QtObjectPtr ptr);
void QRadioTunerControl_DestroyQRadioTunerControl(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif