from words import words
import random


def check_word(word):
    word = list(word)
    for i in range(0, len(word)):
        char = word[i]
        if not char.isalpha():
            word.remove(word[i])
    return word


def get_random_word(words):
    return random.choice(words)


def check_guess(word, answer, guesses, wrong_guesses):
    correct_guess = False
    guess = input("enter your guess:")
    while 1 < len(guess):
        guess = input("enter a 1 letter guess:")

    if guess in word:
        for i in range(len(word)):
            if guess == word[i]:
                answer[i] = guess
                correct_guess = True

    if not correct_guess:
        guesses += 1
        wrong_guesses += guess

    return answer, guesses, wrong_guesses


def check_for_win(word, answer):
    for u in range(0, len(word)):
        if word[u] != answer[u]:
            return False
    else:
        return True


def print_answer(answer):
    string_answer = ""
    for p in range(0, len(answer)):
        string_answer += answer[p]
    print(string_answer)





def game_loop():
    # variables
    answer = []
    word = list(get_random_word(words))
    word = check_word(word)
    guesses = 0
    wrong_guesses = []

    # creating the answer list
    for i in range(0, len(word)):
        answer += "_"
    print_answer(answer)
    while guesses < 5 and check_for_win(word, answer) == False:
        answer, guesses, wrong_guesses = check_guess(word, answer, guesses, wrong_guesses)

        print("you have made", guesses, "wrong guesses. They are:", wrong_guesses)
        print_answer(answer)

    if check_for_win(word, answer):
        print("you win!")
    else:
        print("you lose!")
    print_answer(word)

    again = input("do you want to play again?(y/n)")
    if again == "y":
        game_loop()
    else:
        print("thanks for playing!")


game_loop()
