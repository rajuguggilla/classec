CREATE DATABASE IF NOT EXISTS `GO_DB_Testing` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;
USE `GO_DB_Testing`;

CREATE TABLE IF NOT EXISTS `azure_dynamic` (
  `name` varchar(200) NOT NULL,
  `timestamp` varchar(200) NOT NULL,
  `minimum` varchar(200) NOT NULL,
  `maximum` varchar(200) NOT NULL,
  `average` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;




CREATE TABLE IF NOT EXISTS `azure_instances` (
  `subscriptionid` varchar(100) DEFAULT NULL,
  `name` varchar(150) NOT NULL DEFAULT '',
  `type` varchar(150) DEFAULT NULL,
  `location` varchar(150) DEFAULT NULL,
  `vmsize` varchar(150) DEFAULT NULL,
  `publisher` varchar(150) DEFAULT NULL,
  `offer` varchar(150) DEFAULT NULL,
  `sku` varchar(150) DEFAULT NULL,
  `vmid` varchar(150) NOT NULL DEFAULT '',
  `availabilitysetid` varchar(150) DEFAULT NULL,
  `provisioningstate` varchar(150) DEFAULT NULL,
  `resourcegroupname` varchar(150) DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT NULL,
  `tagname` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



CREATE TABLE IF NOT EXISTS `hos_instances` (
  `Name` varchar(200) NOT NULL,
  `Instance_id` varchar(200) NOT NULL,
  `Flavor_id` varchar(200) NOT NULL,
  `Flavor_Name` varchar(200) NOT NULL,
  `Status` varchar(200) NOT NULL,
  `Image` varchar(200) NOT NULL,
  `Security_Group` varchar(200) NOT NULL,
  `Availability_Zone` varchar(200) NOT NULL,
  `ip_address` varchar(20) NOT NULL,
  `keypair_name` varchar(200) NOT NULL,
  `ram` int(200) NOT NULL,
  `vcpu` int(200) NOT NULL,
  `disk` int(200) NOT NULL,
  `deleted` varchar(50) NOT NULL,
  `tagname` varchar(200) NOT NULL,
  PRIMARY KEY (`Instance_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;



CREATE TABLE IF NOT EXISTS `instances` (
  `name` varchar(255) NOT NULL,
  `instance_id` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL,
  `availability_zone` varchar(255) NOT NULL,
  `flavor` varchar(255) NOT NULL,
  `flavor_id` varchar(255) NOT NULL,
  `ram` varchar(20) NOT NULL,
  `vcpu` varchar(20) NOT NULL,
  `storage` varchar(20) NOT NULL,
  `ip_address` decimal(20,0) NOT NULL,
  `security_group` varchar(255) NOT NULL,
  `keypair_name` varchar(255) NOT NULL,
  `image_name` varchar(255) DEFAULT NULL,
  `volumes` int(4) DEFAULT NULL,
  `insertion_date` date NOT NULL,
  `CreationTime` varchar(255) DEFAULT NULL,
  `deleted` tinyint(1) NOT NULL,
  `tagname` varchar(200) NOT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE IF NOT EXISTS `providers` (
  `InstanceId` varchar(200) NOT NULL,
  `Cloud` varchar(200) NOT NULL,
  `Tagname` varchar(200) NOT NULL,
  PRIMARY KEY (`InstanceId`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE IF NOT EXISTS `vmware_dynamic_details` (
  `Name` varchar(50) NOT NULL,
  `Uuid` varchar(50) NOT NULL,
  `Timestamp` varchar(100) NOT NULL,
  `MaxCpuUsage` varchar(20) NOT NULL,
  `AvgCpuUsage` varchar(20) NOT NULL,
  `MinCouUsage` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE IF NOT EXISTS `vmware_instances` (
  `Name` varchar(200) NOT NULL DEFAULT '',
  `Uuid` varchar(200) NOT NULL DEFAULT '',
  `MemorySizeMB` int(200) DEFAULT NULL,
  `PowerState` varchar(200) DEFAULT NULL,
  `NumofCPU` int(200) DEFAULT NULL,
  `GuestFullName` varchar(200) DEFAULT NULL,
  `IPaddress` varchar(200) DEFAULT NULL,
  `deleted` tinyint(1) NOT NULL,
  `tagname` varchar(200) NOT NULL,
  `classifier` varchar(100) NOT NULL,
  PRIMARY KEY (`Name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


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