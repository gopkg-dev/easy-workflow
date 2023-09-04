package model

type Task struct {
	TaskID       int    `gorm:"column:id"`            //任务ID
	BusinessID   string `gorm:"column:business_id"`   //业务ID
	ProcID       int    `gorm:"column:proc_id"`       //流程ID
	ProcName     string `gorm:"column:name"`          //流程名称
	ProcInstID   int    `gorm:"column:proc_inst_id"`  //流程实例ID
	NodeID       string `gorm:"column:node_id"`       //节点ID
	NodeName     string `gorm:"column:node_name"`     //节点名称
	PrevNodeID   string `gorm:"column:prev_node_id"`  //上一节点ID
	IsCosigned   int    `gorm:"column:is_cosigned"`   //0:任意一人通过即可 1:会签
	BatchCode    string `gorm:"column:batch_code"`    //批次码.节点会被驳回，一个节点可能产生多批task,用此码做分别
	UserID       string `gorm:"column:user_id"`       //分配用户ID
	IsPassed     int    `gorm:"column:is_passed"`     //任务是否通过 0:驳回 1:通过
	IsFinished   int    `gorm:"column:is_finished"`   // 0:任务未处理 1:处理完成
	Comment      string `gorm:"column:comment"`       //评论意见
	CreateTime   string `gorm:"column:create_time"`   //系统创建任务时间
	FinishedTime string `gorm:"column:finished_time"` // 处理任务时间
}


