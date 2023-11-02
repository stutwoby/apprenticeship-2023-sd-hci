package TakeAwayApp.MenuHandler;

import TakeAwayApp.*;

public class InitiateMenu {
    MenuBuilder[] menuCategory;
   DialogHandler dialogHandler = new DialogHandler();
    public InitiateMenu(){
        initiateMenu();
    }
    public void initiateMenu() {
        menuCategory = MenuBuilder.values();
        for (MenuBuilder category : menuCategory) {
            System.out.println(category);
            category.list();
            System.out.println("");
        }
        new ItemChecker(dialogHandler.getDialog("What would you like?"));
    }
}