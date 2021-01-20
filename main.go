package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "log"
)

type Player struct {
    name string
    stake int
    // active bool // reasonable? *question*
}

// type Card struct {
//     value string // "2" to "10", "J", "Q", "K", "A"
//     type string // "Heart", "Ace", "Pike", "Square" // ♥️♦️♣️♠️
//     color string // "red" or "black"
// }

func start_game() {
    fmt.Printf("Ready to start a new game? (y/n) \n")
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSuffix(text, "\n")

    if (text == "N" || text == "n") {
        os.Exit(0)
    } else if (text == "y" || text == "Y") {
        fmt.Println("--------------------------------")
        fmt.Println("Command Line Poker Game started")
        fmt.Println("--------------------------------")
    } else {
        start_game()
    }
}

func set_player_name(requested string, default_value string) string {
    fmt.Printf("Enter %v or hit enter for default value (%v): \n", requested, default_value)
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSuffix(text, "\n")

    if (len(text) < 2) {
        return default_value
    } else {
        return text
    }
}

func set_player_count(requested string, default_value int) int {
    var player_amount int
    fmt.Printf("Enter %v or hit enter for default value (%v): \n", requested, default_value)

    if    _, err := fmt.Scan(&player_amount);    err != nil {
        log.Print("  Scan for i failed, due to ", err)
    }
    if (player_amount < 0 || player_amount > 9) {
        set_player_count("amount of computer players (between 1 and 9)", 7)
    }
    return player_amount
}

func set_pot_size(requested string, default_value int) int {
    var chip_amount int
    fmt.Printf("Enter %v or hit enter for default value (%v €): \n", requested, default_value)

    if    _, err := fmt.Scan(&chip_amount);    err != nil {
        log.Print("  Scan for i failed, due to ", err)
    }
    if (chip_amount < 0 || chip_amount > 1000000) {
        set_pot_size("Pot size", 1000)
    }
    return chip_amount
}

func set_big_blind_size(requested string, chip_amount int) int {
    var big_blind_choice int
    var default_value = (chip_amount / 25)
    fmt.Printf("Enter %v or hit enter for default value (%v €): \n", requested, default_value)

    if    _, err := fmt.Scan(&big_blind_choice);    err != nil {
        log.Print("  Scan for i failed, due to ", err)
    }
    if (big_blind_choice < 0 || big_blind_choice > 1000000) {
        set_pot_size("Pot size", 1000)
    }
    return big_blind_choice
}


func initialize_game () {
    // INITIALIZATION
    /* 
    [X] (y/n) to start game
    [X] welcome message
    [X] insert name (or use default value)
    [X] insert amount of players (or use default value, minimum 1 maximum 9)
    [X] insert pot_size (or use default value)
    [X] insert big_blind_size (or use default value calculated from pot_size)
    [X] calculate small_blind_size from big_blind_size
    [X] calculate player names (check that they're not equal to player_name)
    [X] calculate each player's stake
    */

    var player_name string
    var player_count int
    var pot_size int
    var big_blind_size int
    var small_blind_size int

    start_game()
    player_name = set_player_name("name", "Player 1")
    player_count = set_player_count("amount of computer players (between 1 and 9)", 7)
    pot_size = set_pot_size("Pot size", 1000)
    big_blind_size = set_big_blind_size("Big blind", pot_size)
    small_blind_size = (big_blind_size / 2)

    human_player := Player{name: player_name, stake: (pot_size / (player_count + 1))}
    players := map[Player]int{human_player: 0} 
    for x := 1; x < (player_count + 2); x++ {
        var provisionary_name = fmt.Sprint("Player ", x)
        var previous_name = fmt.Sprint("Player ", (x - 1))
        if ((provisionary_name == player_name) || (provisionary_name == previous_name)) {
            provisionary_name = fmt.Sprint("Player ", (x + 1))
        }
        newPlayer := Player{name: provisionary_name, stake: (pot_size / (player_count + 1))}
        players[newPlayer] = x
    }

    fmt.Printf("Player name: %v; Amount of players: %v; Pot size: %v €; Blinds: big %v €, small %v €;\n", player_name, player_count, pot_size, big_blind_size, small_blind_size)
    fmt.Println(players)

    /*
    generate array with cards
    */
    // 127136 backside (red)
    // 127137 => 127150 Pik (black)
    // 127153 => 127166 Herz (rot)
    // 127169 => 127182 Rauten (rot)
    // 127185 => 127198 Kreuz (black)
    // var cards []string = make([]string, 52) // add hint for color
}

func main() {
    initialize_game()


    // ROUNDS
    /*

        // preflop

    raise minimum of big blind

        // evaluation
            // check from top to bottom who of the players has the best cards => if multiple, split pot
            // move pot to winning player(s)'(s) stake

    */

    // CHECKS
    /*
    who's turn is it? => shift turns after current round ends

    did someone go all in?

    is someone broke and should get deleted?
    is it only one player left? => winner announced, ends
    is the human player broke? => game lost, ends
    */

}
