# Praktikum System Design

## Problem 1 - Diagram

1. ERD Aplikasi Pencatatan Pengeluaran Keuangan
    ![alt text](../screenshots/ERD_AplikasiPencatatanPengeluaran.png "ERD_AplikasiPencatatanPengeluaran")
2. Use Case Diagram Aplikasi Pencatatan Pengeluaran Keuangan
    ![alt text](../screenshots/UseCaseDiagram_AplikasiPencatatanPengeluaran.png "UseCaseDiagram_AplikasiPencatatanPengeluaran")

## Problem 2 - Query

SQL
```sql
SELECT * FROM users;
```

Redis
```
KEYS *users*
```

Neo4J
```
SHOW DATABASE users YIELD *
```

Cassandra
```
SELECT * FROM system_schema.users;
```