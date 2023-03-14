create database alta_online_shop;

show databases;

use alta_online_shop;

create table products (
	ProductID INT primary key not null AUTO_INCREMENT, 
    Product_type varchar(50),
    Product_description text,
    Operator varchar(10), 
    Payment_methods varchar(15)
);

create table users (
	UserID INT primary key not null auto_increment,
    Nama varchar (15), 
    Alamat varchar (20), 
    Tanggal_lahir date,
    Status_user varchar (1), 
    Gender varchar (1), 
    Create_at timestamp, 
    Update_at timestamp
    );
    
-- tidak membuat table transaksi detail karena transakasi detail berhubungan one to one pada transaksi maka detail dapat diamasukan pada table transaksi
create table transaksi (
	TransaksiID INT primary key not null auto_increment,
    ProductID_FK INT not null,
    UserID_FK int not null, 
    Jumlah int, 
    Total double, 
    Create_at timestamp,
    foreign key (ProductID_FK) references products (ProductID), 
    foreign key (UserID_FK) references users (UserID)
    );
    
create table kurirs (
	KurirID int not null primary key auto_increment,
    Nama varchar (15),
    Create_at timestamp, 
    Update_at timestamp
    );
    
alter table kurirs add Ongkos_dasar double;
alter table kurirs rename to shipping;

drop table shipping;

create table user_info (
	User_infoID int primary key not null,
	UserID_FK int not null,
    nomor int,
    email varchar (15),
    sosmed varchar (15),
    foreign key (UserID_FK) references users (UserID)
    );
    
create table pinjaman (
	PinjamanID int not null primary key,
    UserID_FK int not null,
    Total_uang double,
    Waktu_pinjam timestamp, 
    Jatuh_tempo date,
    foreign key (userID_FK) references users(UserID)
    );

create table wishlists (
	WishlistID int primary key not null,
	ProductID_FK INT not null,
    UserID_FK int not null, 
    Jumlah int, 
    Create_at timestamp,
    Update_at timestamp,
    foreign key (ProductID_FK) references products (ProductID), 
    foreign key (UserID_FK) references users (UserID)
    );
    
show tables;	

show columns from pinjaman;