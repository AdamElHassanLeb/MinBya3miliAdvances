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
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `images`
--

LOCK TABLES `images` WRITE;
/*!40000 ALTER TABLE `images` DISABLE KEYS */;
INSERT INTO `images` VALUES (26,'a3746f6f-64f1-4305-8ea5-2435aa7577d0.png',6,0,1,'2024-11-12 19:09:16'),(27,'76abbd58-a766-44d8-a48a-c4021a6333e3.png',6,5,1,'2024-11-18 08:12:53'),(28,'d5412e25-36c4-4929-9a0a-c107d0740481.png',6,4,1,'2024-11-18 08:56:28'),(29,'0f51c2ab-78d8-4808-8019-e3af9372759c.png',6,11,1,'2024-11-18 20:23:33'),(30,'16e139d8-23f9-41a2-957b-f1ec96e299d3.png',6,11,1,'2024-11-18 20:23:33'),(31,'020132e3-2c1f-4f2a-919e-7f1c5c486cef.png',6,12,1,'2024-11-18 20:25:05'),(32,'55e94a5b-7bdb-41b6-9e94-b486a4f964ad.png',6,12,1,'2024-11-18 20:25:05'),(33,'33aeff0a-3603-4784-b253-36fa99d81d50.png',6,16,1,'2024-11-19 13:17:02'),(34,'62199bb5-6cc2-473c-9639-25dce5ea79b7.png',6,17,1,'2024-11-19 18:44:14'),(35,'3f428f54-096a-4b90-8b47-abe97a178ddc.png',6,17,1,'2024-11-19 18:44:14'),(36,'ed74fba2-6831-4f08-a29c-fe9c2274ad6f.png',6,18,1,'2024-11-19 19:32:55'),(37,'ee227802-6ea0-4e2f-955f-1f09f68fd537.png',6,18,1,'2024-11-19 19:32:55'),(38,'7ab08e08-3929-4e38-ae13-d7895bb4b4f0.png',6,4,1,'2024-11-30 16:53:16');
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
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `listings`
--

LOCK TABLES `listings` WRITE;
/*!40000 ALTER TABLE `listings` DISABLE KEYS */;
INSERT INTO `listings` VALUES (4,'Offer',_binary '\0\0\0\0\0\0\0öB\Û{RÀ¶ä „]D@',6,'Professional Lawn Mowing Service','Offering reliable lawn mowing services with 5 years of experience.','2024-11-18 05:13:39',1,'New York','United States'),(11,'Offer',_binary '\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0',6,'2111','21111','2024-11-18 20:23:33',1,'',''),(12,'Request',_binary '\0\0\0\0\0\0\0ù?°\Éb.A@8ÿ+ \êA@',6,'ImageTest','222','2024-11-18 20:25:05',1,'','Cyprus'),(14,'Offer',_binary '\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0',6,'ImageTest','222','2024-11-19 10:44:51',1,'',''),(15,'Offer',_binary '\0\0\0\0\0\0\0«\Ë)1ŸK@qZð¢¯A@',6,'ImageTest','2223','2024-11-19 10:45:00',1,'Ø¯Ù‡Ø³ØªØ§Ù† Ø¨ÛŒØ§Ø¨Ø§Ù†Ú©','Iran');
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
  `price_with_currency` varchar(50) NOT NULL,
  `date_created` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `job_start_date` date NOT NULL,
  `job_end_date` date NOT NULL,
  `details_from_offered` text NOT NULL,
  `details_offering` text NOT NULL,
  PRIMARY KEY (`transaction_id`),
  KEY `user_offered_id` (`user_offered_id`),
  KEY `user_offering_id` (`user_offering_id`),
  KEY `listing_id` (`listing_id`),
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`user_offered_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`user_offering_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `transactions_ibfk_3` FOREIGN KEY (`listing_id`) REFERENCES `listings` (`listing_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
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
  UNIQUE KEY `phone_number_UNIQUE` (`phone_number`),
  SPATIAL KEY `location` (`location`),
  KEY `profile_image_idx` (`profile_image`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (2,'John','Doe','+1234567890','1990-05-15','Software Engineer',_binary '\0\0\0\0\0\0\0«\Ë)1ŸK@qZð¢¯A@','$2a$10$yk0iyuXu/BreqUpTtZI26OO0rsdepXXrW2xP.0xGpLug2y8KoXmFy','Ø¯Ù‡Ø³ØªØ§Ù† Ø¨ÛŒØ§Ø¨Ø§Ù†Ú©','Iran',16),(4,'John','Doe','1234567890','1990-05-15','Software Engineer',_binary '\0\0\0\0\0\0\0«\Ë)1ŸK@qZð¢¯A@','$2a$10$uajyaEVQArQHHdkXCeCuYO/6LjUMVc1p7R9NTx28v25exAGGAJLXy','Ø¯Ù‡Ø³ØªØ§Ù† Ø¨ÛŒØ§Ø¨Ø§Ù†Ú©','Iran',0),(6,'John','Doe','11234567890','1990-05-15','Software Engineer L1',_binary '\0\0\0\0\0\0\0¥\×\Ú\ëO.A@²ývž\êA@','$2a$10$5qGZEiE0nolEnb0/qQ1/deUN1bdolB6vMYtSdMDvpJbTi4TtfmZJ.','','Cyprus',26),(12,'Adam','Elhassan','03601360','2003-08-01','SWE',_binary '\0\0\0\0\0\0\0\0$\0Î§\êA@\á_MPf.A@','$2a$10$Y50XCNXr2Bp5Jz0Cf0l2fe9ZopxlNEBzQNYTdALubM4szFXTtOgwu','Nakhleh','Lebanon',0);
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

-- Dump completed on 2024-11-30 19:24:46
