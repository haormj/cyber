package qos

import "github.com/haormj/cyber/pb"

func CreateQosProfile(history pb.QosHistoryPolicy, depth, mps uint32,
	reliability pb.QosReliabilityPolicy, durability pb.QosDurabilityPolicy) pb.QosProfile {
	return pb.QosProfile{
		History:     &history,
		Depth:       &depth,
		Mps:         &mps,
		Reliability: &reliability,
		Durability:  &durability,
	}
}

var (
	QOS_PROFILE_DEFAULT = CreateQosProfile(
		pb.QosHistoryPolicy_HISTORY_KEEP_LAST,
		pb.Default_QosProfile_Depth,
		pb.Default_QosProfile_Mps,
		pb.QosReliabilityPolicy_RELIABILITY_RELIABLE,
		pb.QosDurabilityPolicy_DURABILITY_VOLATILE,
	)

	QOS_PROFILE_SENSOR_DATA = CreateQosProfile(
		pb.QosHistoryPolicy_HISTORY_KEEP_LAST,
		5,
		pb.Default_QosProfile_Mps,
		pb.QosReliabilityPolicy_RELIABILITY_BEST_EFFORT,
		pb.QosDurabilityPolicy_DURABILITY_VOLATILE,
	)

	QOS_PROFILE_PARAMETERS = CreateQosProfile(
		pb.QosHistoryPolicy_HISTORY_KEEP_LAST,
		1000,
		pb.Default_QosProfile_Mps,
		pb.QosReliabilityPolicy_RELIABILITY_RELIABLE,
		pb.QosDurabilityPolicy_DURABILITY_VOLATILE,
	)

	QOS_PROFILE_SERVICES_DEFAULT = CreateQosProfile(
		pb.QosHistoryPolicy_HISTORY_KEEP_LAST,
		10,
		pb.Default_QosProfile_Mps,
		pb.QosReliabilityPolicy_RELIABILITY_RELIABLE,
		pb.QosDurabilityPolicy_DURABILITY_TRANSIENT_LOCAL,
	)

	QOS_PROFILE_PARAM_EVENT = CreateQosProfile(
		pb.QosHistoryPolicy_HISTORY_KEEP_LAST,
		1000,
		pb.Default_QosProfile_Mps,
		pb.QosReliabilityPolicy_RELIABILITY_RELIABLE,
		pb.QosDurabilityPolicy_DURABILITY_VOLATILE,
	)

	QOS_PROFILE_SYSTEM_DEFAULT = CreateQosProfile(
		pb.QosHistoryPolicy_HISTORY_SYSTEM_DEFAULT,
		pb.Default_QosProfile_Depth,
		pb.Default_QosProfile_Mps,
		pb.QosReliabilityPolicy_RELIABILITY_RELIABLE,
		pb.QosDurabilityPolicy_DURABILITY_TRANSIENT_LOCAL,
	)

	QOS_PROFILE_TF_STATIC = CreateQosProfile(
		pb.QosHistoryPolicy_HISTORY_KEEP_ALL,
		10,
		pb.Default_QosProfile_Depth,
		pb.QosReliabilityPolicy_RELIABILITY_RELIABLE,
		pb.QosDurabilityPolicy_DURABILITY_TRANSIENT_LOCAL,
	)

	QOS_PROFILE_TOPO_CHANGE = CreateQosProfile(
		pb.QosHistoryPolicy_HISTORY_KEEP_ALL,
		10,
		pb.Default_QosProfile_Depth,
		pb.QosReliabilityPolicy_RELIABILITY_RELIABLE,
		pb.QosDurabilityPolicy_DURABILITY_TRANSIENT_LOCAL,
	)
)
