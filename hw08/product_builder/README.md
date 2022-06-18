## HW08
###  Моё приложение представляет из себя простой REST API. API будет предоставлять точки для доступа к «продуктам» и управления ими.

- Создание нового продукта
- Обновление существующего продукта
- Удаление существующего продукта
- Получение существующего продукта
- Получение списка продуктов.

### API Specification

<p> Create a new product in response to a valid POST request at /product,
<p> Update a product in response to a valid PUT request at /product/{id},
<p> Delete a product in response to a valid DELETE request at /product/{id},
<p> Fetch a product in response to a valid GET request at /product/{id}, and
<p> Fetch a list of products in response to a valid GET request at /products.


#### Требуемые модули:
- mux – The Gorilla Mux router and,
- pq – The PostgreSQL driver.

### Структуры таблиц:
Таблица products содержит поля:
- id – the primary key in this table,
- name – the name of the product and,
- price – the price of the product.