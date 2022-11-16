package getPsReport

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/GoStruct"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
)

const Url = "/v1/reportService/getPsReport"
const Disabled = false

type RequestData struct {
	ReportType valueTypes.Integer `json:"report_type" required:"true"`
	DateId     valueTypes.DateTime `json:"date_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return GoStruct.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("report_type: One of 1, 2, 5")
	ret += api.HelpDataType()
	return ret
}

type ResultData struct {
	PsDateRangeTotalData TotalData `json:"ps_date_range_total_data"`
	PsTotalData          TotalData `json:"ps_total_data"`
	ReportListData       struct {
		PageList      []struct {
			GoStruct           GoStruct.GoStruct  `json:"-" PointDeviceFrom:"PsId"`

			PsId               valueTypes.PsId    `json:"ps_id"`
			PsName             valueTypes.String  `json:"ps_name"`
			PsType             valueTypes.Integer `json:"ps_type"`
			ShareType          valueTypes.String  `json:"share_type"`
			Co2Reduce          valueTypes.Float   `json:"co2_reduce" PointUnitFromParent:"PsTotalCo2ReduceUnit" PointUpdateFreq:"UpdateFreqTotal"`
			PowerQuantityTotal valueTypes.Float   `json:"power_quantity_total" PointUnitFromParent:"PowerQuantityTotalUnit" PointUpdateFreq:"UpdateFreqTotal"`
			PsInstalledPower   valueTypes.Float   `json:"ps_installed_power" PointUnitFromParent:"PsInstalledPowerUnit"`
			PsTotalCo2Reduce   valueTypes.Float   `json:"ps_total_co2_reduce" PointUnitFromParent:"PsTotalCo2ReduceUnit" PointUpdateFreq:"UpdateFreqTotal"`
			PsTotalPower       valueTypes.Float   `json:"ps_total_power" PointUnitFromParent:"PsTotalPowerUnit" PointUpdateFreq:"UpdateFreqTotal"`
			TotalProfit        valueTypes.Float   `json:"total_profit" PointUnitFrom:"IncomeUnitName"`
			IncomeUnitName     valueTypes.String  `json:"income_unit_name" PointIgnore:"true"`
			ProfitList         []Profit           `json:"profit_list" PointIdReplace:"true"`
		} `json:"pageList" PointId:"page_list" PointIdFromChild:"PsId" PointIdReplace:"true"`
		Co2ReduceUnit          valueTypes.String  `json:"co2_reduce_unit" PointIgnore:"true"`
		PowerQuantityTotalUnit valueTypes.String  `json:"power_quantity_total_unit" PointIgnore:"true"`
		PsInstalledPowerUnit   valueTypes.String  `json:"ps_installed_power_unit" PointIgnore:"true"`
		PsTotalCo2ReduceUnit   valueTypes.String  `json:"ps_total_co2_reduce_unit" PointIgnore:"true"`
		PsTotalPowerUnit       valueTypes.String  `json:"ps_total_power_unit" PointIgnore:"true"`
		RowCount               valueTypes.Count   `json:"rowCount" PointId:"row_count"`
	} `json:"report_list_data" PointId:"total_by_ps" PointIdReplace:"true"`
}

type TotalData struct {
	MinDateId                      valueTypes.DateTime  `json:"min_date_id"`

	Co2Reduce                      valueTypes.Float     `json:"co2_reduce" PointUnitFrom:"Co2ReduceOriginalUnit"`
	Co2ReduceMap                   valueTypes.UnitValue `json:"co2_reduce_map"`
	Co2ReduceOriginalUnit          valueTypes.String    `json:"co2_reduce_original_unit" PointIgnore:"true"`

	PowerQuantityTotal             valueTypes.Float     `json:"power_quantity_total" PointUnitFrom:"PowerQuantityTotalOriginalUnit"`
	PowerQuantityTotalMap          valueTypes.UnitValue `json:"power_quantity_total_map"`
	PowerQuantityTotalOriginalUnit valueTypes.String    `json:"power_quantity_total_original_unit" PointIgnore:"true"`

	TotalProfit                    valueTypes.Float     `json:"total_profit" PointUnitFrom:"IncomeUnitName"`
	ProfitList                     []Profit             `json:"profit_list" PointIdReplace:"true"`
}

type Profit struct {
	IncomeUnitId   valueTypes.Integer   `json:"income_unit_id"`
	IncomeUnitName valueTypes.String  `json:"income_unit_name" PointIgnore:"true"`
	TotalProfit    valueTypes.Float `json:"total_profit" PointUnitFrom:"IncomeUnitName"`
}

func (e *ResultData) IsValid() error {
	var err error
	return err
}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "", GoStruct.NewEndPointPath(e.Request.DateId.Format(valueTypes.DateTimeLayoutDay)))
	}

	return entries
}
