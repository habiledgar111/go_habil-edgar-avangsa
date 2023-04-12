create database secsion16;
use secsion16;

create table product_descriptions (
	id int not null primary key auto_increment, 
    description text, 
    create_at timestamp default current_timestamp, 
    update_at timestamp default current_timestamp on update current_timestamp
	);
create table product_types(
	id int not null primary key auto_increment,
    name varchar(255), 
    create_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp on update current_timestamp);
create table operators(
	id int not null primary key auto_increment,
    name varchar(255), 
    create_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp on update current_timestamp);
create table products(
	id int not null primary key auto_increment, 
    product_type_id int, 
    operator_id int, 
    foreign key(product_type_id) references product_types(id), 
    foreign key(operator_id)references operators(id),
    code varchar(50), 
    name varchar(100), 
    status smallint, 
    create_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp on update current_timestamp);
create table payment_methods(
	id int not null primary key auto_increment, 
    name varchar(255),
    status smallint, 
    create_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp on update current_timestamp);
create table users(
	id int not null primary key auto_increment, 
    status smallint, 
    dob date, 
    gender char(1), 
    create_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp on update current_timestamp);
create table transactions(
	id int not null primary key auto_increment,
    user_id int, 
    payment_method_id int, 
    foreign key(user_id) references users(id), 
    foreign key(payment_method_id) references payment_methods(id), 
    status varchar(10),
    total_qty int, 
    total_price numeric (25,2), 
    create_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp on update current_timestamp);
