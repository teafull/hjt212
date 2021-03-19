// 2005标准协议内容
package protocol

import (
	"github.com/teafull/hjt212/utils"
	"io"
	"strconv"
)

// ST 系统编码表(可扩充) (GB/T16706-1996)见《环境信息标准化手册》第一卷第 236 页
const (
	SurfaceWaterMonitoring                         = 21 // 地表水监测
	AirQualityMonitoring                           = 22 // 空气质量监测
	RegionalEnvironmentalNoiseMonitoring           = 23 // 区域环境噪声监测
	SourcesOfAtmosphericEnvironmentalPollution     = 31 // 大气环境污染源
	SurfaceWaterEnvironmentalPollution             = 32 // 地表水体环境污染源
	GroundwaterBodyEnvironmentalPollution          = 33 // 地下水体环境污染源
	OceanEnvironmentalPollution                    = 34 // 海洋环境污染源
	SourcesOfSoilEnvironmentalPollution            = 35 // 土壤环境污染源
	AcousticEnvironmentalPollutionSource           = 36 // 声环境污染源
	VibrationEnvironmentalPollutionSource          = 37 // 振动环境污染源
	RadioactiveEnvironmentalPollutionSource        = 38 // 放射性环境污染源
	SourcesOfElectromagneticEnvironmentalPollution = 41 // 电磁环境污染源
	SystemInteraction                              = 91 // 系统交互,用于现场机和上位机的交互
)

// CN
const (
	// 初始化命令
	TimeoutPeriodAndRetransmissionTimes = 1000 // 设置超时时间与重发次数 请求命令
	OverLimitAlarmTime                  = 1001 // 设置超限报警时间 请求命令

	// 参数命令
	ExtractOnSiteMachineTime       = 1011 //用于同步上位机和现场机的系统时间  请求命令&上传命令
	SetTheOnSiteMachineTime        = 1012 // 设置现场机时间 用于同步上位机和现场机的系统时间 请求命令
	ExtractPollutantAlarmThreshold = 1021 // 提取污染物报警门限值,用于污染物超标报警  请求命令&上传命令
	SetPollutantAlarmThreshold     = 1022 //设置污染物报警门限值 请求命令
	ExtractTheHostComputerAddress  = 1031 // 提取上位机地址  请求命令&上传命令
	SetTheHostComputerAddress      = 1032 // 设置上位机地址  请求命令
	ExtractedDataReportingTime     = 1041 // 提取数据上报时间  请求命令&上传命令
	SetDataReportingTime           = 1042 // 设置数据上报时间  请求命令
	ExtractRealTimeDataInterval    = 1061 // 提取实时数据间隔  请求命令&上传命令
	SetRealTimeDataInterval        = 1062 // 设置实时数据间隔  请求命令
	SetAccessPassword              = 1072 // 设置访问密码  请求命令

	// 交互命令
	RequestResponse          = 9011 //用于现场机回应上位机的请求。例如是否执行请求  上传命令
	OperationExecutionResult = 9012 // 用于现场机回应上位机的请 求的执行结果  上传命令
	NotificationResponse     = 9013 // 回应通知命令  请求命令&上传命令
	DataResponse             = 9014 // 数据应答  请求命令&上传命令

	// 数据命令-实时数据
	GetRealTimePollutantData = 2011 // 取污染物实时数据  请求命令&上传命令
	StopViewingRealTimeData  = 2012 // 停止察看实时数据，告诉现场机停止发送实时数据  请求命令

	// 数据命令-设备状态
	GetEquipmentOperatingStatusData        = 2021 // 取设备运行状态数据  请求命令&上传命令
	StopViewingEquipmentOperatingStatusDat = 2022 // 停止察看设备运行状态，告诉现场机停止发送设备运行状态数据  请求命令

	// 数据命令-历史数据
	GetHistoricalPollutantData             = 2031 // 取污染物日历史数据  请求命令&上传命令
	GetHistoricalPollutantDataOfDeviceTime = 2041 // 取设备运行时间日历史数据  请求命令&上传命令

	// 数据命令-分钟数据(可以自定义分钟间隔数,例如 5 或 10 分钟)
	GetPollutantMinuteData = 2051 // 取污染物分钟数据  请求命令&上传命令

	// 数据命令-小时数据
	TakePollutantHourlyData = 2061 // 取污染物小时数据  请求命令&上传命令

	// 数据命令-报警数据
	TakePollutantAlarmRecords = 2071 // 取污染物报警记录  请求命令&上传命令
	UploadAlarmEvent          = 2072 // 上传报警事件  请求命令&上传命

	// 控制命令
	ZeroFull                    = 3011 // 零校满 请求命令
	InstantSamplingCommand      = 3012 // 即时采样命令 请求命令
	EquipmentOperationCommand   = 3013 // 设备操作命令 请求命令
	SetDeviceSamplingTimePeriod = 3014 // 设置设备采样时间周期 请求命令
)

