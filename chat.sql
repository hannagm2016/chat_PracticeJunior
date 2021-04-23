-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Апр 19 2021 г., 09:10
-- Версия сервера: 10.3.22-MariaDB
-- Версия PHP: 7.1.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `chat`
--

-- --------------------------------------------------------

--
-- Структура таблицы `contacts`
--

CREATE TABLE `contacts` (
  `id` int(50) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `phone` varchar(15) DEFAULT NULL,
  `status` varchar(10) DEFAULT NULL,
  `Password` varchar(100) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `contacts`
--

INSERT INTO `contacts` (`id`, `name`, `phone`, `status`, `Password`, `email`) VALUES
(1, 'Tom', '8-800-777-22-33', 'online', '', ''),
(2, 'Sam', '8-800-777-22-55', 'offline', '', ''),
(3, 'Greg', '8-800-777-77-66', 'online', '', ''),
(4, 'Paul', '8-800-777-22-99', 'offline', '', ''),
(6, 'Ann', '', '', '$2a$14$Gl/Od/3n.AxYDMAwT0CA9ef.zZ3.ZB5l5VndYAUuzbylg1kcHaleG', 'ann@c.com'),
(7, 'hannaboiko2016@gmail.com', '', '', NULL, 'hannaboiko2016@gmail.com'),
(8, 'Nik', '', '', '$2a$14$XrWxPnz3jdPZm9W4Y60ay.b1Ge4lMoicMz9i80I7bES/LPuv4zBza', 'nik@k.com'),
(9, 'Hanna Boiko', '', '', NULL, ''),
(10, 'hannaboiko2016@gmail.com', '', '', NULL, 'hannaboiko2016@gmail.com'),
(11, 'Ray', '', '', '$2a$14$VY1TVbI0U9N56Uqbl7DWQOE4FunjgJR1ifaSywym7ACCrazElnpke', 'ray@n.com'),
(12, 'Dan', '', '', '$2a$14$g/iGnMYrryRylxmssM7Wc.cZFP34BYJxB7V8H6J1NsC7Iq.uG7q3C', 'den@m.nn');

-- --------------------------------------------------------

--
-- Структура таблицы `messages`
--

CREATE TABLE `messages` (
  `id` int(50) NOT NULL,
  `user_from_id` int(20) DEFAULT NULL,
  `user_to_id` int(20) DEFAULT NULL,
  `time` varchar(25) NOT NULL DEFAULT 'current_timestamp()',
  `text` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `messages`
--

INSERT INTO `messages` (`id`, `user_from_id`, `user_to_id`, `time`, `text`) VALUES
(1, 3, 1, '19:10', 'Hello, I\'m Tom!'),
(2, 1, 3, '19:30', 'Hi thereother'),
(3, 2, 3, '19:15', 'Hello, Sam!'),
(4, 2, 3, '19:32', 'How are you?'),
(5, 3, 4, '20:00', 'La-la Lend'),
(6, 3, 2, '20:01', 'Chao'),
(7, 2, 4, '13:00', 'will not get this'),
(40, 3, 1, '12:42', 'nice to see you'),
(131, 3, 4, '20:07', 'lalalala'),
(135, 2, 4, '21:39', 'Hi Paul'),
(136, 2, 3, '21:39', 'How are you Greg?'),
(137, 2, 0, '21:48', 'Am I sam?'),
(138, 2, 0, '21:49', 'Now?'),
(139, 0, 2, '22:03', 'hjh'),
(140, 3, 4, '07:55', 'bla bla');

-- --------------------------------------------------------

--
-- Структура таблицы `relations`
--

CREATE TABLE `relations` (
  `id` int(20) NOT NULL,
  `relation` varchar(10) NOT NULL,
  `userId` int(10) NOT NULL,
  `userTo` int(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `relations`
--

INSERT INTO `relations` (`id`, `relation`, `userId`, `userTo`) VALUES
(1, 'Tom', 1, 0),
(2, 'Sam', 2, 0);

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `contacts`
--
ALTER TABLE `contacts`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `messages`
--
ALTER TABLE `messages`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `relations`
--
ALTER TABLE `relations`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `contacts`
--
ALTER TABLE `contacts`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT для таблицы `messages`
--
ALTER TABLE `messages`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=141;

--
-- AUTO_INCREMENT для таблицы `relations`
--
ALTER TABLE `relations`
  MODIFY `id` int(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
