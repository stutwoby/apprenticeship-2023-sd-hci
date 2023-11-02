package TakeAwayApp.ReceiptHandler;

import TakeAwayApp.DialogHandler;
import TakeAwayApp.MenuHandler.InitiateMenu;
import java.util.ArrayList;

public class CustomerAddress {

    private final ArrayList<String> CustomerDetails = new ArrayList<>();
    DialogHandler dialogHandler = new DialogHandler();


    public CustomerAddress() {
    }


    public void setAddress() {
        System.out.println("Please enter your delivery details:");
        addDetails(dialogHandler.getDialog("house name/number: "));
        addDetails(dialogHandler.getDialog("Street Address: "));
        addDetails(dialogHandler.getDialog("Town/City name: "));
        addDetails(dialogHandler.getDialog("Postcode: "));
        addDetails(dialogHandler.getDialog("Contact number: "));
        confirmAddress();
    }

    private void addDetails(String detail) {
        this.CustomerDetails.add(detail);
    }

   private void confirmAddress() {
        printDetails();
        if (dialogHandler.getDialog("Is this correct? yes/no").equals("yes"))
            new InitiateMenu();
        else {
            CustomerDetails.clear();
            System.out.println("<<<Please type carefully.>>>");
            setAddress();
        }
    }

    public void printDetails() {
        for (String i : this.CustomerDetails) {
            System.out.println(i);
        }
    }
}
