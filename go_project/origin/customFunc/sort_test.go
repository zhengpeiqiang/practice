package customFunc_test

import (
	"fmt"
	"origin/customFunc"
	"strings"
	"testing"
)

func TestStringLengthSort(t *testing.T) {
	str := `var wkbGroupMemberService = &service.WkbGroupMemberService{}
var wkbImCustomerServiceGroupService = &service.WkbImCustomerServiceGroupService{}
var wkbImRobotService = &service.WkbImRobotService{}
var wkbGroupService = &service.WkbGroupService{}
var wkbUserDetailService = &service.WkbUserDetailService{}
var wkbGroupChatMessageService = &service.WkbGroupChatMessageService{}
var wkbImMergeMessageService = &service.WkbImMergeMessageService{}
var wkbUserService = &service.WkbUserService{}
var wkbRealNameAuthService = &service.WkbRealNameAuthService{}
var wkbImAchievementService = &service.WkbImAchievementService{}
var wkbImConfigService = &service.WkbImConfigService{}
var wkbImReplySetting = &service.WkbImReplySetting{}
var wkbImKeywordService = &service.WkbImKeywordService{}`
	newStr := strings.Split(str, "\n")

	sortArr := make([]int, 0)
	for _, v := range newStr {
		sortArr = append(sortArr, len(v))
	}
	customFunc.Sort("asc", "quick", sortArr)
	fmt.Println(sortArr)

	customFunc.Sort("asc", "quick", newStr)
	for _, v := range newStr {
		fmt.Println(v)
	}
}