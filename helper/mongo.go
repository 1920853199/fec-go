package helper

import (
   
   
)

// 得到当前时间对应的 mongodb   的 dbName
// nowDateStr 格式：  2009-12-12
func GetDbName(dbName string, nowDateStr string) (string){
    if nowDateStr == "" {
        nowDateStr = DateUTCStr()
    }
    return dbName + "_" + nowDateStr
}
// 得到当前时间对应的 mongodb   的 collection Name
func GetCollName(collName string, websiteId string) string {
    return collName + "_" + websiteId
}


// 初始数据接收的数据库
func GetTraceDbName() (string){
    return GetDbName("trace", "")
}

// dateStr 格式：  2009-12-12 
// 通过时间，得到相应时间的库
func GetTraceDbNameByDate(dateStr string) string {
    return GetDbName("trace", dateStr)
}

func GetTraceDataCollName(websiteId string) (string){
    return GetCollName("trace_data", websiteId)
}

// 得到Browser统计后的数据输出的collection
func GetOutWholeBrowserCollName(websiteId string) (string){
    return GetCollName("trace_whole_browser_data", websiteId)
}

// 得到Browser统计后的数据输出的collection
func GetOutWholeAllCollName(websiteId string) (string){
    return GetCollName("trace_whole_all_data", websiteId)
}

// 得到Refer统计后的数据输出的collection
func GetOutWholeReferCollName(websiteId string) (string){
    return GetCollName("trace_whole_refer_data", websiteId)
}

// 得到 country 统计后的数据输出的collection
func GetOutWholeCountryCollName(websiteId string) (string){
    return GetCollName("trace_whole_country_data", websiteId)
}









