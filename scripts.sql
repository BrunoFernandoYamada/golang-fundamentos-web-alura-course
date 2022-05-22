drop table if exists product;

create table product (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
)

insert into product (nome, descricao, preco, quantidade) values
('Camiseta', 'Preta', 19, 10 ),
('Fone', 'Muito bom',99, 5)
('PC Gamer', 'High End', 19000.59, 1)