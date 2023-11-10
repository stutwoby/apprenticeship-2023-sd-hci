package TakeAwayApp;

import TakeAwayApp.ReceiptHandler.CustomerAddress;

public class Introduction {
    DialogHandler dialogHandler = new DialogHandler();
   public static CustomerAddress customerAddress = new CustomerAddress();

    public Introduction(){
        intro(dialogHandler.getDialog("Welcome to Litva Kebabs, would you like to make an order? yes/no"));
    }

    public void intro(String dialogHandler) {
            if (dialogHandler.equals("yes")) {
                customerAddress.setAddress();
            }
            else if (dialogHandler.equals("no"))
                System.out.println("Have a nice day.");
            else {
                System.out.println("Warning '" + dialogHandler + "' is not an option.\n");
                intro(this.dialogHandler.getDialog("Welcome to Litva Kebabs, would you like to make an order? yes/no"));
            }
        }
}

