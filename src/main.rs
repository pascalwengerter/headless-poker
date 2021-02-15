mod setup;

fn main() {
    let game = setup::run();
    let mut ongoing: bool = true;
    let mut rounds: u32 = 0;
    let mut current_cashpool: u32 = 0;

    let human_player = Player {
        name: game.human_player,
        stack: (game.full_pot_size as u32 / game.players as u32),
        in_game: true,
        in_round: true
    };
    // initialize computer player structs in loop
    // put all players in vector?

    let statistics = loop {
        if ongoing == false {
            let winner: &str = "Patrick"; // name of the one player left that is in_game
            break (winner, rounds);
        }
        current_cashpool = 0;


        // main game loop goes here
        println!("{} has {}$", human_player.name, human_player.stack);
        println!("Big blind: {}", game.big_blind_size);
        println!("Small blind: {}", game.small_blind_size);
    
        // if (only one player left that is in_game) {
            ongoing = false;
        // } else {
            rounds = rounds + 1;
        // }
    };
    announce_winner(statistics.0, statistics.1);
}

fn announce_winner(winning_player: &str, played_rounds: u32) {
    println!("---\n{} has won the game after {} rounds!\nGame finished.", winning_player, played_rounds);
}

struct Player{
    name: String,
    stack: u32,
    in_game: bool,
    in_round: bool
}
