**Ejercicio**: Una mueblería necesita la implementación de una base de datos para controlar las ventas que realiza por día, el stock de sus artículos (productos) y la lista de sus clientes que realizan las compras.

+ ¿Cuáles serían las entidades de este sistema?: Producto, cliente, comprobante y linea de comprobante
+ ¿Qué atributos se determinarán para cada entidad? (Considerar los que se crean necesarios): 
  + Productos: ID, nombre, precio, stock
  + Cliente: ID, nombre
  + Comprobante: n_comprobante, id_cliente, fecha, total
  + Linea comprobante: id, n_comprobante, id_producto, cantidad
+ ¿Cómo se conformarán las relaciones entre entidades? ¿Cuáles serían las cardinalidades?: Un producto puede aparecer en muchas lineas de comprobantes pero una linea de comprobante solo puede tener un producto; Un comprobante puede tener muchas lineas de comprobante pero una linea de comprobante solo puede tener un comprobante; Un cliente puede tener muchos comprobantes pero un comprobante solo le pertenece a un cliente.

```mermaid
erDiagram
    PRODUCTO{
        int ID PK
        string nombre
        float precio
        int stock
    }

    CLIENTE{
        int ID PK
        string nombre
    }

    LINEACOMPROBANTE{
        int ID PK
        int n_comprobante FK
        int id_producto FK
        int cantidad
    }

    COMPROBANTE{
        int n_comprobante PK
        int id_cliente FK
        date fecha
        float total
    }

    PRODUCTO ||--o{ LINEACOMPROBANTE : aparece
    CLIENTE ||--o{ COMPROBANTE: tiene
    COMPROBANTE ||--|{ LINEACOMPROBANTE: tiene

```

**Ejercicio**: Realizar un diagrama de entidad - relación para el sistema de una concesionaria, que desea gestionar los servicios de los coches de sus clientes. 

Para el módulo del sistema, se necesita almacenar información de los clientes, los coches que estos poseen y los service/revisiones de cada uno de esto

```mermaid
erDiagram
    COCHE{
        int ID PK
        int id_cliente FK
        int anio
        string marca
        string modelo
        string color
    }

    CLIENTE{
        int ID PK
        string nombre
    }

    REVISION{
        int ID PK
        int id_coche FK
        date fecha_revision
        string descripcion
    }

    CLIENTE ||--o{ COCHE: tiene
    COCHE ||--o{ REVISION: tiene

```
