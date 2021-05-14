struct Token {
    value: String,
    r#type: String,
}

impl std::fmt::Display for Token {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "{} {}", self.value, self.r#type)
    }
}

fn main() {
    fn tokenizer(input: &str) {
        let mut tokens: Vec<Token> = Vec::new();
        let mut current = 0;
        for char in input.chars() {
            match char.to_string().as_str() {
                "[" => tokens.push(Token {
                    value: char.to_string(),
                    r#type: "paren".to_string(),
                }),
                _ => tokens.push(Token {
                    value: char.to_string(),
                    r#type: "poop".to_string(),
                }),
            }
            println!("{}", tokens[current]);
            current = current + 1;
        }
    }

    tokenizer("[");
}
