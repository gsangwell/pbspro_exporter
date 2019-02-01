package collector

import (
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/taylor840326/go_pbspro/qstat"
)

func init() {
	registerCollector(qstatCollectorSubSystem, defaultEnabled, NewQstatCollector)
}

type qstatCollector struct {
	server_state string
}

func (c *qstatCollector) Update(ch chan<- prometheus.Metric) error {
	log.Infoln("Update Qstat Server Status")
	c.updateQstatServer(ch)
	log.Infoln("Update Qstat Queue Status")
	c.updateQstatQueue(ch)
	log.Infoln("Update Qstat Node Status")
	c.updateQstatNode(ch)
	log.Infoln("Update Qstat Jobs Status")
	c.updateQstatJobs(ch)
	return nil
}

type qstatMetric struct {
	name            string
	desc            string
	value           float64
	metricType      prometheus.ValueType
	extraLabel      []string
	extraLabelValue string
}

func NewQstatCollector() (Collector, error) {
	qc := new(qstatCollector)
	return &qstatCollector{server_state: qc.server_state}, nil
}

func (c *qstatCollector) updateQstatServer(ch chan<- prometheus.Metric) {

	var allMetrics []qstatMetric
	//var metrics []qstatMetric
	var labelsValue []string

	qstat, err := qstat.NewQstat(*pbsproURL)
	if err != nil {
		log.Errorln("Create New Qstat Failed ", err.Error())
	}

	qstat.SetAttribs(nil)
	qstat.SetExtend("")

	log.Infoln("Connecting PBS Server ..")
	err = qstat.ConnectPBS()
	if err != nil {
		log.Fatalln("Connecting PBS Server Failed. ", err.Error())
	}
	defer qstat.DisconnectPBS()

	err = qstat.PbsServerState()
	if err != nil {
		log.Errorln("Gather PBS Server Informations Failed", err.Error())
	}

	for _, ss := range qstat.ServerState {
		allMetrics = []qstatMetric{
			{
				name:       "server_state",
				desc:       "pbspro_exporter: server state. 1 is Active",
				value:      float64(ss.ServerState),
				metricType: prometheus.GaugeValue,
			},
			{

				name:       "server_scheduling",
				desc:       "pbspro_exporter: Server Scheduling. 1 is True",
				value:      float64(ss.ServerScheduling),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_total_jobs",
				desc:       "pbspro_exporter: Server Total Jobs.",
				value:      float64(ss.TotalJobs),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_transit_state_count",
				desc:       "pbspro_exporter: Server Transit State Count.",
				value:      float64(ss.StateCountTransit),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_queued_state_count",
				desc:       "pbspro_exporter: Server Queued State Count.",
				value:      float64(ss.StateCountQueued),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_held_state_count",
				desc:       "pbspro_exporter: Server Held State Count.",
				value:      float64(ss.StateCountHeld),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_waiting_state_count",
				desc:       "pbspro_exporter: Server Waiting State Count.",
				value:      float64(ss.StateCountWaiting),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_running_state_count",
				desc:       "pbspro_exporter: Server Running State Count.",
				value:      float64(ss.StateCountRunning),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_exiting_state_count",
				desc:       "pbspro_exporter: Server Exiting State Count.",
				value:      float64(ss.StateCountExiting),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_begun_state_count",
				desc:       "pbspro_exporter: Server Begun State Count.",
				value:      float64(ss.StateCountBegun),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_log_events",
				desc:       "pbspro_exporter: Server Log Events.",
				value:      float64(ss.LogEvents),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_query_other_jobs",
				desc:       "pbspro_exporter: Server Query Other Jobs. 1 is True",
				value:      float64(ss.QueryOtherJobs),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_resources_default_ncpus",
				desc:       "pbspro_exporter: Server Resources Default Ncpus.",
				value:      float64(ss.ResourcesDefaultNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_default_chunk_ncpus",
				desc:       "pbspro_exporter: Server Default Chunk Ncpus.",
				value:      float64(ss.DefaultChunkNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_resources_assigned_ncpus",
				desc:       "pbspro_exporter: Server Resources Assigned Ncpus.",
				value:      float64(ss.ResourcesAssignedNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_resources_assigned_nodect",
				desc:       "pbspro_exporter: Server Resources Assigned Nodect.",
				value:      float64(ss.ResourcesAssignedNodect),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_scheduler_iteration",
				desc:       "pbspro_exporter: Server Scheudler Iteration.",
				value:      float64(ss.SchedulerIteration),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_flicenses",
				desc:       "pbspro_exporter: Server Flicense.",
				value:      float64(ss.Flicenses),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_resv_enable",
				desc:       "pbspro_exporter: Server Resv Enable. 1 is True",
				value:      float64(ss.ResvEnable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_node_fail_requeue",
				desc:       "pbspro_exporter: Server Node Fail Requeue.",
				value:      float64(ss.NodeFailRequeue),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_max_array_size",
				desc:       "pbspro_exporter: Server Max Array Size.",
				value:      float64(ss.MaxArraySize),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_pbs_license_min",
				desc:       "pbspro_exporter: Server PBS License Min.",
				value:      float64(ss.PBSLicenseMin),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_pbs_license_max",
				desc:       "pbspro_exporter: Server PBS License Max.",
				value:      float64(ss.PBSLicenseMax),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_pbs_license_linger_time",
				desc:       "pbspro_exporter: Server PBS License Linger Time.",
				value:      float64(ss.PBSLicenseLingerTime),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_license_count_avail_global",
				desc:       "pbspro_exporter: Server License Count Avail Global.",
				value:      float64(ss.LicenseCountAvailGlobal),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_license_count_avail_local",
				desc:       "pbspro_exporter: Server License Count Avail Global.",
				value:      float64(ss.LicenseCountAvailLocal),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_license_count_used",
				desc:       "pbspro_exporter: Server License Used.",
				value:      float64(ss.LicenseCountUsed),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_license_count_high_use",
				desc:       "pbspro_exporter: Server License Count High Use.",
				value:      float64(ss.LicenseCountHighUse),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_eligible_time_enable",
				desc:       "pbspro_exporter: Server Eligible Time Enable.1 is True",
				value:      float64(ss.EligibleTimeEnable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_job_history_enable",
				desc:       "pbspro_exporter: Server Job History Enable.1 is True",
				value:      float64(ss.JobHistoryEnable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_job_history_duration",
				desc:       "pbspro_exporter: Server Job History Duration.",
				value:      float64(ss.JobHistoryDuration),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_max_concurrent_provision",
				desc:       "pbspro_exporter: Server Max Concurrent Provision.",
				value:      float64(ss.MaxConcurrentProvision),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "server_power_provisioning",
				desc:       "pbspro_exporter: Server Power Provisioning. 1 is True",
				value:      float64(ss.PowerProvisioning),
				metricType: prometheus.GaugeValue,
			},
		}
		labelsValue = []string{ss.ServerName, ss.ServerHost, ss.DefaultQueue, ss.MailFrom, ss.PBSVersion}
	}

	for _, m := range allMetrics {

		labelsName := []string{"ServerName", "ServerHost", "DefaultQueue", "MailFrom", "PBSVersion"}

		desc := prometheus.NewDesc(
			prometheus.BuildFQName(namespace, qstatCollectorSubSystem, m.name),
			m.desc,
			labelsName,
			nil,
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			m.metricType,
			m.value,
			labelsValue...,
		)
	}

}

func (c *qstatCollector) updateQstatQueue(ch chan<- prometheus.Metric) {

	var allMetrics []qstatMetric
	//var metrics []qstatMetric
	var labelsValue []string

	qstat, err := qstat.NewQstat(*pbsproURL)
	if err != nil {
		log.Errorln("Create New Qstat Failed. ", err.Error())
	}

	qstat.SetAttribs(nil)
	qstat.SetExtend("")

	log.Infoln("Connecting PBS Server ..")
	err = qstat.ConnectPBS()
	if err != nil {
		log.Fatalln("Connect PBS Server Failed. ", err.Error())
	}
	defer qstat.DisconnectPBS()

	err = qstat.PbsQueueState()
	if err != nil {
		log.Errorln("Update Queue State Failed. ", err.Error())
	}

	for _, ss := range qstat.QueueState {
		allMetrics = []qstatMetric{
			{
				name:       "queue_total_jobs",
				desc:       "pbspro_exporter: Queue Total Jobs.",
				value:      float64(ss.TotalJobs),
				metricType: prometheus.GaugeValue,
			},
			{

				name:       "queue_transit_state_count",
				desc:       "pbspro_exporter: Queue Transit State Count.",
				value:      float64(ss.StateCountTransit),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_queued_state_count",
				desc:       "pbspro_exporter: Queue Queued State Count.",
				value:      float64(ss.StateCountQueued),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_held_state_count",
				desc:       "pbspro_exporter: Queue Held State Count.",
				value:      float64(ss.StateCountHeld),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_waiting_state_count",
				desc:       "pbspro_exporter: Queue Waiting State Count.",
				value:      float64(ss.StateCountWaiting),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_running_state_count",
				desc:       "pbspro_exporter: Queue Running State Count.",
				value:      float64(ss.StateCountRunning),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_exiting_state_count",
				desc:       "pbspro_exporter: Queue Exiting State Count.",
				value:      float64(ss.StateCountExiting),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_begun_state_count",
				desc:       "pbspro_exporter: Queue Begun State Count.",
				value:      float64(ss.StateCountBegun),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_resources_assigned_ncpus",
				desc:       "pbspro_exporter: Queue Resources Assigned Ncpus.",
				value:      float64(ss.ResourcesAssignedNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_resources_assigned_nodect",
				desc:       "pbspro_exporter: Queue Resources Assigned Nodect.",
				value:      float64(ss.ResourcesAssignedNodect),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_enable",
				desc:       "pbspro_exporter: Queue Enable. 1 is True",
				value:      float64(ss.Enable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "queue_started",
				desc:       "pbspro_exporter: Queue Started. 1 is True",
				value:      float64(ss.Started),
				metricType: prometheus.GaugeValue,
			},
		}
		labelsValue = []string{ss.QueueName, ss.QueueType}
	}

	for _, m := range allMetrics {

		labelsName := []string{"QueueName", "QueueType"}

		desc := prometheus.NewDesc(
			prometheus.BuildFQName(namespace, qstatCollectorSubSystem, m.name),
			m.desc,
			labelsName,
			nil,
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			m.metricType,
			m.value,
			labelsValue...,
		)
	}

}

func (c *qstatCollector) updateQstatNode(ch chan<- prometheus.Metric) {

	var allMetrics []qstatMetric
	//var metrics []qstatMetric
	var labelsValue []string

	qstat, err := qstat.NewQstat(*pbsproURL)
	if err != nil {
		log.Errorln("Create New Qstat Failed. ", err.Error())
	}

	qstat.SetAttribs(nil)
	qstat.SetExtend("")

	log.Infoln("Connecting PBS Server.. ")
	err = qstat.ConnectPBS()
	if err != nil {
		log.Fatalln("Connect PBS Server Failed. ", err.Error())
	}
	defer qstat.DisconnectPBS()

	err = qstat.PbsNodeState()
	if err != nil {
		log.Errorln("Update Node State Failed ", err.Error())
	}

	for _, ss := range qstat.NodeState {
		allMetrics = []qstatMetric{
			{
				name:       "node_pcpus",
				desc:       "pbspro_exporter: Node Pcpus.",
				value:      float64(ss.Pcpus),
				metricType: prometheus.GaugeValue,
			},
			{

				name:       "node_resources_available_mem",
				desc:       "pbspro_exporter: Node Resources Available Mem",
				value:      float64(ss.ResourcesAvailableMem),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "node_resources_available_ncpus",
				desc:       "pbspro_exporter: Node Resources Available Ncpus.",
				value:      float64(ss.ResourcesAvailableNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "node_resources_assigned_accelerator_memory",
				desc:       "pbspro_exporter: Node Resources Assigned Accelerator Memory.",
				value:      float64(ss.ResourcesAssignedAcceleratorMemory),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "node_resources_assigned_hbmem",
				desc:       "pbspro_exporter: Node Resources Assigned HBmem.",
				value:      float64(ss.ResourcesAssignedHbmem),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "node_resources_assigned_mem",
				desc:       "pbspro_exporter: Node Resources Assigned Mem.",
				value:      float64(ss.ResourcesAssignedMem),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "node_resources_assigned_naccelerators",
				desc:       "pbspro_exporter: Node Resources Assigned Naccelerators.",
				value:      float64(ss.ResourcesAssignedNaccelerators),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "node_resources_assigned_ncpus",
				desc:       "pbspro_exporter: Node Resources Assigned Ncpus.",
				value:      float64(ss.ResourcesAssignedNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "node_resources_assigned_vmem",
				desc:       "pbspro_exporter: Node Resources Assigned Vmem.",
				value:      float64(ss.ResourcesAssignedVmem),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "node_resv_enable",
				desc:       "pbspro_exporter: Node Resv Enable. 1 is True",
				value:      float64(ss.ResvEnable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "node_last_change_time",
				desc:       "pbspro_exporter: Node Last Change Time",
				value:      float64(ss.LastStateChangeTime),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "node_last_used_time",
				desc:       "pbspro_exporter: Node Last Used Time",
				value:      float64(ss.LastUsedTime),
				metricType: prometheus.GaugeValue,
			},
		}
		labelsValue = []string{ss.NodeName, ss.Mom, ss.Ntype, ss.State, ss.Jobs, ss.ResourcesAvailableArch, ss.ResourcesAvailableHost, ss.ResourcesAvailableApplications, ss.ResourcesAvailablePlatform, ss.ResourcesAvailableSoftware, ss.ResourcesAvailableVnodes, ss.Sharing}
	}

	for _, m := range allMetrics {

		labelsName := []string{"NodeName", "Mom", "Ntype", "NodeState", "RunningJobs", "ResourcesAvailableArch", "ResourcesAvailableHost", "ResourcesAvailableApplications", "ResourcesAvailablePlatform", "ResourcesAvailableSoftware", "ResourcesAvailableVnodes", "Sharing"}

		desc := prometheus.NewDesc(
			prometheus.BuildFQName(namespace, qstatCollectorSubSystem, m.name),
			m.desc,
			labelsName,
			nil,
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			m.metricType,
			m.value,
			labelsValue...,
		)
	}

}

func (c *qstatCollector) updateQstatJobs(ch chan<- prometheus.Metric) {

	var allMetrics []qstatMetric
	var metrics []qstatMetric
	var labelsValue []string

	qstat, err := qstat.NewQstat(*pbsproURL)
	if err != nil {
		log.Errorln("Create New Qstat Failed. ", err.Error())
	}

	qstat.SetAttribs(nil)
	qstat.SetExtend("")

	log.Infoln("Connecting PBS Server..")
	err = qstat.ConnectPBS()
	if err != nil {
		log.Fatalln("Connect PBS Server Failed. ", err.Error())
	}
	defer qstat.DisconnectPBS()

	err = qstat.PbsJobsState()
	if err != nil {
		log.Errorln("Update Jobs State Failed. ", err.Error())
	}

	for _, ss := range qstat.JobsState {
		metrics = []qstatMetric{
			{
				name:       "jobs_resources_used_cpupercent",
				desc:       "pbspro_exporter: Jobs Resources Used CpuPercent.",
				value:      ss.ResourcesUsedCpuPercent,
				metricType: prometheus.GaugeValue,
			},
			{

				name:       "jobs_resources_used_cput",
				desc:       "pbspro_exporter: Jobs Resources Used Cput",
				value:      float64(ss.ResourcesUsedCput),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_resources_used_mem",
				desc:       "pbspro_exporter: Jobs Resources Used Mem.",
				value:      float64(ss.ResourcesUsedMem),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_resources_used_ncpus",
				desc:       "pbspro_exporter: Jobs Resources Used Ncpus.",
				value:      float64(ss.ResourcesUsedNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_resources_used_vmem",
				desc:       "pbspro_exporter: Jobs Resources Used Vmem.",
				value:      float64(ss.ResourcesUsedVmem),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_resources_used_walltime",
				desc:       "pbspro_exporter: Jobs Resources Used WallTime.",
				value:      float64(ss.ResourcesUsedWallTime),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_ctime",
				desc:       "pbspro_exporter: Jobs Çtime.",
				value:      float64(ss.Ctime),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_mtime",
				desc:       "pbspro_exporter: Jobs Mtime.",
				value:      float64(ss.Mtime),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_priority",
				desc:       "pbspro_exporter: Jobs Priority.",
				value:      float64(ss.Priority),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_qtime",
				desc:       "pbspro_exporter: Jobs Qtime",
				value:      float64(ss.Qtime),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_rerunable",
				desc:       "pbspro_exporter: Jobs Rerunable",
				value:      float64(ss.Rerunable),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_resources_list_ncpus",
				desc:       "pbspro_exporter: Jobs Resources List Ncpus",
				value:      float64(ss.ResourceListNcpus),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_resources_list_nodect",
				desc:       "pbspro_exporter: Jobs Resources List Nodect",
				value:      float64(ss.ResourceListNodect),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_resources_list_walltime",
				desc:       "pbspro_exporter: Jobs Resources List WallTime",
				value:      float64(ss.ResourceListWallTime),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_stime",
				desc:       "pbspro_exporter: Jobs stime",
				value:      float64(ss.Stime),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_sessionid",
				desc:       "pbspro_exporter: Jobs Session ID",
				value:      float64(ss.SessionID),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_substate",
				desc:       "pbspro_exporter: Jobs SubState",
				value:      float64(ss.SubState),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_etime",
				desc:       "pbspro_exporter: Jobs Etime",
				value:      float64(ss.Etime),
				metricType: prometheus.GaugeValue,
			},
			{
				name:       "jobs_runcount",
				desc:       "pbspro_exporter: Jobs RunCount",
				value:      float64(ss.RunCount),
				metricType: prometheus.GaugeValue,
			},
		}
		labelsValue = []string{ss.JobName,
			strings.Replace(ss.JobOwner, "@", "_", -1),
			ss.JobState,
			ss.Queue,
			ss.Server,
			ss.CheckPoint,
			ss.ErrorPath,
			ss.ExecHost,
			ss.ExecVnode,
			ss.HoldType,
			ss.JoinPath,
			ss.KeepFiles,
			ss.MailPoints,
			ss.OutputPath,
			ss.ResourceListPlace,
			ss.ResourceListSelect,
			ss.ResourceListSoftware,
			ss.JobDir,
			ss.VariableList,
			strings.Replace(ss.VariableListHome, "/", "-1", -1),
			strings.Replace(ss.VariableListLang, ".", "_", -1),
			ss.VariableListLogname,
			ss.VariableListPath,
			ss.VariableListMail,
			ss.VariableListShell,
			ss.VariableListWorkdir,
			ss.VariableListSystem,
			ss.VariableListQueue,
			ss.VariableListHost,
			ss.Comment,
			ss.SubmitArguments,
			ss.Project,
			time.Now().String(),
		}

		allMetrics = append(allMetrics, metrics...)
	}

	for _, m := range allMetrics {

		labelsName := []string{"JobName",
			"JobOwner",
			"JobState",
			"Queue",
			"Server",
			"CheckPoint",
			"ErrorPath",
			"ExecHost",
			"ExecVnode",
			"HoldType",
			"JoinPath",
			"KeepFiles",
			"MailPoints",
			"OutputPath",
			"ResourceListPlace",
			"ResourceListSelect",
			"ResourceListSoftware",
			"JobDir",
			"VariableList",
			"VariableListHome",
			"VariableListLang",
			"VariableListLogname",
			"VariableListPath",
			"VariableListMail",
			"VariableListShell",
			"VariableListWrokdir",
			"VariableListSystem",
			"VariableListQueue",
			"VariableListHost",
			"Comment",
			"SubmitArguments",
			"Project",
			"LocalTime",
		}
		desc := prometheus.NewDesc(
			prometheus.BuildFQName(namespace, qstatCollectorSubSystem, m.name),
			m.desc,
			labelsName,
			nil,
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			m.metricType,
			m.value,
			labelsValue...,
		)
	}

}
