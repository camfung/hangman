package P1;
import java.util.*;

public class HangMan {
	final static String[] wordBank = {"thunder", "lightning", "cloud", "capital", "washington", "weather"};
	
	
	
	public static void main(String[] args) {
		List<Character> guesses = new ArrayList<>();
		List<Character> anslst = new ArrayList<>();
		List<Character> guessWord = new ArrayList<>();
		
		String[] man = {" O\n", "/", "|", "\\\n", "/", " \\\n"};  
		Scanner input = new Scanner(System.in);
		
		int wrongGuesses = 0;
		int correctGuesses = 0;
		
		// getting random word from bank
		Random gen = new Random();
		int index; 
		index = gen.nextInt(wordBank.length);
		String answer = wordBank[index];
		
		for (int i = 0; i < answer.length(); i++){
			anslst.add(answer.charAt(i));
			guessWord.add('_');
		}
		
		while (true) {
			boolean goodGuess = false;
			
			System.out.println("Enter your guess:");
			
			for (int i = 0; i < answer.length(); i++){
				System.out.print(guessWord.get(i) + " ");
			}
			
			
			System.out.println();
			char guess = input.next().charAt(0);
			boolean doubleGuess = false;
			for (int i = 0; i < guesses.size(); i++){
				if (guess == guesses.get(i))
				{
					System.out.println("You made this guess already!");
					doubleGuess = true;
					break;
				}
			}
			if (doubleGuess)
			{
				continue;
			}
			guesses.add(guess);
			for (int i = 0 ; i < answer.length(); i++){
				if (guess == anslst.get(i)) {
					guessWord.set(i, guess);
					goodGuess = true;
				}
			}
			
			if (!goodGuess) {
				wrongGuesses++;
				System.out.println("Wrong Guesses:" + wrongGuesses);
				
				 for (int i = 0; i < wrongGuesses; i++)
				 {
					 System.out.print(man[i]);
				 }
	
				 System.out.println();
			}
			else {
				correctGuesses++;
			}
			if (wrongGuesses > 5)
			{
				System.out.println("you lose! \nThe word was " 
									+ answer + ".");	
				break;
			}
			if (correctGuesses >= answer.length()) {
				System.out.println("you won!" + answer.length());
			}
			
		}
	}
}
