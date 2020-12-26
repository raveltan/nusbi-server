-- MySQL dump 10.13  Distrib 8.0.22, for Linux (x86_64)
--
-- Host: localhost    Database: nusbiam
-- ------------------------------------------------------
-- Server version	8.0.22-0ubuntu0.20.04.3

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Absence`
--

DROP TABLE IF EXISTS `Absence`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Absence` (
  `student_id` varchar(60) NOT NULL,
  `schedule_id` varchar(60) NOT NULL,
  KEY `student_id` (`student_id`),
  KEY `schedule_id` (`schedule_id`),
  CONSTRAINT `Absence_ibfk_1` FOREIGN KEY (`student_id`) REFERENCES `Students` (`student_id`) ON DELETE CASCADE,
  CONSTRAINT `Absence_ibfk_2` FOREIGN KEY (`schedule_id`) REFERENCES `Schedules` (`schedule_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Absence`
--

LOCK TABLES `Absence` WRITE;
/*!40000 ALTER TABLE `Absence` DISABLE KEYS */;
/*!40000 ALTER TABLE `Absence` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Class`
--

DROP TABLE IF EXISTS `Class`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Class` (
  `class_id` varchar(60) NOT NULL,
  `class_name` varchar(40) NOT NULL,
  `course_id` varchar(60) NOT NULL,
  `batch` smallint NOT NULL,
  PRIMARY KEY (`class_id`),
  KEY `course_id` (`course_id`),
  CONSTRAINT `Class_ibfk_1` FOREIGN KEY (`course_id`) REFERENCES `Courses` (`course_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Class`
--

LOCK TABLES `Class` WRITE;
/*!40000 ALTER TABLE `Class` DISABLE KEYS */;
INSERT INTO `Class` VALUES ('0e08f5f2-9f64-4215-a163-52be08157dd1','l3ac','678c1ba3-a073-4fb1-bf5c-fb31a9f334f6',2020),('30298851-8bab-4fae-937a-7f9fd3f9ff54','l3ac','8545e0e2-9fff-4de6-95cf-731388e83a9a',2020),('6316e712-1c20-4c46-919a-cfe294d3f0c0','l3ac','80e7fa33-b13b-430e-a3b4-3e2bce46c1f5',2020),('68a73a70-5062-4bae-96ff-aecac9fc5697','lobc','011fd7a2-a317-49fe-8881-17cfe3338019',2011),('bffb3538-76fb-4678-a255-eb00053797b3','l4ac','678c1ba3-a073-4fb1-bf5c-fb31a9f334f6',2020);
/*!40000 ALTER TABLE `Class` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Courses`
--

DROP TABLE IF EXISTS `Courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Courses` (
  `course_id` varchar(60) NOT NULL,
  `course_name` varchar(40) NOT NULL,
  `lecturer_id` varchar(60) NOT NULL,
  `scu` tinyint NOT NULL,
  PRIMARY KEY (`course_id`),
  KEY `lecturer_id` (`lecturer_id`),
  CONSTRAINT `Courses_ibfk_1` FOREIGN KEY (`lecturer_id`) REFERENCES `Lecturers` (`lecturer_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Courses`
--

LOCK TABLES `Courses` WRITE;
/*!40000 ALTER TABLE `Courses` DISABLE KEYS */;
INSERT INTO `Courses` VALUES ('011fd7a2-a317-49fe-8881-17cfe3338019','Networking','0ef5cbbe-2106-4080-b618-41cffa2813d7',4),('678c1ba3-a073-4fb1-bf5c-fb31a9f334f6','intro to programming','0ef5cbbe-2106-4080-b618-41cffa2813d7',5),('80e7fa33-b13b-430e-a3b4-3e2bce46c1f5','Computer forensics','0ef5cbbe-2106-4080-b618-41cffa2813d7',3),('8545e0e2-9fff-4de6-95cf-731388e83a9a','computational math','28294550-c206-4517-a8d5-54f322eb56bc',4);
/*!40000 ALTER TABLE `Courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Enrolled_Courses`
--

DROP TABLE IF EXISTS `Enrolled_Courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Enrolled_Courses` (
  `student_id` varchar(60) NOT NULL,
  `class_id` varchar(60) NOT NULL,
  `mid_score` tinyint DEFAULT NULL,
  `final_score` tinyint DEFAULT NULL,
  KEY `student_id` (`student_id`),
  KEY `class_id` (`class_id`),
  CONSTRAINT `Enrolled_Courses_ibfk_1` FOREIGN KEY (`student_id`) REFERENCES `Students` (`student_id`) ON DELETE CASCADE,
  CONSTRAINT `Enrolled_Courses_ibfk_2` FOREIGN KEY (`class_id`) REFERENCES `Class` (`class_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Enrolled_Courses`
--

LOCK TABLES `Enrolled_Courses` WRITE;
/*!40000 ALTER TABLE `Enrolled_Courses` DISABLE KEYS */;
INSERT INTO `Enrolled_Courses` VALUES ('656d6bf2-d8da-4870-b808-e1b2d15d732d','0e08f5f2-9f64-4215-a163-52be08157dd1',97,88),('1a567f11-66e5-4e38-b8e0-5bffa53b1af8','6316e712-1c20-4c46-919a-cfe294d3f0c0',99,98),('14009124-ef12-45e0-8b64-4004d40ad95f','0e08f5f2-9f64-4215-a163-52be08157dd1',-1,-1),('1a567f11-66e5-4e38-b8e0-5bffa53b1af8','bffb3538-76fb-4678-a255-eb00053797b3',-1,-1),('14009124-ef12-45e0-8b64-4004d40ad95f','30298851-8bab-4fae-937a-7f9fd3f9ff54',-1,-1),('1a567f11-66e5-4e38-b8e0-5bffa53b1af8','68a73a70-5062-4bae-96ff-aecac9fc5697',-1,-1),('14009124-ef12-45e0-8b64-4004d40ad95f','6316e712-1c20-4c46-919a-cfe294d3f0c0',99,55),('6de353ec-a038-4dba-8200-2386e9abdb83','68a73a70-5062-4bae-96ff-aecac9fc5697',-1,-1),('6de353ec-a038-4dba-8200-2386e9abdb83','bffb3538-76fb-4678-a255-eb00053797b3',-1,-1);
/*!40000 ALTER TABLE `Enrolled_Courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Lecturers`
--

DROP TABLE IF EXISTS `Lecturers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Lecturers` (
  `lecturer_id` varchar(60) NOT NULL,
  `first_name` varchar(40) NOT NULL,
  `last_name` varchar(40) NOT NULL,
  `gender` char(1) NOT NULL,
  `dob` date NOT NULL,
  `email` varchar(40) NOT NULL,
  `user_id` varchar(60) NOT NULL,
  PRIMARY KEY (`lecturer_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `Lecturers_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `Users` (`user_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Lecturers`
--

LOCK TABLES `Lecturers` WRITE;
/*!40000 ALTER TABLE `Lecturers` DISABLE KEYS */;
INSERT INTO `Lecturers` VALUES ('0ef5cbbe-2106-4080-b618-41cffa2813d7','sumarno','berjaya','M','2003-12-18','sumarno@nosebee.com','sumarno'),('28294550-c206-4517-a8d5-54f322eb56bc','yowen','yowen','M','2000-12-28','yowen@nosebee.com','yowen');
/*!40000 ALTER TABLE `Lecturers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Majors`
--

DROP TABLE IF EXISTS `Majors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Majors` (
  `major_id` varchar(60) NOT NULL,
  `major_name` varchar(40) NOT NULL,
  PRIMARY KEY (`major_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Majors`
--

LOCK TABLES `Majors` WRITE;
/*!40000 ALTER TABLE `Majors` DISABLE KEYS */;
INSERT INTO `Majors` VALUES ('2309d801-287a-4f28-895f-6f10de9797ad','Network system'),('3019ee5d-2555-4fa2-b9c5-948bda3b63e9','internaet system'),('343d40ca-b59e-4938-8f9c-4efaf0f5e8ff','Computer Science'),('8eb24d8d-1df7-4665-b53f-f2d5f9b05f7a','International business'),('a1b01cdc-6fc2-44f7-a6b8-6c0ea3f2134a','Information system');
/*!40000 ALTER TABLE `Majors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Schedules`
--

DROP TABLE IF EXISTS `Schedules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Schedules` (
  `schedule_id` varchar(60) NOT NULL,
  `date_time` datetime NOT NULL,
  `class_id` varchar(60) NOT NULL,
  PRIMARY KEY (`schedule_id`),
  KEY `class_id` (`class_id`),
  CONSTRAINT `Schedules_ibfk_1` FOREIGN KEY (`class_id`) REFERENCES `Class` (`class_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Schedules`
--

LOCK TABLES `Schedules` WRITE;
/*!40000 ALTER TABLE `Schedules` DISABLE KEYS */;
INSERT INTO `Schedules` VALUES ('14ddf383-9b4a-4b60-a5df-2c4ad9ccc40f','2020-12-08 00:00:00','0e08f5f2-9f64-4215-a163-52be08157dd1'),('1fda5502-90db-4d6e-a679-ce3bf4eca4f9','2020-12-31 00:00:00','bffb3538-76fb-4678-a255-eb00053797b3'),('3f7f6872-b4f0-470d-a82d-addc1509f164','2020-12-08 00:00:00','68a73a70-5062-4bae-96ff-aecac9fc5697'),('81428838-6d6e-4619-9317-f9078a19da9e','2020-12-14 00:00:00','30298851-8bab-4fae-937a-7f9fd3f9ff54'),('8d1240dd-e069-4b3a-86a5-9992e406e513','2020-12-23 00:00:00','0e08f5f2-9f64-4215-a163-52be08157dd1'),('9c205387-4447-478e-bb25-fdc1e1d0f4b5','2020-12-16 00:00:00','30298851-8bab-4fae-937a-7f9fd3f9ff54'),('ecfe85d3-f3c9-4639-9c08-b6ae4cca6802','2020-12-18 00:00:00','6316e712-1c20-4c46-919a-cfe294d3f0c0'),('f806e819-efb9-4d6f-9423-e96d59cafb79','2020-12-08 00:00:00','6316e712-1c20-4c46-919a-cfe294d3f0c0');
/*!40000 ALTER TABLE `Schedules` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Students`
--

DROP TABLE IF EXISTS `Students`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Students` (
  `student_id` varchar(60) NOT NULL,
  `first_name` varchar(40) NOT NULL,
  `last_name` varchar(40) NOT NULL,
  `gender` char(1) NOT NULL,
  `dob` date NOT NULL,
  `email` varchar(40) NOT NULL,
  `user_id` varchar(60) NOT NULL,
  `major_id` varchar(60) NOT NULL,
  `batch` smallint NOT NULL,
  PRIMARY KEY (`student_id`),
  KEY `user_id` (`user_id`),
  KEY `major_id` (`major_id`),
  CONSTRAINT `Students_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `Users` (`user_id`) ON DELETE CASCADE,
  CONSTRAINT `Students_ibfk_2` FOREIGN KEY (`major_id`) REFERENCES `Majors` (`major_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Students`
--

LOCK TABLES `Students` WRITE;
/*!40000 ALTER TABLE `Students` DISABLE KEYS */;
INSERT INTO `Students` VALUES ('14009124-ef12-45e0-8b64-4004d40ad95f','william','marugame','M','2000-12-29','william@nosebee.com','william','a1b01cdc-6fc2-44f7-a6b8-6c0ea3f2134a',2020),('1a567f11-66e5-4e38-b8e0-5bffa53b1af8','vincent','tandra','M','2020-12-15','vincent@nosebee.com','vincent','343d40ca-b59e-4938-8f9c-4efaf0f5e8ff',2020),('656d6bf2-d8da-4870-b808-e1b2d15d732d','sutardi','berjaya','F','2006-12-21','sutardi@nosebee.com','sutardi','343d40ca-b59e-4938-8f9c-4efaf0f5e8ff',2019),('6de353ec-a038-4dba-8200-2386e9abdb83','johny','sumarno','M','1985-12-19','johny@john.com','johny','a1b01cdc-6fc2-44f7-a6b8-6c0ea3f2134a',2023);
/*!40000 ALTER TABLE `Students` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Users`
--

DROP TABLE IF EXISTS `Users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Users` (
  `user_id` varchar(60) NOT NULL,
  `password` varchar(60) NOT NULL,
  `role` char(1) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Users`
--

LOCK TABLES `Users` WRITE;
/*!40000 ALTER TABLE `Users` DISABLE KEYS */;
INSERT INTO `Users` VALUES ('johny','$2a$07$Xt4JFNRtxqVzurnksm.sSuqjRH..umbS5i7cGxtvMThz2N2wV5yGq','s'),('ravel','$2a$07$rKtdH7dfNFQ5trlCQhSOz.GNnFsUe1gXkfWeT7hoJ6p7KcsQOSI/y','a'),('sumarno','$2a$07$qWRfW8Vp/04ADYxvwnPcfe88yUS21CFpkUzJjUgR.2JMYgcyav5sK','t'),('sutardi','$2a$07$9RpbkNeVTQgMeXDQQf4CmubPb3.TupjtHiFime/wS5k3f3fWxhdNC','s'),('tanjaya','$2a$07$lLm4za0ARCUMxRZ5rbL3O.arxS6qdwxfD1jz6Xl3KWEvh06nDHusi','a'),('vincent','$2a$07$jQMdlDwfs86i0lhPaT3zAOBhHuv6iThSrPffKxtuv7sEHIZlu832O','s'),('william','$2a$07$1NSVhshtRBTxYHIXJJHLPOwtfvNUi3lfNaJK4EjnvctoQ76AMNOxW','s'),('yowen','$2a$07$QmVEdrljotwNMLvcJu2f3O9UkNsjWm/k62aNBcu9sXd0ZhnSNe5oq','t');
/*!40000 ALTER TABLE `Users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-12-26 23:02:06
