CREATE DATABASE  IF NOT EXISTS `minbya3mili` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `minbya3mili`;
-- MySQL dump 10.13  Distrib 8.0.36, for Linux (x86_64)
--
-- Host: localhost    Database: minbya3mili
-- ------------------------------------------------------
-- Server version	8.0.40-0ubuntu0.24.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `images`
--

DROP TABLE IF EXISTS `images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `images` (
  `image_id` int NOT NULL AUTO_INCREMENT,
  `url` varchar(255) NOT NULL,
  `user_id` int NOT NULL,
  `listing_id` int NOT NULL DEFAULT '0',
  `show_on_profile` tinyint(1) DEFAULT '0',
  `date_created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`image_id`),
  UNIQUE KEY `url_UNIQUE` (`url`),
  KEY `fk_images_1_idx` (`user_id`),
  CONSTRAINT `fk_images_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `images`
--

LOCK TABLES `images` WRITE;
/*!40000 ALTER TABLE `images` DISABLE KEYS */;
INSERT INTO `images` VALUES (26,'a3746f6f-64f1-4305-8ea5-2435aa7577d0.png',6,0,1,'2024-11-12 19:09:16'),(27,'76abbd58-a766-44d8-a48a-c4021a6333e3.png',6,5,1,'2024-11-18 08:12:53'),(28,'d5412e25-36c4-4929-9a0a-c107d0740481.png',6,4,1,'2024-11-18 08:56:28'),(29,'0f51c2ab-78d8-4808-8019-e3af9372759c.png',6,11,1,'2024-11-18 20:23:33'),(30,'16e139d8-23f9-41a2-957b-f1ec96e299d3.png',6,11,1,'2024-11-18 20:23:33'),(31,'020132e3-2c1f-4f2a-919e-7f1c5c486cef.png',6,12,1,'2024-11-18 20:25:05'),(32,'55e94a5b-7bdb-41b6-9e94-b486a4f964ad.png',6,12,1,'2024-11-18 20:25:05'),(33,'33aeff0a-3603-4784-b253-36fa99d81d50.png',6,16,1,'2024-11-19 13:17:02'),(34,'62199bb5-6cc2-473c-9639-25dce5ea79b7.png',6,17,1,'2024-11-19 18:44:14'),(35,'3f428f54-096a-4b90-8b47-abe97a178ddc.png',6,17,1,'2024-11-19 18:44:14'),(36,'ed74fba2-6831-4f08-a29c-fe9c2274ad6f.png',6,18,1,'2024-11-19 19:32:55'),(37,'ee227802-6ea0-4e2f-955f-1f09f68fd537.png',6,18,1,'2024-11-19 19:32:55'),(38,'7ab08e08-3929-4e38-ae13-d7895bb4b4f0.png',6,4,1,'2024-11-30 16:53:16'),(39,'fb1e2222-a3ce-4985-b3aa-8850e2f46652.png',6,0,1,'2024-12-01 04:58:40'),(40,'78c22fac-cd4e-43fc-8887-a5d27e751393.png',6,0,1,'2024-12-01 05:18:40'),(41,'ece1705e-bc98-4626-9f81-a42a66bb78ab.png',6,0,1,'2024-12-01 05:21:13'),(42,'72b0c794-0ceb-40c4-95f4-47a312f4c293.png',6,0,1,'2024-12-01 05:23:12'),(43,'c365a886-58b4-46a8-9f0a-094088d4d8db.png',6,0,1,'2024-12-01 05:23:21'),(44,'4b23a9d8-c406-47a1-be58-fbbde04c8843.jpg',6,19,1,'2024-12-01 09:29:45'),(45,'f0efbd61-db08-4e38-86ac-7825f9804b90.jpg',6,19,1,'2024-12-01 09:29:45'),(46,'1a6a5665-aa19-4fc8-9039-3c060d0b4baf.jpg',6,19,1,'2024-12-01 09:29:45'),(47,'659ed22b-dcfd-4379-bdff-decae6f758b1.jpg',6,20,1,'2024-12-01 09:30:42'),(48,'362390ed-77db-43c8-bba7-72d611c6bc7f.jpg',6,20,1,'2024-12-01 09:30:42'),(49,'b72da60d-fd36-4036-902f-085c73e337f1.jpg',6,22,1,'2024-12-01 09:31:59'),(50,'f4b956f5-ebf4-4537-b215-ae2ec0fc1fe7.jpg',6,22,1,'2024-12-01 09:31:59'),(51,'e3bd06c7-15ad-4705-867c-bde6f3a5c56e.jpg',6,23,1,'2024-12-01 09:32:57'),(52,'a989373b-fc86-4d06-b680-07b0c9b9a124.jpg',6,23,1,'2024-12-01 09:32:57'),(53,'d01ae7a0-0e91-4071-8fcc-65d7d65c20b4.jpg',6,24,1,'2024-12-01 09:33:38'),(54,'9c478e1c-9083-42a0-acfa-bcefb7d4afe0.jpg',6,24,1,'2024-12-01 09:33:38'),(55,'002c85a7-cc6b-424b-a343-568c4b4146c9.jpg',6,25,1,'2024-12-01 09:34:56'),(56,'9a502cf5-7574-4472-80e5-21a722dbda14.jpg',6,25,1,'2024-12-01 09:34:56'),(57,'9d0a6c07-24a0-489a-babf-36f62be07edc.jpg',6,0,1,'2024-12-01 09:36:53'),(58,'ba0151c2-1a87-4ecc-92cb-0c7d37ca112d.jpg',6,0,1,'2024-12-01 13:40:25'),(59,'833b9c4d-a458-420a-841a-fdea89eedaf0.jpg',6,26,1,'2024-12-13 06:45:31'),(60,'64c0148e-b5ce-4c37-85ce-158aee54cf1d.jpg',6,27,1,'2024-12-13 07:01:42'),(61,'801a8427-f485-484f-a21d-930574a8f8e4.jpg',6,28,1,'2024-12-13 07:04:58'),(62,'d459d35d-7937-4bfe-910b-f16f37a758a6.jpg',6,29,1,'2024-12-13 07:07:20'),(63,'74872d70-47ae-4ee7-897b-92106432c439.jpg',6,30,1,'2024-12-13 07:08:59'),(64,'139d3c4f-e9d1-498d-ad8d-d8f1b418fb39.jpg',6,31,1,'2024-12-13 07:50:16'),(65,'2adf6434-5b74-4ddd-a2de-e415eba33581.jpg',6,32,1,'2024-12-13 07:56:55'),(66,'04a2d0c2-a72e-4ef7-b858-16427886b84c.jpeg',31,0,1,'2024-12-15 22:31:45'),(67,'c9dc8e25-7e1a-4580-8273-fe810ad04f92.jpg',6,33,1,'2024-12-16 00:34:24');
/*!40000 ALTER TABLE `images` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `listings`
--

DROP TABLE IF EXISTS `listings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `listings` (
  `listing_id` int NOT NULL AUTO_INCREMENT,
  `type` enum('Request','Offer') NOT NULL,
  `location` point NOT NULL,
  `user_id` int NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `date_created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `active` tinyint(1) NOT NULL DEFAULT '1',
  `city` varchar(255) NOT NULL,
  `country` varchar(255) NOT NULL,
  PRIMARY KEY (`listing_id`),
  SPATIAL KEY `location` (`location`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `listings_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `listings`
--

LOCK TABLES `listings` WRITE;
/*!40000 ALTER TABLE `listings` DISABLE KEYS */;
INSERT INTO `listings` VALUES (19,'Offer',_binary '\0\0\0\0\0\0\0˜?¯tÀI@\0\0Y4=·¿',6,'Potato Farmer','20$ per hour good job','2024-12-01 09:29:45',1,'',''),(20,'Request',_binary '\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0',6,'Manga Taza lal bei3333','batata','2024-12-01 09:30:42',1,'',''),(22,'Request',_binary '\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0',6,'Makdoos Baladi','fresh','2024-12-01 09:31:59',1,'',''),(23,'Offer',_binary '\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0',6,'Mjadra','','2024-12-01 09:32:57',1,'',''),(24,'Request',_binary '\0\0\0\0\0\0\0¥\Ã	K\á%A@\0\0N\ZöA@',6,'Maraba MeshMosh','','2024-12-01 09:33:38',1,'',''),(25,'Offer',_binary '\0\0\0\0\0\0\0µ\"K®\nA@\0§\ËL\ÔA@',6,'Shawarma ','Extra Toum','2024-12-01 09:34:55',1,'','Cyprus'),(26,'Offer',_binary '\0\0\0\0\0\0\0-¬A8A@\0§+\Ç\êA@',6,'Batoun','M3alem Nb1 Batoun','2024-12-13 06:45:31',1,'','Cyprus'),(27,'Offer',_binary '\0\0\0\0\0\0\0h\é›R<8A@€\ÓE<\ëA@',6,'King of  Ceramic','Best from italy nb1','2024-12-13 07:01:42',1,'','Cyprus'),(28,'Offer',_binary '\0\0\0\0\0\0\0‡§\Ìj+A@\0\0\0˜\ìA@',6,'Washing Machine Repair','Specialized in modern brands but can work on all types','2024-12-13 07:04:58',1,'',''),(29,'Request',_binary '\0\0\0\0\0\0\0F¸E|(A@\0\0\0\Ù\çA@',6,'House Cleaner','Get it squeeky clean','2024-12-13 07:07:20',1,'','Cyprus'),(30,'Request',_binary '\0\0\0\0\0\0\0„·¿¤2A@\0\0\0n\éA@',6,'Car Detailing','Ceramic','2024-12-13 07:08:59',1,'','Cyprus'),(31,'Request',_binary '\0\0\0\0\0\0\0\0\0\0-\îA@li0\Å6A@',6,'Blat','M3alem Blat','2024-12-13 07:50:16',1,'Mejdlaiya','Lebanon'),(32,'Request',_binary '\0\0\0\0\0\0\0\0\0\0˜\ìA@AW\é0A@',6,'ADS_Connectivity_Graph','2121','2024-12-13 07:56:55',1,'Dahr El Ain','Lebanon'),(33,'Offer',_binary '\0\0\0\0\0\0\0\0\0\à\ÙB@/Ø®ÿ˜A@',6,'Mountain','POtato','2024-12-16 00:34:24',1,'El Arz','Lebanon');
/*!40000 ALTER TABLE `listings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `transaction_id` int NOT NULL AUTO_INCREMENT,
  `user_offered_id` int NOT NULL,
  `user_offering_id` int NOT NULL,
  `listing_id` int NOT NULL,
  `price` double NOT NULL,
  `date_created` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `job_start_date` date NOT NULL,
  `job_end_date` date NOT NULL,
  `details_from_offered` varchar(1000) NOT NULL,
  `details_from_offering` varchar(1000) NOT NULL DEFAULT '',
  `currency` varchar(10) NOT NULL DEFAULT 'USD',
  `status` enum('Pending','Accepted','Completed') NOT NULL DEFAULT 'Pending',
  PRIMARY KEY (`transaction_id`),
  KEY `user_offered_id` (`user_offered_id`),
  KEY `user_offering_id` (`user_offering_id`),
  KEY `listing_id` (`listing_id`),
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`user_offered_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`user_offering_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `transactions_ibfk_3` FOREIGN KEY (`listing_id`) REFERENCES `listings` (`listing_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES (4,31,6,32,21,'2024-12-15 15:42:26','2024-12-09','2024-12-27','Price is per hour','','USD','Completed'),(5,31,6,32,21,'2024-12-15 15:42:45','2024-12-09','2024-12-27','Price is per hour','','USD','Accepted'),(6,31,6,32,21,'2024-12-15 15:43:10','2024-12-09','2024-12-27','Price is per hour','','USD','Completed'),(7,31,6,32,2111,'2024-12-15 15:43:35','2024-12-06','2024-12-25','Matata','','EUR','Completed'),(8,31,6,32,2111,'2024-12-15 15:44:40','2024-12-06','2024-12-25','Matata','','EUR','Pending'),(9,31,6,32,2111,'2024-12-15 15:44:59','2024-12-06','2024-12-25','Matata','','EUR','Pending'),(10,31,6,32,2111,'2024-12-15 15:47:36','2024-12-06','2024-12-25','Matata','','EUR','Pending'),(11,31,6,32,2111,'2024-12-15 17:39:41','2024-12-09','2024-12-27','ssss','','USD','Pending');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `first_name` varchar(50) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `phone_number` varchar(20) NOT NULL,
  `date_of_birth` date NOT NULL,
  `profession` varchar(100) NOT NULL,
  `location` point NOT NULL,
  `password` varchar(255) NOT NULL,
  `city` varchar(25) NOT NULL,
  `country` varchar(25) NOT NULL,
  `profile_image` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_id_UNIQUE` (`user_id`),
  SPATIAL KEY `location` (`location`),
  KEY `profile_image_idx` (`profile_image`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (2,'John','Doe','+1234567890','1990-05-15','Software Engineer',_binary '\0\0\0\0\0\0\0«\Ë)1ŸK@qZð¢¯A@','$2a$10$yk0iyuXu/BreqUpTtZI26OO0rsdepXXrW2xP.0xGpLug2y8KoXmFy','Ø¯Ù‡Ø³ØªØ§Ù† Ø¨ÛŒØ§Ø¨Ø§Ù†Ú©','Iran',16),(4,'John','Doe','1234567890','1990-05-15','Software Engineer',_binary '\0\0\0\0\0\0\0«\Ë)1ŸK@qZð¢¯A@','$2a$10$uajyaEVQArQHHdkXCeCuYO/6LjUMVc1p7R9NTx28v25exAGGAJLXy','Ø¯Ù‡Ø³ØªØ§Ù† Ø¨ÛŒØ§Ø¨Ø§Ù†Ú©','Iran',0),(6,'John','Doe','11234567890','1990-05-15','Software Engineer L1',_binary '\0\0\0\0\0\0\0\ß\ã\Û\æ=.A@\0\0\0|\êA@','$2a$10$g4W31r/KgVfTx2hf2gqe0eW74q4cdC7Td08bI4AzOEM1MyW4.e03a','','Cyprus',58),(14,'Adam','Elhassan','9613601360','2003-08-01','Software Engineer',_binary '\0\0\0\0\0\0\0\0\0\\¶\êA@}†Ž´Z.A@','$2a$10$g0QWBoh3BSGr3Os4Gxauu.3p6sJB0owQk1Zacz.ff35c64LLVo/dG','Nakhleh','Lebanon',0),(31,'Adam','Elhassan','03601360','2024-12-03','Software Engineer1',_binary '\0\0\0\0\0\0\0}\çyB¾A@\"@ñ@@','$2a$10$8MmBO4kvbbaQdJL589jhyeN4fAk25RVrP6bGNpgYMISJe.q0zs95K','Dar Al Fatwa','Lebanon',66),(32,'Maji','Wajdi','1123456789012121','2024-12-03','Batata',_binary '\0\0\0\0\0\0\0J¼¥‰ò@@hß±\Úu½A@','$2a$10$rMC/zB9K3H9VSv/SkSld2umR2AcZytHhWHNyq3SCk.S1yJt30tv3u','','Cyprus',0),(33,'Majdi','Wajdi','11234567890111','2024-12-01','Batata',_binary '\0\0\0\0\0\0\0O7­y{½A@õ…¶\ä\êñ@@','$2a$10$qJlfWq1Ikq4adQp7.p9rBO3VhhvYqJpOUD5TvfoV6ZE4vdOKll0v2','Beirut','Lebanon',0);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-12-16  7:11:23
