Create database librerias;
use librerias;
create table libro(
Idlibro int not null auto_increment,
Nombre varchar(255) not null,
Autor varchar(255) not null,
Editorial varchar(255) not null,
primary key (Idlibro)
);