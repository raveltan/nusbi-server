-- MySQL dump 10.13  Distrib 8.0.22, for Linux (x86_64)
--
-- Host: localhost    Database: nusbiam
-- ------------------------------------------------------
-- Server version       8.0.22-0ubuntu0.20.04.2

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
-- Table structure for table `Answers`
--

DROP TABLE IF EXISTS `Answers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Answers` (
                           `answer_id` varchar(60) NOT NULL,
                           `student_id` varchar(60) NOT NULL,
                           `question_id` varchar(60) NOT NULL,
                           `answer` varchar(600) DEFAULT NULL,
                           `score` tinyint DEFAULT NULL,
                           `submitted_on` datetime DEFAULT NULL,
                           PRIMARY KEY (`answer_id`),
                           KEY `student_id` (`student_id`),
                           KEY `question_id` (`question_id`),
                           CONSTRAINT `Answers_ibfk_1` FOREIGN KEY (`student_id`) REFERENCES `Students` (`student_id`) ON DELETE CASCADE,
                           CONSTRAINT `Answers_ibfk_2` FOREIGN KEY (`question_id`) REFERENCES `Questions` (`question_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Answers`
--

LOCK TABLES `Answers` WRITE;
/*!40000 ALTER TABLE `Answers` DISABLE KEYS */;
/*!40000 ALTER TABLE `Answers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Attendance`
--

DROP TABLE IF EXISTS `Attendance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Attendance` (
                              `attendance_id` varchar(60) NOT NULL,
                              `student_id` varchar(60) NOT NULL,
                              `schedule_id` varchar(60) NOT NULL,
                              `status` tinyint NOT NULL,
                              PRIMARY KEY (`attendance_id`),
                              KEY `student_id` (`student_id`),
                              KEY `schedule_id` (`schedule_id`),
                              CONSTRAINT `Attendance_ibfk_1` FOREIGN KEY (`student_id`) REFERENCES `Students` (`student_id`) ON DELETE CASCADE,
                              CONSTRAINT `Attendance_ibfk_2` FOREIGN KEY (`schedule_id`) REFERENCES `Schedules` (`schedule_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Attendance`
--

LOCK TABLES `Attendance` WRITE;
/*!40000 ALTER TABLE `Attendance` DISABLE KEYS */;
/*!40000 ALTER TABLE `Attendance` ENABLE KEYS */;
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
/*!40000 ALTER TABLE `Courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Enrolled_Courses`
--

DROP TABLE IF EXISTS `Enrolled_Courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Enrolled_Courses` (
                                    `enrolled_id` varchar(60) NOT NULL,
                                    `student_id` varchar(60) NOT NULL,
                                    `class_id` varchar(60) NOT NULL,
                                    `mid_score` tinyint DEFAULT NULL,
                                    `final_score` tinyint DEFAULT NULL,
                                    `enrolled_date` date NOT NULL,
                                    PRIMARY KEY (`enrolled_id`),
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
INSERT INTO `Majors` VALUES ('3019ee5d-2555-4fa2-b9c5-948bda3b63e9','internaet system'),('343d40ca-b59e-4938-8f9c-4efaf0f5e8ff','Computer Science'),('8eb24d8d-1df7-4665-b53f-f2d5f9b05f7a','International business'),('a1b01cdc-6fc2-44f7-a6b8-6c0ea3f2134a','Information system');
/*!40000 ALTER TABLE `Majors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Questions`
--

DROP TABLE IF EXISTS `Questions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Questions` (
                             `question_id` varchar(60) NOT NULL,
                             `class_id` varchar(60) NOT NULL,
                             `question` varchar(300) NOT NULL,
                             `due_date` date NOT NULL,
                             PRIMARY KEY (`question_id`),
                             KEY `class_id` (`class_id`),
                             CONSTRAINT `Questions_ibfk_1` FOREIGN KEY (`class_id`) REFERENCES `Class` (`class_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Questions`
--

LOCK TABLES `Questions` WRITE;
/*!40000 ALTER TABLE `Questions` DISABLE KEYS */;
/*!40000 ALTER TABLE `Questions` ENABLE KEYS */;
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
                            `gpa` decimal(5,2) DEFAULT NULL,
                            `scu` tinyint DEFAULT NULL,
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
INSERT INTO `Users` VALUES ('ravel','$2a$07$rKtdH7dfNFQ5trlCQhSOz.GNnFsUe1gXkfWeT7hoJ6p7KcsQOSI/y','a'),('tanjaya','$2a$07$lLm4za0ARCUMxRZ5rbL3O.arxS6qdwxfD1jz6Xl3KWEvh06nDHusi','a');
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

-- Dump completed on 2020-11-20 17:57:59
