package TakeAwayApp.MenuHandler;

import TakeAwayApp.DialogHandler;
import TakeAwayApp.ReceiptHandler.Total;

public class MenuLoop {

    public MenuLoop(String customerDecision){
        menuLoop(customerDecision);

    }

    public void menuLoop(String string) {
        if (string.equals("yes")){
            new InitiateMenu();
        }
        else if (string.equals("no")){
            Total.getBreakDown();
        }
        else {
            System.out.println("Warning: Not an option.");
            menuLoop(new DialogHandler().getDialog("Would you like anything else? yes/no"));
        }
    }
}
