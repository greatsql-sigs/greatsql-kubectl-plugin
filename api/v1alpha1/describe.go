package v1alpha1

type Describe struct {
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	UID        string            `json:"uid"`
	Age        string            `json:"age"`
	Replicas   int               `json:"replicas"`
	Image      string            `json:"image"`
	Members    []Members         `json:"members"`
	Conditions []ConditionStatus `json:"conditions"`
	Primary    string            `json:"primary"`
	ReadyCount int               `json:"readyCount"`
}

type Members struct {
	Name     string `json:"name"`     // e.g. "mycluster-0"
	Role     string `json:"role"`     // "Arbitrator" / "Voter"
	Mode     string `json:"mode"`     // 读写类型，R/O / R/W
	Storage  string `json:"storage"`  // PVC name or ephemeral
	Affinity string `json:"affinity"` // 节点亲和性简要描述
}

type ConditionStatus struct {
	Type    string `json:"type"`   // e.g. Ready, Healthy
	Status  string `json:"status"` // True/False
	Reason  string `json:"reason"`
	Message string `json:"message"`
}