create table transaction_details(
	transaction_id int, 
    product_id int, 
    foreign key(transaction_id) references transactions(id), 
    foreign key(product_id) references products(id), 
    status varchar(10),
    qty int,
    price numeric (25,2),
    create_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp on update current_timestamp);
    
    show tables;
    
    insert into operators (name) values ("operator 1"), ("operator 2"),("operator 3"),("operator 4"),("operator 5");
    select * from operators;
    
    insert into product_types (name) values ("type 1"),("type 2"), ("type 3");
    select * from product_types;
    
    insert into products (product_type_id,operator_id,code, name, status) values (1,3 ,"code11","barang 1",10),(1,3 ,"code22","barang 2",10),(1,3 ,"code33","barang 3",10);
    select * from products;
    
    -- Insert 3 product dengan product type id = 2, dan operators id = 1.
    insert into products (product_type_id,operator_id,code, name, status) values (2,3 ,"code44","barang 4",10),(2,3 ,"code55","barang 5",10),(2,3 ,"code66","barang 6",10);
    select * from products;
    
    -- Insert 3 product dengan product type id = 3, dan operators id = 4.
	insert into products (product_type_id,operator_id,code, name, status) values (3,4 ,"code77","barang 7",10),(3,4 ,"code88","barang 8",10),(3,4 ,"code99","barang 9",10);
    select * from products;
    
    -- Insert 3 product dengan product type id = 3, dan operators id = 4. 
    alter table products add column product_description_id int;
	ALTER TABLE products ADD CONSTRAINT product_description_id foreign key (product_description_id) references product_descriptions(id);
    show columns from products;
    insert into product_descriptions (description) values ("ini desc bara 1"), ("ini desc barang 2");
    update products set product_description_id = 1;
    select * from products;
    
    -- Insert 3 payment methods.
    insert into payment_methods (name, status) values ("payment 1",1),("payment 2",1),("payment 3",1);
    select * from payment_methods;
    
    -- Insert 5 user pada tabel user.
    insert into users (status,gender)values (1,"L"),(1,"P"),(1,"L"),(1,"P"),(1,"L");
    select * from users;
    
    -- Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j) dan Insert 3 product di masing-masing transaksi.
    insert into transactions (user_id,payment_method_id,status,total_qty,total_price) values (1,1,"berhasil",5,10000),(2,2,"belum",5,10000),(3,3,"berhasil",5,10000);
    select * from transactions;
    insert into transaction_details(transaction_id,product_id, status, qty,price) values (1,1,"berhasil",5,10000),(2,2,"belum",5,10000),(3,3,"berhasil",5,10000);
    select * from transaction_details;
    
    
    select * from users where gender = "L";
    
    -- Tampilkan product dengan id = 3.
    select * from products where id = 3;
    
    -- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
    -- namun pada contoh tidak ada column nama pada table users
    alter table users add column name varchar(10);
    update users set name = "habil";
    update users set name = "no nme" where id in (4,5);
	select * from users;
    select * from users where create_at >=  current_date()-7 and name like '%a%';
    
    -- Hitung jumlah user / pelanggan dengan status gender Perempuan.
    select count(id) from users where gender = "P";
    
    -- Tampilkan data pelanggan dengan urutan sesuai nama abjad
    -- agar telihat dapat diurutkan saya mengubah satu data 
    update users set name = "abdi" where id = 3;
    select * from users order by name asc;
    
    -- Tampilkan 5 data pada data product
    select * from products limit 5;
    
    -- Ubah data product id 1 dengan nama ‘product dummy’.
    update products set name = "product dummy" where id = 1;
    select * from products where id = 1;
    
    -- Update qty = 3 pada transaction detail dengan product id = 1.
    update transaction_details set qty = 3 where product_id = 1;
    select * from transaction_details where product_id = 1;
    
    -- Delete data pada tabel product dengan id = 1
    -- karena id = 1 digunakan pada table transaction detail maka saya akan menghapus data pada transaction detail
    delete from transaction_details where product_id = 1;
    delete from products where id = 1;
    select * from products;
    
    -- Delete pada pada tabel product dengan product type id 1.
    -- karena data dengan product type = 1 digunakan pada table lain maka akan menghapus data dengan product type = 3
    select * from product_types;
    update products set product_type_id = 3 where id in (7,8,9);
    delete from products where product_type_id = 3;
    select * from products;
    
    -- Gabungkan data transaksi dari user id 1 dan user id 2.
    select * from transactions inner join users on transactions.id = users.id where users.id in (1,2);
    
    -- Tampilkan jumlah harga transaksi user id 1.
    select sum(total_price) from transactions where id = 1;
    select * from transactions;
    
    -- Tampilkan total transaksi dengan product type 2.
    insert into transactions (user_id,payment_method_id, status, total_qty,total_price) value (1,1,"berhasil", 2,2000);
    insert into transaction_details (transaction_id, product_id ,status,qty,price) values (4,4,"berhasil",2,2000);
    select * from transaction_details;
    select * from products;
    select * from transactions;
    select count(td.product_id) as total_transaction from transaction_details td, products p  where p.product_type_id = 2 and td.product_id = p.id;
    
    -- Tampilkan semua field table product dan field name table product type yang saling berhubungan.
    select p.*, pt.* from products p , product_types pt where p.product_type_id = pt.id; -- tanpa menggunakan inner join
    select * from products inner join product_types on products.product_type_id = product_types.id; -- dengan menggunakan inner join
    select * from products;
    select * from product_types;
    
    -- Tampilkan semua field table transaction, field name table product dan field name table user.
    -- select * from (select status,total_qty, total_price from transactions union all select code, name, status from products union all select name, gender from users) as jawaban;
    select transactions.*, products.*, users.* from transactions, products, users;
    select * from transactions;
    
    -- Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
    delimiter $$
    create function delete_transaction ( t_id int) 
    returns int
    reads sql data 
    DETERMINISTIC
    begin 
		delete from transaction_details where transaction_id = t_id;
        delete from transactions where id = t_id;
        return t_id; 
	end; $$
    
    delimiter ;
    
    drop function delete_transaction;
    
    select * from transactions;
    select delete_transaction(1);
    
    -- Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
    delimiter $$
    create function delete_transaction_detail ( t_id int) 
    returns int
    reads sql data 
    DETERMINISTIC
    begin 
		declare temp int;
        declare temp2 int;
		-- t_id = transaction id
        set temp = (select sum(qty) from transaction_details where transaction_id = t_id);
        set temp2 =(select total_qty from transactions where id = t_id);
        delete from transaction_details where transaction_id = t_id; 
        update transactions set total_qty = (temp+temp2) where id = t_id;
        return t_id; 
	end; $$
    
    delimiter ;
    
    select * from transactions;
    select delete_transaction_detail(3);
    
    -- Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
    select * from products; 
    select * from transaction_details;
    
    -- menggunakan left join 
    select * from products left join transaction_details on products.id = transaction_details.product_id where transaction_details.product_id is null;
    -- menggunakan sub query
    select * from products where id not in (select p.id from products p, transaction_details td where p.id = td.product_id); 