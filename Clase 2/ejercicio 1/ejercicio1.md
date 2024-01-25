## Segunda Parte

Se propone realizar las siguientes consultas a la base de datos movies_db.sql trabajada en la primera clase. Importar el archivo movies_db.sql desde PHPMyAdmin o MySQL Workbench y resolver las siguientes consultas:

1. Mostrar el título y el nombre del género de todas las series.
```sql
    SELECT s.title, g.name
      FROM series s 
INNER JOIN genres g ON g.id = s.genre_id 
```

 2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.

```sql
    SELECT s2.title, s.title, e.title, a.first_name, a.last_name 
      FROM episodes e
INNER JOIN seasons s ON s.id  = e.id 
INNER JOIN series s2 ON s2.id = s.serie_id
INNER JOIN actor_episode ae ON ae.episode_id = e.id 
INNER JOIN actors a ON a.id = ae.actor_id 
```

3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.

```sql
    SELECT s2.title, s.title, count(*) as episode_count
      FROM seasons s 
INNER JOIN series s2 ON s.serie_id = s2.id 
INNER JOIN episodes e ON e.season_id = s.id 
  GROUP BY s.id 
```

4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.

```sql
    SELECT g.name, count(*) movie_count
      FROM genres g 
INNER JOIN movies m ON m.genre_id = g.id 
  GROUP BY g.id 
```

5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
```sql
SELECT DISTINCT a.first_name, a.last_name
           FROM actors a 
     INNER JOIN actor_movie am ON am.actor_id = a.id 
     INNER JOIN movies m ON m.id = am.movie_id  
          WHERE m.title LIKE "%La guerra de las galaxias%"
```


## Ejercicios de la clase sincronica

1. Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
```sql
    SELECT e.nombre, e.puesto, d.localidad
      FROM empleados e
INNER JOIN departamento d ON d.depto_nro = e.depto_nro
```
2. Visualizar los departamentos con más de cinco empleados.
```sql
    SELECT d.depto_nro, d.depto_nombre, d.localidad
      FROM departamento d
INNER JOIN empleado e ON e.depto_nro = d.depto_nro
  GROUP BY d.depto_nro
    HAVING count(*) > 5
```
3. Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
```sql
    SELECT e.nombre, e.salario, d.depto_nombre
      FROM empleado e
INNER JOIN departamento d on d.depto_nro = e.depto_nro
     WHERE e.puesto = "Presidente"
```
4. Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
```sql
    SELECT e.cod_emp, e.nombre, e.apellido, e.puesto, e.fecha_alta, e.salario, e.comision
      FROM empleado e
INNER JOIN departamento d ON d.depto_nro = e.depto_nro
     WHERE d.nombre_dpto = "Contabilidad"
  ORDER BY e.nombre DESC
```
5. Mostrar el nombre del empleado que tiene el salario más bajo.
```sql
  SELECT e.nombre
    FROM empleado e
ORDER BY e.salario ASC 
   LIMIT 1
```
6. Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
```sql
    SELECT e.nombre
      FROM empleado e
INNER JOIN departamento d ON d.depto_nro = e.depto_nro
     WHERE d.nombre_dpto = "Ventas"
  ORDER BY e.salario ASC 
     LIMIT 1
```

## Ejercicios de la practica grupal

### Ejercicio 1

Se tiene el siguiente DER que corresponde al esquema que presenta la base de datos de una “biblioteca”.

