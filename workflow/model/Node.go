package model

type NodeType int

const (
	RootNode    NodeType = 0 //开始(根)节点
	TaskNode    NodeType = 1 //任务节点,指的是需要人完成的节点
	GateWayNode NodeType = 2 //参考activiti的网关.此项目中使用混合网关,等于activiti中排他、并行网关、包含网关的混合体
	EndNode     NodeType = 3 //结束节点,结束节点不需要人参与，到了此节点，则流程实例完成
)

type Node struct {
	NodeID      string        //节点名称
	NodeName    string        //节点名字
	NodeType    NodeType      //节点类型 0:开始节点 1:任务节点,指的是需要人完成的节点 2:网关 3:结束节点
	PrevNodeIDs []string      //上级节点(不管是任务节点还是结束节点，因为分支的存在，所以它的上级节点可能都会有多个)
	UserIDs     []string      //节点处理人数组
	Roles       []string      //节点处理角色数组。注意，因为系统无法预先知道角色中存在多少用户，所以必须用StartEvents解析角色，将角色中的用户加到UserIDs中
	GWConfig    HybridGateway //网关。只有在节点类型为GateWay的情况下此字段才会有值
	IsCosigned  int8          //是否会签  只有任务节点才会用到，会签的情况下需要所有任务通过才能进行下一节点，只要有一人反对，则整个节点驳回
	StartEvents []string      //节点开始时触发的事件
	EndEvents   []string      //节点结束时触发的事件
}

/*思考
一、开始节点与结束节点的设计思想：
1、BPMBN 2.0标准是必须要有开始和结束节点的。
2、在java世界常用的activiti流程引擎中，开始节点并不是一个任务节点，这也符合BPMBN标准。
按照activiti的流程图设计一个请假流程是这样的：开始节点-->开始请假(任务节点)。用过activiti的都知道，流程开始之后“开始请假”节点并不会自动结束，一般
需要写一个事件判断是不是开始节点，是则自动结束。这就很麻烦，而且不符合人类思维：我都开始流程了，难道还有不想让流程继续下去的？
3、所以，我设计的流程引擎中，确实包含开始节点，但这个开始节点同时也是一个特殊的任务节点，流程一旦开始，这个节点就自动通过。
4、activiti中结束节点是这样的：领导审批完毕-->结束节点。
我们可以同样图省事把结束节点可以也放在任务节点上么？
能，当然能。但是会有问题：在分支较多的时候，可能由于疏忽，某个分支流转的最后一个任务节点忘了把它设置为结束节点，此时流程就卡死了。
所以，明确的设置一个结束节点，也是一种防呆设计。

二、关于事件的注意点：
1、结束节点只处理StartEvents，不处理EndEvents。因为结束节点只是把数据库表中流程数据做归档，不涉及到任务分配、节点路径分配等，所以没有必要再做结束事件处理。
2、不推荐在GateWay节点上加事件，容易使流程节点分配出错；而且对于上级节点>1个的GW节点，由于每次上级节点的task结束，都会到该GW中做一下判断（上级节点是否结束，
是否可以开始在GW中进行下一步节点分配），所以会执行多次StartEvents。这无疑会把事情搞复杂。

三、关于会签的注意点:
会签节点：全部通过才能通过，一人驳回即驳回
非会签节点：一人通过即通过，一人驳回即驳回
*/