// 常用部分污染物相关参数编码表, 引自《环境信息标准化手册》第三卷, 定义格式： 编码:名称:应用范围
var (
	PollutantParameterCodingTable = []string{
		"B03:噪声:噪声",
		"L10:累计百分声级L10:噪声",
		"L5:累计百分声级L5:噪声",
		"L50:累计百分声级L50:噪声",
		"L90:累计百分声计L90:噪声",
		"L95:累计百分声级L95:噪声",
		"Ld:夜间等效声级Ld:噪声",
		"Ldn:昼夜等效声级Ldn:噪声",
		"Leq:30:秒等效声级Leq:噪声",
		"LMn:最小的瞬时声级:噪声",
		"LMx:最大的瞬时声级:噪声",
		"Ln:昼间等效声级:Ln:噪声",
		"S01:O.2含量:废气",
		"S02:烟气流速:废气",
		"S03:烟气温度:废气",
		"S04:烟气动压:废气",
		"S05:烟气湿度:废气",
		"S06:制冷温度:废气",
		"S07:烟道截面积:废气",
		"S08:烟气压力:废气",
		"B02:废气:废气",
		"01:烟尘:废气",
		"02:二氧化硫:废气",
		"03:氮氧化物:废气",
		"04:一氧化碳:废气",
		"05:硫化氢:废气",
		"06:氟化物:废气",
		"07:氰化物(含氰化氢):废气",
		"08:氯化氢:废气",
		"09:沥青烟:废气",
		"10:氨:废气",
		"11:氯气:废气",
		"12:二硫化碳:废气",
		"13:硫醇:废气",
		"14:硫酸雾:废气",
		"15:铬酸雾:废气",
		"16:苯系物:废气",
		"17:甲苯:废气",
		"18:二甲苯:废气",
		"19:甲醛:废气",
		"20:苯并(a)芘:废气",
		"21:苯胺类:废气",
		"22:硝基苯类:废气",
		"23:氯苯类:废气",
		"24:光气:废气",
		"25:碳氢化合物(含非甲烷总烃):废气",
		"26:乙醛:废气",
		"27:酚类:废气",
		"28:甲醇:废气",
		"29:氯乙烯:废气",
		"30:二氧化碳:废气",
		"31:汞及其化合物:废气",
		"32:铅及其化合物:废气",
		"33:镉及其化合物:废气",
		"34:锡及其化合物:废气",
		"35:镍及其化合物:废气",
		"36:铍及其化合物:废气",
		"37:林格曼黑度:废气",
		"99:其他气污染物:废气",
		"B01:污水:污水",
		"001:pH:值:污水",
		"002:色度:污水",
		"003:悬浮物:污水",
		"010:生化需氧量(BOD5)",
		"011:化学需氧量(CODcr)",
		"015:总有机碳:污水",
		"020:总汞:污水",
		"021:烷基汞:污水",
		"022:总镉:污水",
		"023:总铬:污水",
		"024:六价铬:污水",
		"025:三价铬:污水",
		"026:总砷:污水",
		"027:总铅:污水",
		"028:总镍:污水",
		"029:总铜:污水",
		"030:总锌:污水",
		"031:总锰:污水",
		"032:总铁:污水",
		"033:总银:污水",
		"034:总铍:污水",
		"035:总硒:污水",
		"036:锡:污水",
		"037:硼:污水",
		"038:钼:污水",
		"039:钡:污水",
		"040:钴:污水",
		"041:铊:污水",
		"060:氨氮:污水",
		"061:有机氮:污水",
		"065:总氮:污水",
		"080:石油类:污水",
		"101:总磷:污水",
	}
)

// 编写一些常用的2005协议中的函数，对外提供常用控制方法

//1.设置现场机访问密码
func makeSetPw(MN, oldPW, NewPw string, ST int) []byte {
	hjt212Cmd := Hjt212Cmd{
		QN:     utils.GetQN(),
		ST:     ST,
		CN:     SetAccessPassword,
		MN:     []byte(MN),
		PW:     []byte(oldPW),
		Flag:   3,
		Params: map[string]string{"PW": NewPw},
	}

	h := &HjtEncoder{}
	pwCmd, _ := h.Encoder(hjt212Cmd)
	return pwCmd
}

func SetPW(MN, oldPW, NewPw string, ST int, dest io.ReadWriteCloser) {
	// step one
	req := makeSetPw(MN, oldPW, NewPw, ST)
	dest.Write(req) // 上位机 发送设置现场机访问密码

	// step two, 请求应答
	reqRsp := make([]byte, 256)
	rspLen, err := dest.Read(reqRsp)
	if rspLen < 0 || err != nil {
		return
	}

	// 解析请求应答 ST=91;CN=9011;PW=123456;MN=88888880000001;Flag=0;CP=&&QN=20040516010101001;QnRtn=1&&
	h := &HjtDecoder{}
	hjt212Package, err := h.Decoder(reqRsp)
	if err != nil || hjt212Package.ST != ST || string(hjt212Package.MN) != MN || hjt212Package.CN != RequestResponse {
		return
	}
	if len(hjt212Package.Package) > 0 {
		QnRtn, ok := hjt212Package.Package[0]["QnRtn"]
		if QnRtnI, _ := strconv.Atoi(QnRtn); QnRtnI != 1 || !ok {
			return
		}
	} else {
		return
	}

	// step three, 返回操作执行结果
	reqRsp = make([]byte, 256)
	rspLen, err = dest.Read(reqRsp)
	if rspLen < 0 || err != nil {
		return
	}

	// 解析操作执行结果
}