![img](https://i.ibb.co/RY6YBNH/Screenshot-2024-01-25-at-10-15-49.png)

En base al mismo, plantear las consultas SQL para resolver los siguientes requerimientos:


1. Listar los datos de los autores.
```sql
SELECT a.nombre, a.nacionalidad
  FROM autor a 
```
2. Listar nombre y edad de los estudiantes
```sql
SELECT e.nombre, e.apellido, e.edad
  FROM estudiante e
```
3. ¿Qué estudiantes pertenecen a la carrera informática?
```sql
SELECT e.nombre, e.apellido 
  FROM estudiante e
 WHERE e.Carrera = "Informatica" 
```
4. ¿Qué autores son de nacionalidad francesa o italiana?
```sql
SELECT a.nombre, a.nacionalidad
  FROM autor a
 WHERE a.nacionalidad = "Francesa" OR
       a.nacionalidad = "Italiana"
```
5. ¿Qué libros no son del área de internet?
```sql
SELECT l.titulo, l.area
  FROM libro l
 WHERE l.area != "Internet"
```
6. Listar los libros de la editorial Salamandra.
```sql
SELECT l.titulo, l.area
  FROM libro l
 WHERE l.editorial = "Salamandra"
```
7. Listar los datos de los estudiantes cuya edad es mayor al promedio.
```sql
SELECT e.nombre, e.apellido, e.direccion, e.carrera, e.edad
  FROM estudiante e
 WHERE e.edad > (SELECT AVG(edad) FROM estudiante) 
```
8. Listar los nombres de los estudiantes cuyo apellido comience con la letra G.
```sql
SELECT e.nombre, e.apellido
  FROM estudiante e
 WHERE e.apellido like "G%"
```
9. Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).
```sql
    SELECT a.nombre
      FROM autor a
INNER JOIN libro l ON l.id_autor = a.id_autor
     WHERE l.titulo = "El Universo: Guía de viaje" 
```
10. ¿Qué libros se prestaron al lector “Filippo Galli”?
```sql
    SELECT l.titulo, l.editorial, l.area
      FROM libro l
INNER JOIN prestamo p ON p.id_libro = l.id_libro
INNER JOIN estudiante e ON e.id_lector = p.id_lector
     WHERE e.nombre = "Filippo" AND
           e.apellido = "Galli" 
```
11. Listar el nombre del estudiante de menor edad.
```sql
  SELECT e.nombre, e.apellido
    FROM estudiante e
ORDER BY e.edad ASC
   LIMIT 1
```
12. Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.
```sql
    SELECT e.nombre, e.apellido
      FROM estudiante e
INNER JOIN prestamo p ON p.id_lector = e.id_lector
INNER JOIN libro l ON l.id_libro = p.id_libro
     WHERE l.titulo like "%Bases_de_Datos%"
```
13. Listar los libros que pertenecen a la autora J.K. Rowling.
```sql
    SELECT l.titulo, l.editorial, l.area
      FROM libro l
INNER JOIN libroautor la ON la.id_libro = l.id_libro
INNER JOIN autor a ON a.id_autor = la.id_autor
     WHERE a.nombre = "J.K. Rowling"
```
14. Listar títulos de los libros que debían devolverse el 16/07/2021.
```sql
    SELECT l.titulo, l.editorial, l.area
      FROM libro l
INNER JOIN prestamo p ON p.id_libro = l.id_libro
     WHERE p.fecha_devolucion < "2021-07-16"
```

### Ejercicio 2

Implementar la base de datos en PHPMyAdmin o MySQL Workbench, cargar cinco registros en cada tabla y probar algunas consultas planteadas en el Ejercicio 1. 

```sql
   CREATE DATABASE IF NOT EXISTS prestamos_libros;

    USE prestamos_libros;   -- Seleccionamos la base de datos nueva

    /* Definicion del esquema de la base de datos */
    CREATE TABLE Autor(
        id_autor        int NOT NULL AUTO_INCREMENT,
        nombre          varchar(60),
        nacionalidad    varchar(30),
        PRIMARY KEY (id_autor)
    );

    CREATE TABLE Libro(
        id_libro        int NOT NULL AUTO_INCREMENT,
        titulo          varchar(60),
        ediotrial       varchar(30),
        area            varchar(30),
        PRIMARY KEY (id_libro) 
    );

    CREATE TABLE LibroAutor(
        ID              int NOT NULL AUTO_INCREMENT,
        id_autor        int NOT NULL,
        id_libro        int NOT NULL,
        PRIMARY KEY (ID),
        FOREIGN KEY (id_autor) REFERENCES Autor(id_autor),
        FOREIGN KEY (id_libro) REFERENCES Libro(id_libro)
    );

    CREATE TABLE Estudiante(
        id_estudiante   int NOT NULL AUTO_INCREMENT,
        nombre          varchar(30),
        apellido        varchar(30),
        direccion       varchar(60),
        carrera         varchar(30),
        edad            smallint,
        PRIMARY KEY (id_estudiante)
    );

    CREATE TABLE Prestamo(
        ID               int NOT NULL AUTO_INCREMENT,
        id_estudiante    int NOT NULL,
        id_libro         int NOT NULL,
        fecha_prestamo   date NOT NULL,
        fecha_devolucion date NOT NULL,
        CHECK (fecha_prestamo < fecha_devolucion),
        PRIMARY KEY (ID),
        FOREIGN KEY (id_estudiante) REFERENCES Estudiante(id_estudiante),
        FOREIGN KEY (id_libro) REFERENCES Libro(id_libro) 
    );
```