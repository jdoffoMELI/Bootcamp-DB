DROP DATABASE IF EXISTS warehouses_test;

CREATE DATABASE warehouses_test;

USE warehouses_test;

CREATE TABLE `warehouses` (
  `id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `address` varchar(150) NOT NULL,
  `telephone` varchar(150) NOT NULL,
  `capacity` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Asignar la columna `id` 
ALTER TABLE `warehouses`
  ADD PRIMARY KEY (`id`);

-- Modificar tabla warehouses para que el id sea autoincrementable
ALTER TABLE `warehouses`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1;

INSERT INTO `warehouses` (`name`, `address`, `telephone`, `capacity`) VALUES
('Main Warehouse', '221 Baker Street', "4555666", 100);


INSERT INTO `warehouses` (`name`, `address`, `telephone`, `capacity`) VALUES
('Secondary Warehouse', '221 Baker Street', "45556222", 50);
