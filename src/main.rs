use std::env;
use std::str;
use std::process::Command;


fn main() {
    let input: Vec<String> = env::args().collect();
    if input.len() == 2 {
        match &input[1][..] {
            "list" => list(),
            "change" => change(),
            _ => println!(),
        }
    }
}


fn list() {
    // execute
    let output = Command::new("bash")
        .arg("-c")
        .arg("tr ':' '\n' <<< $PATH")
        .output()
        .expect("execution error");

    // error
    let output = match str::from_utf8(&output.stdout) {
        Ok(output) => output,
        Err(error) => panic!("{}", error),
    };

    // output
    let folder: Vec<&str> = output.split("\n").collect();
    for path in folder {
        println!("{}", path);
    }
}


fn change() {
    println!("change");
}

