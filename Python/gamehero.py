import random

# Define the hero and monster classes
class Hero:
    def __init__(self, name, health, attack):
        self.name = name
        self.health = health
        self.attack = attack

class Monster:
    def __init__(self, name, health, attack):
        self.name = name
        self.health = health
        self.attack = attack

# Define the stages
stages = [
    {"name": "Stage 1", "monsters": [Monster("Goblin", 10, 2), Monster("Orc", 15, 3)]},
    {"name": "Stage 2", "monsters": [Monster("Troll", 20, 4), Monster("Giant Spider", 25, 5)]},
    {"name": "Stage 3", "monsters": [Monster("Dragon", 50, 10)]}
]

# Define the game logic
def game():
    hero = Hero("Hero", 100, 10)
    current_stage = 0

    while current_stage < len(stages):
        stage = stages[current_stage]
        print(f"Welcome to {stage['name']}!")

        for monster in stage["monsters"]:
            print(f"A {monster.name} appears!")
            while monster.health > 0:
                print(f"Hero's health: {hero.health}, Monster's health: {monster.health}")
                action = input("What do you do? (A)ttack or (R)un? ")

                if action.lower() == "a":
                    monster.health -= hero.attack
                    print(f"You attack the {monster.name} for {hero.attack} damage!")
                elif action.lower() == "r":
                    print("You run away!")
                    break
                else:
                    print("Invalid action!")

                if monster.health > 0:
                    hero.health -= monster.attack
                    print(f"The {monster.name} attacks you for {monster.attack} damage!")

            if hero.health <= 0:
                print("You died! Game over.")
                return
            else:
                print(f"You defeated the {monster.name}!")

        current_stage += 1

    print("Congratulations, you won the game!")

# Start the game
game()