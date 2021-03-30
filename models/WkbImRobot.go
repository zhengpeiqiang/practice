package models

import "time"

// 机器人信息表
type WkbImRobot struct {
	Id int `json:"id"` // 自增主键id
	RobotCode string `json:"robot_code"` // 机器人编码
	RobotImg string `json:"robot_img"` // 机器人头像
	RobotNickname string `json:"robot_nickname"` // 机器人昵称
	WelcomeFlag int `json:"welcome_flag"` // 欢迎语开启状态:1=开启,2=关闭
	WelcomeText string `json:"welcome_text"` // 欢迎语
	AnswerFlag int `json:"answer_flag"` // 应答开启状态:1=开启,2=关闭
	AnswerTitle string `json:"answer_title"` // 应答标题
	AnswerText string `json:"answer_text"` // 应答文本
	KeywordAutoAnswer int `json:"keyword_auto_answer"` // 是否开启关键词自动回复:1=是,2=否
	KeywordAutoLink int `json:"keyword_auto_link"` // 是否开启关键词自动跳转:1=是,2=否
	Deleted int `json:"deleted"` // 是否删除:1=删除,2=否
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedAt time.Time `json:"created_at"` // 创建时间
}