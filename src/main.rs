use std::io;
use std::fs;
use std::env;
use std::str;
use std::io::BufRead;
use std::fs::OpenOptions;
use std::io::{Write, BufReader};

const PATH: &str = "/home/kolu/.bashrc";

fn main() {
    let input: Vec<String> = env::args().collect();
    let args = input.len();

    match args {
        1 => list(),
        2 => match &input[1][..] {
            "add" => add(),
            _ => println!(),
        }
        3 => match &input[1][..] {
            "del" => del(&input[2]),
            _ => println!(),
        }
        _ => println!(),
    }

}


fn list() {
    // open
    let file = OpenOptions::new()
        .read(true)
        .open(PATH)
        .expect("error: open file");

    // read
    let reader = BufReader::new(file);
    
    // output
    println!();
    for (index, line) in reader.lines().enumerate()  {
        let line = line.expect("error: reading line");
        if line.contains("export PATH") {
            println!("  {}. {}", index+1, &line[19..line.len()-1]);
        }
    }
    println!();
}


fn add() {
    // input
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("error: read input");
    let new_path = format!(r#"export PATH="$PATH:{}""#, input.trim());

    // open
    let mut file = OpenOptions::new()
        .append(true)
        .open(PATH)
        .expect("error: open file");

    // write
    writeln!(file, "{}", new_path).expect("error: write file");
}


fn del(index_to_delete: &str) {
    // error
    let check = index_to_delete.chars().all(char::is_numeric);
    if !check {
        panic!("error: index must be a number");
    }
    let index_to_delete: usize = index_to_delete.parse().expect("error: index parsing");

    // open
    let file = OpenOptions::new()
        .read(true)
        .open(PATH)
        .expect("error: open file");
    
    // read
    let reader = BufReader::new(file);
    let mut indices: Vec<usize> = Vec::new();
    let mut contents: Vec<String> = Vec::new();

    for (index, line) in reader.lines().enumerate() {
        let line = line.expect("error: read line");
        if line.contains("export PATH") {
            indices.push(index+1);
        }
        contents.push(line);
    }

    // error
    if !indices.contains(&index_to_delete) {
        panic!("error: index out of range");
    }
    
    // create bashrc temp
    let mut temp_file = create_file(".bashrctemp", "error: create temp file"); 
    for line in &contents {
        writeln!(temp_file, "{}", line).expect("error: write temp file");
    }
    
    // remove, create and write to bashrc
    fs::remove_file(PATH).expect("error: remove bashrc file");
    let mut file = create_file(PATH, "error: create new bashrc"); 
    for (index, line) in contents.iter().enumerate() {
        if index+1 != index_to_delete {
            writeln!(file, "{}", line).expect("error: write file");
        }
    }

    // remove bashrc temp
    fs::remove_file(".bashrctemp").expect("error: remove temp file");
}


fn create_file(path: &str, error: &str) -> std::fs::File {
    OpenOptions::new()
        .write(true)
        .create(true)
        .open(path)
        .expect(error)
}
