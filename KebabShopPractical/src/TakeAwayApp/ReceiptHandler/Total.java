package TakeAwayApp.ReceiptHandler;

import TakeAwayApp.Introduction;

import java.util.ArrayList;
public class Total{
    /* A sub-class of Takeaway creates an instance of total, used for end point of the program to calculate,
     delivery time and total price, methods needs to be brokendown.*/

    private static float delTime = 20.00f;
    static ArrayList<String> itemsCart = new ArrayList<>();
    static ArrayList<Float> pricesBroke = new ArrayList<>();
    static float total;


    Total() {
    }


    // prints out a breakdown of all items, and their prices, added to the cart. should be its own class?
    public static void getBreakDown() {
        System.out.println("\u001B[35m" + "\n\nDelivery Address:" +"\u001B[0m");
        Introduction.customerAddress.printDetails();


        System.out.println("\u001B[35m" + "\nBreakdown:" + "\u001B[0m");
        for (int i = 0; i < itemsCart.size(); i++) {
            System.out.printf(itemsCart.get(i) + "  £" + "%.2f", pricesBroke.get(i));
            System.out.println("");
        }

        System.out.printf("\u001B[35m" + "\nDelivery Fee: " + "\u001B[0m" + "£%.2f", 2.49f);
        System.out.println("");
        totalAmount();
    }

    //Prints the customers details, the delivery time estimated, and the total amount to the screen
    public static void totalAmount() {

        setPrice(2.49f);
        for (Float aFloat : pricesBroke) {
            total += aFloat;
        }
        System.out.printf("\u001B[35m" + "Total: " + "\u001B[0m" + "£%.2f", total);
        System.out.println("");
        System.out.printf("\u001B[35m"+ "Estimated delivery time: " + "\u001B[0m" + "%.2f minutes",delTime);
        System.out.println("");
        delTime = 20.00f;
    }

    // adds an item to itemsCart list.
    public static void setItem(String item) {
        itemsCart.add(item);
    }

    // adds an items price to the pricesCart list.
    public static void setPrice(Float price) {
        pricesBroke.add(price);
    }

    // adds 2.00 minutes to the estimated delivery time
    public static void setDelTime() {delTime += 2.00f;}
}
