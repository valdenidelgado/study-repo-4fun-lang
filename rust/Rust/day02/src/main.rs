const PI: f64 = 3.1415; // Constantes são declaradas com a palavra reservada const e seu valor não pode ser alterado.

fn main() {
    let name = "Valdeni";
    let mut lastName = "Alves"; // mut significa que a variável é mutável, ou seja, seu valor pode ser alterado.
    println!("Hello, world! {}", name);
    let name = "Outro Nome";
    lastName = "Delgado"; // Para alterar o valor de uma variável mutável, basta atribuir um novo valor a ela.
    println!("Hello, world! {}", name);
    println!("Hello, world! {}", lastName);
    println!("PI: {}", PI);
}