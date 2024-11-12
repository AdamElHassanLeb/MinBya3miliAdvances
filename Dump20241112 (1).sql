-- MySQL dump 10.13  Distrib 8.0.36, for Linux (x86_64)
--
-- Host: localhost    Database: minbya3mili
-- ------------------------------------------------------
-- Server version	8.0.39-0ubuntu0.24.04.2

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
  `listing_id` int DEFAULT NULL,
  `show_on_profile` tinyint(1) DEFAULT '0',
  `date_created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`image_id`),
  UNIQUE KEY `url_UNIQUE` (`url`),
  KEY `user_id` (`user_id`),
  KEY `listing_id` (`listing_id`),
  CONSTRAINT `images_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
  CONSTRAINT `images_ibfk_2` FOREIGN KEY (`listing_id`) REFERENCES `listings` (`listing_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `images`
--

LOCK TABLES `images` WRITE;
/*!40000 ALTER TABLE `images` DISABLE KEYS */;
INSERT INTO `images` VALUES (2,'2a5784e2-5070-4dce-8004-4503221f10c5.jpeg',1,1,1,'2024-11-12 01:25:48'),(3,'382b2f5b-f25b-4514-a85a-27b6cca40639.jpeg',1,2,1,'2024-11-12 02:06:07'),(4,'97fc408e-2920-4469-8fe8-e84f3ce1b50f.jpeg',1,1,1,'2024-11-12 02:06:26');
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `listings`
--

LOCK TABLES `listings` WRITE;
/*!40000 ALTER TABLE `listings` DISABLE KEYS */;
INSERT INTO `listings` VALUES (1,'Offer',_binary '\0\0\0\0\0\0\0ªñ\ÒMb€R@^K\È=[D@',1,'Sample Listing Title','This is a sample description for the test listing.','2024-11-11 01:33:05',1,'Osh Region','Kyrgyzstan'),(2,'Request',_binary '\0\0\0\0\0\0\0ªñ\ÒMbÀR@^K\È=[D@',1,'Sample Listing Title','This is a sample description for the test listing.','2024-11-11 01:33:28',1,'Naryn Region','Kyrgyzstan');
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
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_id_UNIQUE` (`user_id`),
  UNIQUE KEY `phone_number_UNIQUE` (`phone_number`),
  SPATIAL KEY `location` (`location`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Doe','John','+123456789021','1990-01-15','Software Engineer',_binary '\0\0\0\0\0\0\0\Ð\ÕV\ì/\ãB@¡ø1æ®µF@','$2a$10$AMw4iQgt3NMRJ65lDPCoSeD4T8MuQf5rptlTDKeHs7X78QSs2SDfq','Krasnodar Krai','Russia');
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

-- Dump completed on 2024-11-12  4:42:34
