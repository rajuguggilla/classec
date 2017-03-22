CREATE DATABASE IF NOT EXISTS `GO_DB_Testing` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;

USE `GO_DB_Testing`;

CREATE TABLE `azure_cpu` (
  `name` varchar(200) NOT NULL,
  `minimum` varchar(200) DEFAULT NULL,
  `maximum` varchar(200) DEFAULT NULL,
  `average` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `azure_dynamic` (
  `name` varchar(200) DEFAULT NULL,
  `timestamp` varchar(200) DEFAULT NULL,
  `minimum` varchar(200) DEFAULT NULL,
  `maximum` varchar(200) DEFAULT NULL,
  `average` varchar(200) DEFAULT NULL,
  `unit` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `azure_instances` (
  `subscriptionid` varchar(100) DEFAULT NULL,
  `name` varchar(150) NOT NULL DEFAULT '',
  `type` varchar(150) DEFAULT NULL,
  `location` varchar(150) DEFAULT NULL,
  `vmsize` varchar(150) DEFAULT NULL,
  `publisher` varchar(150) DEFAULT NULL,
  `offer` varchar(150) DEFAULT NULL,
  `sku` varchar(150) DEFAULT NULL,
  `vmid` varchar(150) DEFAULT NULL,
  `availabilitysetid` varchar(150) DEFAULT NULL,
  `provisioningstate` varchar(150) DEFAULT NULL,
  `resourcegroupname` varchar(150) DEFAULT NULL,
  `status` varchar(200) DEFAULT NULL,
  `storage` int(200) DEFAULT NULL,
  `ram` int(200) DEFAULT NULL,
  `numcpu` int(200) DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `dynamic_instances` (
  `Vm_Name` varchar(200) NOT NULL,
  `InstanceID` varchar(100) NOT NULL,
  `Count` int(50) NOT NULL,
  `DurationStart` varchar(100) NOT NULL,
  `Min` float NOT NULL,
  `DurationEnd` varchar(100) NOT NULL,
  `Max` float NOT NULL,
  `Sum` float NOT NULL,
  `Period` int(200) NOT NULL,
  `PeriodEnd` varchar(100) NOT NULL,
  `Duration` float NOT NULL,
  `PeriodStart` varchar(100) NOT NULL,
  `Avg` float NOT NULL,
  `Groupby` varchar(100) NOT NULL,
  `Unit` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



CREATE TABLE `hos_cpu` (
  `name` varchar(200) NOT NULL,
  `min` varchar(200) DEFAULT NULL,
  `max` varchar(200) DEFAULT NULL,
  `avg` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `hos_dynamic_instances` (
  `Name` varchar(250) DEFAULT NULL,
  `Instance_id` varchar(250) DEFAULT NULL,
  `Count` int(50) DEFAULT NULL,
  `Duration_start` timestamp(5) NULL DEFAULT NULL,
  `Duration_end` timestamp(5) NULL DEFAULT NULL,
  `Min` float DEFAULT NULL,
  `Max` float DEFAULT NULL,
  `Sum` float DEFAULT NULL,
  `Period` int(50) DEFAULT NULL,
  `Period_end` timestamp(5) NULL DEFAULT NULL,
  `Duration` int(11) DEFAULT NULL,
  `Period_start` timestamp(5) NULL DEFAULT NULL,
  `Avg` float DEFAULT NULL,
  `Unit` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `hos_instances` (
  `Name` varchar(200) NOT NULL,
  `Instance_id` varchar(200) NOT NULL,
  `Flavor_id` varchar(200) DEFAULT NULL,
  `Flavor_Name` varchar(200) DEFAULT NULL,
  `Status` varchar(200) DEFAULT NULL,
  `Image` varchar(200) DEFAULT NULL,
  `Security_Group` varchar(200) DEFAULT NULL,
  `Availability_Zone` varchar(200) DEFAULT NULL,
  `ip_address` varchar(20) DEFAULT NULL,
  `keypair_name` varchar(200) DEFAULT NULL,
  `ram` int(200) DEFAULT NULL,
  `vcpu` int(200) DEFAULT NULL,
  `disk` int(200) DEFAULT NULL,
  `deleted` varchar(50) DEFAULT NULL,
  `classifier` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `hos_test` (
  `Name` varchar(25) NOT NULL,
  `Instance_id` varchar(250) DEFAULT NULL,
  `Count` int(50) DEFAULT NULL,
  `Duration_end` varchar(250) DEFAULT NULL,
  `Min` int(50) DEFAULT NULL,
  `Test` varchar(50) DEFAULT NULL,
  `Duration_start` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `instances` (
  `name` varchar(255) NOT NULL,
  `instance_id` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `availability_zone` varchar(255) DEFAULT NULL,
  `flavor` varchar(255) DEFAULT NULL,
  `flavor_id` varchar(255) DEFAULT NULL,
  `ram` varchar(20) DEFAULT NULL,
  `vcpu` varchar(20) DEFAULT NULL,
  `storage` varchar(20) DEFAULT NULL,
  `ip_address` varchar(200) DEFAULT NULL,
  `security_group` varchar(255) DEFAULT NULL,
  `keypair_name` varchar(255) DEFAULT NULL,
  `image_name` varchar(255) DEFAULT NULL,
  `insertion_date` varchar(200) DEFAULT NULL,
  `CreationTime` varchar(255) DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT NULL,
  `classifier` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `openstack_cpu` (
  `name` varchar(200) NOT NULL,
  `min` varchar(200) DEFAULT NULL,
  `max` varchar(200) DEFAULT NULL,
  `avg` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `tags` (
  `InstanceId` varchar(200) NOT NULL,
  `InstanceName` varchar(200) DEFAULT NULL,
  `Cloud` varchar(200) DEFAULT NULL,
  `Tagname` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `vmware_cpu` (
  `name` varchar(200) NOT NULL,
  `min` varchar(200) DEFAULT NULL,
  `max` varchar(200) DEFAULT NULL,
  `avg` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `vmware_dynamic_details` (
  `Name` varchar(50) DEFAULT NULL,
  `Uuid` varchar(50) DEFAULT NULL,
  `Timestamp` varchar(100) DEFAULT NULL,
  `MaxCpuUsage` varchar(20) DEFAULT NULL,
  `AvgCpuUsage` varchar(20) DEFAULT NULL,
  `MinCpuUsage` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE `azure_cpu`
  ADD PRIMARY KEY (`name`);

ALTER TABLE `azure_instances`
  ADD PRIMARY KEY (`name`);

ALTER TABLE `hos_cpu`
  ADD PRIMARY KEY (`name`);

ALTER TABLE `hos_instances`
  ADD PRIMARY KEY (`Instance_id`);

ALTER TABLE `instances`
  ADD PRIMARY KEY (`name`);

ALTER TABLE `openstack_cpu`
  ADD PRIMARY KEY (`name`);

ALTER TABLE `tags`
  ADD PRIMARY KEY (`InstanceId`);

ALTER TABLE `vmware_cpu`
  ADD PRIMARY KEY (`name`);

ALTER TABLE `vmware_instances`
  ADD PRIMARY KEY (`Name`);


CREATE DATABASE IF NOT EXISTS `cloud_assessment` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;

USE `cloud_assessment`;

CREATE TABLE IF NOT EXISTS `rds_dynamic` (
  `DBInstanceIdentifier` varchar(50) NOT NULL DEFAULT '',
  `StartTime` datetime DEFAULT NULL,
  `EndTime` datetime DEFAULT NULL,
  `Period` bigint(20) DEFAULT NULL,
  `BinLogDiskUsage_min` double DEFAULT NULL,
  `BinLogDiskUsage_max` double DEFAULT NULL,
  `BinLogDiskUsage_avg` double DEFAULT NULL,
  `CPUUtilization_min` double DEFAULT NULL,
  `CPUUtilization_max` double DEFAULT NULL,
  `CPUUtilization_avg` double DEFAULT NULL,
  `CPUCreditUsage_min` double DEFAULT NULL,
  `CPUCreditUsage_max` double DEFAULT NULL,
  `CPUCreditUsage_avg` double DEFAULT NULL,
  `CPUCreditBalance_min` double DEFAULT NULL,
  `CPUCreditBalance_max` double DEFAULT NULL,
  `CPUCreditBalance_avg` double DEFAULT NULL,
  `DatabaseConnections_min` double DEFAULT NULL,
  `DatabaseConnections_max` double DEFAULT NULL,
  `DatabaseConnections_avg` double DEFAULT NULL,
  `DiskQueueDepth_min` double DEFAULT NULL,
  `DiskQueueDepth_max` double DEFAULT NULL,
  `DiskQueueDepth_avg` double DEFAULT NULL,
  `FreeableMemory_min` double DEFAULT NULL,
  `FreeableMemory_max` double DEFAULT NULL,
  `FreeableMemory_avg` double DEFAULT NULL,
  `FreeStorageSpace_min` double DEFAULT NULL,
  `FreeStorageSpace_max` double DEFAULT NULL,
  `FreeStorageSpace_avg` double DEFAULT NULL,
  `ReplicaLag_min` double DEFAULT NULL,
  `ReplicaLag_max` double DEFAULT NULL,
  `ReplicaLag_avg` double DEFAULT NULL,
  `SwapUsage_min` double DEFAULT NULL,
  `SwapUsage_max` double DEFAULT NULL,
  `SwapUsage_avg` double DEFAULT NULL,
  `ReadIOPS_min` double DEFAULT NULL,
  `ReadIOPS_max` double DEFAULT NULL,
  `ReadIOPS_avg` double DEFAULT NULL,
  `WriteIOPS_min` double DEFAULT NULL,
  `WriteIOPS_max` double DEFAULT NULL,
  `WriteIOPS_avg` double DEFAULT NULL,
  `ReadLatency_min` double DEFAULT NULL,
  `ReadLatency_max` double DEFAULT NULL,
  `ReadLatency_avg` double DEFAULT NULL,
  `WriteLatency_min` double DEFAULT NULL,
  `WriteLatency_max` double DEFAULT NULL,
  `WriteLatency_avg` double DEFAULT NULL,
  `ReadThroughput_min` double DEFAULT NULL,
  `ReadThroughput_max` double DEFAULT NULL,
  `ReadThroughput_avg` double DEFAULT NULL,
  `WriteThroughput_min` double DEFAULT NULL,
  `WriteThroughput_max` double DEFAULT NULL,
  `WriteThroughput_avg` double DEFAULT NULL,
  `NetworkReceiveThroughput_min` double DEFAULT NULL,
  `NetworkReceiveThroughput_max` double DEFAULT NULL,
  `NetworkReceiveThroughput_avg` double DEFAULT NULL,
  `NetworkTransmitThroughput_min` double DEFAULT NULL,
  `NetworkTransmitThroughput_max` double DEFAULT NULL,
  `NetworkTransmitThroughput_avg` double DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `vw_rds` (
`API_Name` varchar(30)
,`Linux_On_Demand_cost` varchar(30)
,`Linux_Reserved_cost` varchar(30)
,`Windows_On_Demand_cost` varchar(30)
,`Windows_Reserved_cost` varchar(30)
);