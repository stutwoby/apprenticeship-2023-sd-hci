import java.util.Scanner;

public class Order {
//fields
    static String[] order = new String[6];
    static double[] price = new double[6];


//constructors
    Order()
    {
        intro();
       
    }

    
//methods
    public static void breakDown(String item, double itemPrice)
    {
        order[0] = item;
        price[0] = itemPrice;
        System.out.printf(order[0] + " £" + "%.2f",price[0]);
    }

    private void intro()
    {
        System.out.println("Welcome to Litva Kebabs.\nWould you like to place an order? \nType Yes or no: ");
        Scanner reponse = new Scanner(System.in);
        String aInput = reponse.nextLine();
       
        if (aInput.equals("no"))
            System.out.println("Have a nice day.");
        else
            aDrink();
            
        reponse.close();
    }

    private void aDrink()
    {
        System.out.println("What drink would you like: coke, lemonade, orange juice or iced tea? ");
        Scanner line = new Scanner(System.in);
        String input = line.nextLine();
        Drink drink = new Drink(input);
        line.close();
    }

}