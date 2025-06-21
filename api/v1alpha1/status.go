package v1alpha1

type Status struct {
	Cluster      string         `json:"cluster"`
	Namespace    string         `json:"namespace"`
	Replicas     int            `json:"replicas"`
	TopologyMode string         `json:"topologyMode"` // 集群模式
	Nodes        []MemberStatus `json:"nodes"`
}

// MemberStatus 表示单个节点（Pod）的运行与复制状态
type MemberStatus struct {
	Name                string `json:"name"`                // Pod 名称
	Role                string `json:"role"`                // Primary / Secondary
	State               string `json:"state"`               // ONLINE / RECOVERING / OFFLINE
	Ready               bool   `json:"ready"`               // Pod 是否就绪
	Image               string `json:"image"`               // 容器镜像
	UptimeSeconds       int64  `json:"uptimeSeconds"`       // 启动时长
	TrxToBeCertified    int64  `json:"trxToBeCertified"`    // 待认证事务数量
	RelayLogToBeApplied int64  `json:"relayLogToBeApplied"` // Relay log 待应用事务数
	TrxChecked          int64  `json:"trxChecked"`          // 已校验事务
	TrxDone             int64  `json:"trxDone"`             // 已处理完成事务
	TrxProposed         int64  `json:"trxProposed"`         // 发起的事务数量
}
